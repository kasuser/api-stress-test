package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"net/http"
	"net/http/pprof"
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

func RegisterHandlers(r *mux.Router) {
	for i := 0; i < 50; i++ {
		app := Application{Code: getRandCode(), Usages: 0}
		Apps = append(Apps, app)
	}

	go updateApplications()

	r.HandleFunc("/request", handleRequest).Methods("GET")
	r.HandleFunc("/admin/request", handleAdminRequest).Methods("GET")

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	randAppKey := rand.Intn(len(Apps))

	randApp := Apps[randAppKey]
	randApp.Usages++

	Apps[randAppKey] = randApp

	jsonValue, _ := json.Marshal(randApp.Code)
	writer.Write(jsonValue)
}

func handleAdminRequest(writer http.ResponseWriter, request *http.Request) {
	values := map[string]string{}

	for _, app := range Apps {
		if app.Usages != 0 {
			values[app.Code] = strconv.Itoa(int(app.Usages))
		}
	}

	jsonValue, _ := json.Marshal(values)

	writer.Write(jsonValue)
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