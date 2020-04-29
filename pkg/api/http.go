package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Application struct {
	gorm.Model
	Code string
	Canceled bool
	Usages uint
}

var Apps []Application
var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var db *gorm.DB

func RegisterHandlers(r *mux.Router) {
	db1, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		errors.Wrap(err, "failed to connect database")
	}

	db = db1

	//defer db.Close()

	db.DropTableIfExists(&Application{})
	db.AutoMigrate(&Application{})

	for i := 0; i < 50; i++ {
		app := Application{Code: getRandCode(), Canceled: false, Usages: 0}
		db.Create(&app)
		Apps = append(Apps, app)
	}

	go timer(db)

	r.HandleFunc("/request", handleRequest).Methods("GET")
	r.HandleFunc("/admin/request", handleAdminRequest).Methods("GET")
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	randAppKey := rand.Intn(len(Apps))

	randApp := Apps[randAppKey]
	randApp.Usages++
	db.Save(&randApp)

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

func timer(db *gorm.DB) {
	t := time.NewTicker(200 * time.Millisecond)
	for range t.C {
		randAppKey := rand.Intn(len(Apps))

		randApp := Apps[randAppKey]
		db.Delete(&randApp)

		Apps = append(Apps[:randAppKey], Apps[randAppKey + 1:]...)

		app := Application{Code: getRandCode(), Canceled: false, Usages: 0}
		db.Create(&app)
		Apps = append(Apps, app)
		log.Println(len(Apps))
	}
}

func getRandCode() string {
	c := make([]rune, 2)
	for i := range c {
		c[i] = letters[rand.Intn(len(letters))]
	}
	return string(c)
}