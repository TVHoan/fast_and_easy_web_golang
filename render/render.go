package render

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"regexp"
)

func Html(path string, data interface{}, ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	filehtml, err := ioutil.ReadFile(path)
	regex := `{{\w+}}`
	var re = regexp.MustCompile(regex)
	htmlContent := string(filehtml)
	//newdata := map[interface{}]interface{}{}
	//iter := reflect.ValueOf(data).MapRange()
	//for iter.Next() {
	//	newdata[iter.Key().Interface()] = iter.Value().Interface()
	//}
	if m, _ := regexp.MatchString(regex, htmlContent); m {
		if matched := re.FindAllString(htmlContent, -1); matched != nil {
			//for _, value := range matched {
			//	keydata := strings.Trim(value, "{{")
			//	last_keydata := strings.Trim(keydata, "}}")
			//
			//	htmlContent = strings.Replace(htmlContent, value, string(data[last_keydata]), -1)
			//}
		}

	}
	if err == nil {

		fmt.Fprintf(ctx, string(htmlContent))
	} else {
		fmt.Println("Check your pathfile")
	}
}
