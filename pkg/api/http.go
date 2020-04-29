package api

import (
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/valyala/fasthttp"
	"math/rand"
	pprof2 "runtime/pprof"
	"strconv"
	"time"
)

type Application struct {
	gorm.Model
	Code string
	Usages uint
}

var Apps []Application
var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func RegisterHandlers(r *router.Router) {
	for i := 0; i < 50; i++ {
		app := Application{Code: getRandCode(), Usages: 0}
		Apps = append(Apps, app)
	}

	go updateApplications()

	r.GET("/request", handleRequest)
	r.GET("/admin/request", handleAdminRequest)

	r.GET("/debug/pprof/profile", handleProfilerRequest)
}

func handleRequest(ctx *fasthttp.RequestCtx) {
	randAppKey := rand.Intn(len(Apps))

	randApp := Apps[randAppKey]
	randApp.Usages++

	Apps[randAppKey] = randApp

	jsonValue, _ := json.Marshal(randApp.Code)
	ctx.Write(jsonValue)
}

func handleAdminRequest(ctx *fasthttp.RequestCtx) {
	values := map[string]string{}

	for _, app := range Apps {
		if app.Usages != 0 {
			values[app.Code] = strconv.Itoa(int(app.Usages))
		}
	}

	jsonValue, _ := json.Marshal(values)

	ctx.Write(jsonValue)
}

func handleProfilerRequest(ctx *fasthttp.RequestCtx) {
	ctx.Request.Header.Set("X-Content-Type-Options", "nosniff")

	sec := 30

	pprof2.StartCPUProfile(ctx.Response.BodyWriter())

	time.Sleep( time.Duration(sec)*time.Second)
	pprof2.StopCPUProfile()
}

func updateApplications() {
	t := time.NewTicker(200 * time.Millisecond)
	for range t.C {
		randAppKey := rand.Intn(len(Apps))
		Apps = append(Apps[:randAppKey], Apps[randAppKey + 1:]...)

		app := Application{Code: getRandCode(), Usages: 0}
		Apps = append(Apps, app)
	}
}

func getRandCode() string {
	c := make([]rune, 2)
	for i := range c {
		c[i] = letters[rand.Intn(len(letters))]
	}
	return string(c)
}