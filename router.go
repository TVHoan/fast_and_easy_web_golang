package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"golang/controller"
)

func handleroute() {
	router.get["/a"] = handleA
	router.get["/"] = controller.HomeIndex
}
func handleA(ctx *fasthttp.RequestCtx) {
	array := make(map[string]string)
	array["name"] = "hoan"
	array["age"] = "27"
	array["born"] = "Hai Phong"
	result, err := json.Marshal(array)
	if err != nil {
		ctx.SetBody(result)
	} else {
		ctx.SetBody(result)
	}
}
func fooHandlerFunc(ctx *fasthttp.RequestCtx) {
	array := make(map[string]string)
	array["name"] = "hoan"
	array["age"] = "27"
	array["born"] = "Hai Phong"
	result, err := json.Marshal(array)
	if err != nil {
		ctx.SetBody(result)
	} else {
		ctx.SetBody(result)
	}
}