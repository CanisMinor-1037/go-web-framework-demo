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
	Name string `v:"required" dc:"姓名"`
	Age  int `v:"required" dc:"年龄"`
}

type HelloRes struct {}
type Hello struct {}

func (Hello) Say(ctx context.Context, helloReq *HelloReq) (HelloRes *HelloRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.Writef("Hello, %s!, Your age is %d",
		helloReq.Name,
		helloReq.Age,
	)
	return
}

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.Write("error occurs: ", err.Error())
		return
	}
}

func main() {
	fmt.Println("Hello, GF:", gf.VERSION)
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ErrorHandler)
		group.Bind(
			new(Hello),
		)
	})
	s.SetPort(8080)
	s.Run()
}
