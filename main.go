package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"golang/database"
	"golang/entity"
	"golang/helper"
	"golang/middleware"
	"log"
	"os"
	"regexp"
	"strings"
)

var address = ":8001"

type route struct {
	get    map[string]func(ctx *fasthttp.RequestCtx)
	post   map[string]func(ctx *fasthttp.RequestCtx)
	put    map[string]func(ctx *fasthttp.RequestCtx)
	delete map[string]func(ctx *fasthttp.RequestCtx)
	patch  map[string]func(ctx *fasthttp.RequestCtx)
}

var router = route{
	get:    make(map[string]func(ctx *fasthttp.RequestCtx)),
	post:   make(map[string]func(ctx *fasthttp.RequestCtx)),
	put:    make(map[string]func(ctx *fasthttp.RequestCtx)),
	delete: make(map[string]func(ctx *fasthttp.RequestCtx)),
	patch:  make(map[string]func(ctx *fasthttp.RequestCtx)),
}

func main() {
	db := database.Db()
	db.AutoMigrate(entity.RegisterEntity...)
	//handle route function
	handleroute()
	handle := func(ctx *fasthttp.RequestCtx) {
		Routing(ctx)
	}
	//define generic middleware
	middleware := middleware.GenericMiddleWare(handle)

	fmt.Println("Server Run in " + address)
	fasthttp.ListenAndServe(address, middleware)
}
func IsPathRoute(path string) bool {
	if strings.Index(path, "/") == 0 {
		return true
	}
	return false
}
func CheckPathRoute(path string) {
	if !IsPathRoute(path) {
		fmt.Println("path in route must have '/' first ")
	}
}
func LoadStaticfile(ctx *fasthttp.RequestCtx) {
	for _, value := range helper.Acept {

		regex := value + `$`
		path := string(ctx.Path())
		data, err := os.ReadFile(strings.TrimLeft(path, "/"))
		if err != nil {
			log.Printf("Error reading file: %s", err)
			return
		}
		if m, _ := regexp.MatchString(regex, path); m {
			ctx.SetContentType(helper.Typefile[value])
			ctx.SetBody(data)
			return
		}
	}
}
func Routing(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	switch string(ctx.Method()) {
	case "GET":

		method := router.get[path]
		if method != nil {
			method(ctx)
		} else {
			LoadStaticfile(ctx)
			//fmt.Println("method not found")
			//ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		}
	case "POST":
		method := router.post[path]
		if method != nil {
			method(ctx)
		} else {
			fmt.Println("method not found")
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		}
	case "DELETE":
		method := router.delete[path]
		if method != nil {
			method(ctx)
		} else {
			fmt.Println("method not found")
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		}
	case "PUT":
		method := router.put[path]
		if method != nil {
			method(ctx)
		} else {
			fmt.Println("method not found")
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		}
	case "PATCH":
		method := router.patch[path]
		if method != nil {
			method(ctx)
		} else {
			fmt.Println("method not found")
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		}
	}
}
func init() {
	fmt.Println(`    
	  _____      ______
	 / ___/     //    \\
	/ / ____    ||    ||
	| | |_ |    ||    ||
	| |__| |    | \__/ |   
	\______/     \____/  `)
}
