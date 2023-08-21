package helper

type Type string

const (
	javascript Type = ".js"
	css        Type = ".css"
	scss       Type = ".scss"
)

var Acept = []string{
	".js",
	".css",
	".scss",
}
