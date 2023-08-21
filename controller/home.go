package controller

import (
	"github.com/valyala/fasthttp"
	"golang/render"
)

func HomeIndex(ctx *fasthttp.RequestCtx) {
	//array := make(map[string]string)
	//array["name"] = "hoan"
	//array["age"] = "27"
	//array["born"] = "Hai Phong"
	//result, oke := json.Marshal(array)
	//if oke == nil {
	ctx.SetContentType("text/html; charset=utf-8")
	render.Html("templates/index.html", nil, ctx)
	//}
	// Write the HTML content to the response

}
