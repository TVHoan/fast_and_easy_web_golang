package helper

type Type string

var Typefile = map[string]string{
	".js":   "text/javascript; charset=utf-8",
	".css":  "text/css; charset=utf-8",
	".scss": "text/scss; charset=utf-8",
}

var Acept = []string{
	".js",
	".css",
	".scss",
}
