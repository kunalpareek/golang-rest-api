package main

import (
	"log"
	"net/http"

	"github.com/kunalpareek/golang-rest-api/config"

	"github.com/gorilla/mux"
	"github.com/kunalpareek/golang-rest-api/apis"
	"github.com/kunalpareek/golang-rest-api/daos"
	"github.com/kunalpareek/golang-rest-api/middleware"
	"github.com/kunalpareek/golang-rest-api/services"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	conf, err := config.LoadConfig()
	session, err := mgo.Dial(conf.DbConnectionString)
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	if err := http.ListenAndServe(":"+conf.Port, buildRouter(session)); err != nil {
		log.Fatal(err)
	}
}

func buildRouter(db *mgo.Session) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.RequestLogging)

	bookDao := daos.NewBookDAO(db)
	r = apis.ServeBookResource(r, services.NewBookService(bookDao))

	return r
}
