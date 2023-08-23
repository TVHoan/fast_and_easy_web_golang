package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"golang/entitys"
	"golang/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	//handle route function
	dsn := "root:1@tcp(127.0.0.1:3306)/web_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(entitys.Register...)
	HandleRoute()
	handle := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		switch string(ctx.Method()) {
		case "GET":

			method, ok := router.get[path]
			if ok {
				method(ctx)
			} else {
				LoadStaticfile(ctx)
				//fmt.Println("method not found")
				//ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		case "POST":
			method, ok := router.get[path]
			if ok {
				method(ctx)
			} else {
				fmt.Println("method not found")
				ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		case "DELETE":
			method, ok := router.delete[path]
			if ok {
				method(ctx)
			} else {
				fmt.Println("method not found")
				ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		case "PUT":
			method, ok := router.put[path]
			if ok {
				method(ctx)
			} else {
				fmt.Println("method not found")
				ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		case "PATCH":
			method, ok := router.patch[path]
			if ok {
				method(ctx)
			} else {
				fmt.Println("method not found")
				ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		}

		//switch string(ctx.Path()) {
		//case "/foo":
		//	fooHandlerFunc(ctx)
		//default:
		//	ctx.Error("not found", fasthttp.StatusNotFound)
		//}
	}
	fmt.Println("Server Run in " + address)
	fasthttp.ListenAndServe(address, handle)
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
