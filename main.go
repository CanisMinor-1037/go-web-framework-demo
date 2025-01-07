package main

import (
	"fmt"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Hello, GF:", gf.VERSION)
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
			var user User
			if err := r.Parse(&user); err != nil {
				r.Response.Write(err.Error())
				return
			}
			if user.Name == "" {
				r.Response.Write("Name should not be empty")
				return
			}
			if user.Age <= 0 {
				r.Response.Write("invalid age value")
				return
			}
			r.Response.Writef("Hello, %s!, Your age is %d",
			user.Name,
			user.Age,
		)
	})
	s.SetPort(8080)
	s.Run()
}
