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
	r.GET("/request", func(ctx *fasthttp.RequestCtx) { ctx.Write(HandleRequest()) })
	r.GET("/admin/request", func(ctx *fasthttp.RequestCtx) { ctx.Write(HandleAdminRequest()) })

	// pprof request handler
	r.GET("/debug/pprof/profile", handleProfilerRequest)
}

func HandleRequest() []byte {
	randOrderKey := rand.Int63() % int64(len(model.Orders()))

	model.OrdMutex.Lock()
	defer model.OrdMutex.Unlock()

	randApp := &model.Orders()[randOrderKey]
	randApp.Usages++

	model.Orders()[randOrderKey] = *randApp

	return []byte(randApp.Code)
}

func HandleAdminRequest() []byte {
	buf := new(bytes.Buffer)

	for _, o := range model.Orders() {
		fmt.Fprintf(buf, "%s-%d\n", o.Code, o.Usages)
	}

	return buf.Bytes()
}

func handleProfilerRequest(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Set("X-Content-Type-Options", "nosniff")

	pprof.StartCPUProfile(ctx.Response.BodyWriter())

	time.Sleep(time.Duration(30) * time.Second)
	pprof.StopCPUProfile()
}
