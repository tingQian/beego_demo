package main

import (
	"net/http"
	"github.com/astaxie/beego/context"
	_ "hello/routers"
	"github.com/astaxie/beego"
	"strings"
)

//add client
func TransparentClient(ctx *context.Context) {
	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "client/"+ctx.Request.URL.Path)
}

func main() {
	beego.InsertFilter("/", beego.BeforeRouter, TransparentClient) 
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentClient)
	beego.Run()
}

