package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"testpr/pkg/api"
)

func main() {
	r := router.New()
	api.RegisterHandlers(r)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}