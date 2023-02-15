package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/rdstihz/SimpleAccountBook/cmd/api/hanlder"
	"github.com/rdstihz/SimpleAccountBook/cmd/api/rpc"
	userapi "github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"github.com/rdstihz/SimpleAccountBook/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	h := server.New(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithHandleMethodNotAllowed(true),
	)

	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,

		//用户登录, 登录成功返回用户id
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			// 获取传入的username和password
			var loginValues handler.UserParam
			err := c.Bind(&loginValues)
			if err != nil {
				return int64(0), jwt.ErrMissingLoginValues
			}
			username := loginValues.Username
			password := loginValues.Password
			if len(username) == 0 || len(password) == 0 {
				return int64(0), jwt.ErrMissingLoginValues
			}
			// 调用rpc验证登录，获取user id
			return rpc.CheckUser(&userapi.CheckUserRequest{
				Username: username,
				Password: password,
			})
		},

		//向jwt token payload中写入用户id
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		//登录成功的响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"status_code": 0,
				"status_msg":  "success",
				"user_id":     0,
				"token":       token,
				"expire":      expire.Format(time.RFC3339),
			})
		},

		//登录失败的响应
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": 1,
				"status_msg":  "authorization failed",
			})
		},

		//刷新token的响应
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 0,
				"status_msg":  "success",
				"token":       token,
				"expire":      expire.Format(time.RFC3339),
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		IdentityKey:   constants.IdentityKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	// 添加中间件，打印日志信息
	h.Use(middleware.HertzCommonMiddleware())
	//用户相关api
	user := h.Group("user/")
	//用户登录
	user.POST("login/", authMiddleware.LoginHandler)
	//刷新jwt token
	user.POST("refresh/", authMiddleware.RefreshHandler)
	//用户注册
	user.POST("register/", handler.RegisterHandler)

	user.Use(authMiddleware.MiddlewareFunc())
	//获取用户信息
	user.GET("/", handler.GetUserInfoHandler)

	//account相关api
	account := h.Group("account/")
	account.Use(authMiddleware.MiddlewareFunc())
	//创建帐户
	account.POST("/", handler.CreateAccountHandler)
	//删除帐户
	account.DELETE(":account_id/", handler.DeleteAccountHandler)
	//修改帐户
	account.PUT(":account_id/", handler.UpdateAccountHandler)
	//查询帐户
	account.GET(":account_id/", handler.GetAccountInfoHandler)
	//查询用户所有帐户
	account.GET("list/", handler.ListAccountHandler)

	//bill相关api
	bill := h.Group("bill/")
	bill.Use(authMiddleware.MiddlewareFunc())
	//创建bill
	bill.POST("/", handler.CreateBillHandler)
	//删除bill
	bill.DELETE(":bill_id/", handler.DeleteBillHandler)
	//修改bill
	bill.PUT(":bill_id/", handler.UpdateBillHandler)
	//查询bill
	bill.GET(":bill_id/", handler.GetBillInfoHandler)
	//查询account下所有bill
	bill.GET("list/", handler.ListBillHandler)
	h.Spin()
}
