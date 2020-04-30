package api

import (
	"bytes"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"math/rand"
	"runtime/pprof"
	"strconv"
	"time"
)

type Application struct {
	Code   string
	Usages uint
}

var apps [50]Application
var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func RegisterHandlers(r *router.Router) {
	for i := 0; i < 50; i++ {
		apps[i] = Application{Code: getRandCode(), Usages: 0}
	}

	go updateApplications()

	r.GET("/request", handleRequest)
	r.GET("/admin/request", handleAdminRequest)

	r.GET("/debug/pprof/profile", handleProfilerRequest)
}

func handleRequest(ctx *fasthttp.RequestCtx) {
	randAppKey := rand.Int63() % int64(len(letters))

	randApp := apps[randAppKey]
	randApp.Usages++

	apps[randAppKey] = randApp

	ctx.WriteString(randApp.Code)
}

func handleAdminRequest(ctx *fasthttp.RequestCtx) {
	values := map[string]string{}

	for _, app := range apps {
		values[app.Code] = strconv.Itoa(int(app.Usages))
	}

	ctx.Write(createKeyValuePairs(values))
}

func createKeyValuePairs(m map[string]string) []byte {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s-%s", key, value)
	}
	return b.Bytes()
}

func handleProfilerRequest(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Set("X-Content-Type-Options", "nosniff")

	pprof.StartCPUProfile(ctx.Response.BodyWriter())

	time.Sleep(time.Duration(30) * time.Second)
	pprof.StopCPUProfile()
}

func updateApplications() {
	t := time.NewTicker(200 * time.Millisecond)
	for range t.C {
		apps[rand.Int63() % int64(len(letters))] = Application{Code: getRandCode(), Usages: 0}
	}
}

func getRandCode() string {
	c := make([]rune, 2)
	for i := range c {
		c[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(c)
}
