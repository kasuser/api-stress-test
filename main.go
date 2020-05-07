//Package Stresstest is just an example of solving a test task.
package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"stresstest/pkg/api"
	"stresstest/pkg/model"
	"time"
)

func main() {
	model.InitializeOrders()

	go func() {
		t := time.NewTicker(200 * time.Millisecond)
		for range t.C {
			model.UpdateActualOrders()
		}
	}()

	r := router.New()
	api.RegisterHandlers(r)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}