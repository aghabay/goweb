package main

//
// go run .\main.go .\handlers.go .\render.go
//go run .\cmd\web\main.go .\cmd\handlers\handlers.go .\cmd\render\render.go

import (
	"net/http"

	"github.com/aghabay/goweb/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(":9000", nil)
}
