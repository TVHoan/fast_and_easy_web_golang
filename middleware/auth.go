package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"golang/auth"
)

func Auth(ctx *fasthttp.RequestCtx) *fasthttp.RequestCtx {

	tokenFormCookie := string(ctx.Request.Header.Cookie("token"))
	token, err := auth.ParseToken(tokenFormCookie)
	if tokenFormCookie != "" || err != nil {
		return ctx
	} else {
		res, _ := json.Marshal(map[string]string{"messenge": "unauthenticaton"})
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBody(res)
	}
	claims := token.Claims.(jwt.MapClaims)
	ctx.SetUserValue("username", claims["username"])
	return ctx
}
func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// Thực hiện các bước xác thực tại đây
		// Ví dụ: kiểm tra mã thông báo, token,...

		tokenFormCookie := string(ctx.Request.Header.Cookie("token"))
		token, err := auth.ParseToken(tokenFormCookie)
		if tokenFormCookie == "" || err != nil || !token.Valid {
			res, _ := json.Marshal(map[string]string{"messenge": "unauthenticaton"})
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBody(res)
		} else {
			claims := token.Claims.(jwt.MapClaims)
			ctx.SetUserValue("username", claims["username"])
			// Nếu xác thực thành công, gọi next() để tiếp tục xử lý RequestHandler kế tiếp
			next(ctx)
		}

	}
}
