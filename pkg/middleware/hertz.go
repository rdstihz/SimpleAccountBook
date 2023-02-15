package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

// HertzCommonMiddleware Hertz中间件，用于打印日志信息
func HertzCommonMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		remoteAddr := c.RemoteAddr().String()
		uri := c.URI().String()
		method := string(c.Method())
		log.Printf("[%v] %v %v\n", remoteAddr, method, uri)
		queryArgs := c.QueryArgs()
		params := c.Params
		raw_data := c.GetRawData()
		log.Println("QueryArgs:", queryArgs)
		log.Println("params:", params)
		log.Println("raw_data:", string(raw_data))
		c.Next(ctx)
	}
}
