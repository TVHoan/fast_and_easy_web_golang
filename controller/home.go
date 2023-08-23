package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"

	"golang/render"
)

type App struct {
	*fasthttp.RequestCtx
}

func HomeIndex(ctx *fasthttp.RequestCtx) {
	render.Html("templates/index.html", nil, ctx)
}
func UserInfo(ctx *fasthttp.RequestCtx) {
	user := make(map[string]string)
	user["username"] = "admin"
	user["lang"] = "vi"
	user["name"] = "Trần Vũ Hoàn"
	user["title"] = "Trang chủ"
	if reponse, err := json.Marshal(user); err != nil {
		ctx.SetBody(reponse)
	} else {
		ctx.SetBody(reponse)
	}
}
