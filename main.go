//Package Stresstest is just an example of solving a test task.
package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"stresstest/pkg/api"
)

func main() {
	r := router.New()
	api.RegisterHandlers(r)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}