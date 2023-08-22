package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"golang/auth"
	"golang/database"
	"golang/entity"
)

func Register(ctx *fasthttp.RequestCtx) {
	username := string(ctx.FormValue("username"))
	password := string(ctx.FormValue("password"))
	hash_pw, _ := auth.HashPassword(password)
	user := entity.User{
		Username: username,
		Password: hash_pw,
	}
	db := database.Db()
	result := db.Create(&user)
	if reponse, err := json.Marshal(result); err != nil {
		ctx.SetBody(reponse)
	} else {
		ctx.SetBody(reponse)
	}
}
func Login(ctx *fasthttp.RequestCtx) {
	username := string(ctx.FormValue("username"))
	password := string(ctx.FormValue("password"))
	db := database.Db()
	user := entity.User{}
	db.Where("username = (?)", username).Select("username", "password").Find(&user)
	isLogin := auth.CheckPasswordHash(password, user.Password)
	if isLogin {
		token, _ := auth.CreateToken(username)

		cook := fasthttp.Cookie{}
		cook.SetKey("token")
		cook.SetValue(token)
		cook.SetMaxAge(3600000)
		cook.SetSameSite(fasthttp.CookieSameSiteLaxMode)
		ctx.Response.Header.SetCookie(&cook)
		reponse, _ := json.Marshal(map[string]string{
			"authentication": "success",
		})
		ctx.SetBody(reponse)
	} else {
		reponse, _ := json.Marshal(map[string]string{
			"authentication": "fail",
		})
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBody(reponse)
	}

}
