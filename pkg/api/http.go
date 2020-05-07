package api

import (
	"bytes"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"math/rand"
	"runtime/pprof"
	"stresstest/pkg/model"
	"time"
)

// RegisterHandlers just registers handlers
func RegisterHandlers(r *router.Router) {
	r.GET("/request", handleRequest)
	r.GET("/admin/request", handleAdminRequest)

	// pprof request handler
	r.GET("/debug/pprof/profile", handleProfilerRequest)
}

func handleRequest(ctx *fasthttp.RequestCtx) {
	randOrderKey := rand.Int63() % int64(len(model.Orders()))

	model.OrdMutex.Lock()
	defer model.OrdMutex.Unlock()

	randApp := &model.Orders()[randOrderKey]
	randApp.Usages++

	model.Orders()[randOrderKey] = *randApp

	ctx.WriteString(randApp.Code)
}

func handleAdminRequest(ctx *fasthttp.RequestCtx) {
	buf := new(bytes.Buffer)

	for _, o := range model.Orders() {
		fmt.Fprintf(buf, "%s-%d\n", o.Code, o.Usages)
	}

	ctx.Write(buf.Bytes())
}

func handleProfilerRequest(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Set("X-Content-Type-Options", "nosniff")

	pprof.StartCPUProfile(ctx.Response.BodyWriter())

	time.Sleep(time.Duration(30) * time.Second)
	pprof.StopCPUProfile()
}
