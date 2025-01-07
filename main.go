package main

import (
	"fmt"
	"context"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	g.Meta `path:"/" method:"get"`
	Name string `v:"required" json:"name" dc:"姓名"`
	Age  int `v:"required" json:"age" dc:"年龄"`
}

type HelloRes struct {
	Content string `json:"content" dc:"返回结果"`
}
type Hello struct {}

type Response struct {
	Message string `json:"message" dc:"消息提示"`
	Data interface{} `json:"data" dc:"执行结果"`
}

func (Hello) Say(ctx context.Context, helloReq *HelloReq) (helloRes *HelloRes, err error) {
	helloRes = &HelloRes{
		Content: fmt.Sprintf(
			"Hello, %s!, Your age is %d",
			helloReq.Name,
			helloReq.Age,
		),
	}
	return
}

// 定义中间件
func Middleware(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg string
		res = r.GetHandlerResponse()
		err = r.GetError()
	)

	if err != nil {
		msg = err.Error()
	} else {
		msg = "OK"
	}
	r.Response.WriteJson(Response{
		Message: msg,
		Data: res,
	})
}


func main() {
	fmt.Println("Hello, GF:", gf.VERSION)
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(Middleware)
		group.Bind(
			new(Hello),
		)
	})
	s.SetPort(8080)
	s.Run()
}
