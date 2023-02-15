package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	accountapi "github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	userapi "github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
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
			type UserParams struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			var userParams UserParams
			err := c.Bind(&userParams)
			if err != nil {
				return int64(0), jwt.ErrMissingLoginValues
			}
			log.Println(userParams)
			username := userParams.Username
			password := userParams.Password

			if username == "admin" && password == "12345" {
				return int64(1), nil
			} else {
				return int64(0), jwt.ErrFailedAuthentication
			}
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

	//用户相关api
	user := h.Group("user/")
	//用户登录
	user.POST("login/", authMiddleware.LoginHandler)
	//刷新jwt token
	user.POST("refresh/", authMiddleware.RefreshHandler)
	//用户注册
	user.POST("register/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "register succeed",
		})
	})

	user.Use(authMiddleware.MiddlewareFunc())
	//获取用户信息
	user.GET("/", func(ctx context.Context, c *app.RequestContext) {
		//从jwt token payload中获取用户id
		claims := jwt.ExtractClaims(ctx, c)
		userID := int64(claims[constants.IdentityKey].(float64))
		if userID == 1 {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 0,
				"status_msg":  "string",
				"user": &userapi.User{
					UserId:   1,
					Username: "admin",
				},
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 1,
				"status_msg":  "authorization failed",
			})
		}
	})

	//account相关api
	account := h.Group("account/")
	account.Use(authMiddleware.MiddlewareFunc())
	//创建帐户
	account.POST("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "create succeed",
		})
	})
	//删除帐户
	account.DELETE(":account_id/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "delete succeed",
		})
	})
	//修改帐户
	account.PUT(":account_id/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "update succeed",
		})
	})
	//查询帐户
	account.GET(":account_id/", func(ctx context.Context, c *app.RequestContext) {
		accountID, _ := strconv.ParseInt(c.Param("account_id"), 10, 64)
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "string",
			"account": accountapi.Account{
				AccountId:   accountID,
				AccountName: "account_name",
				Balance:     0,
			},
		})
	})
	//查询用户所有帐户
	account.GET("list/:user_id/", func(ctx context.Context, c *app.RequestContext) {
		userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "string",
			"accounts": []*accountapi.Account{
				{
					AccountId:   userID,
					AccountName: "account_name",
					Balance:     0,
				},
			},
		})
	})

	h.Spin()
}
