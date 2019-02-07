package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunalpareek/golang-rest-api/apis"
	"github.com/kunalpareek/golang-rest-api/daos"
	"github.com/kunalpareek/golang-rest-api/middleware"
	"github.com/kunalpareek/golang-rest-api/services"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017/lookup")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// r := mux.NewRouter()
	// r.Use(middleware.Auth)
	// r = ip.Routes(r)
	if err := http.ListenAndServe(":3000", buildRouter(session)); err != nil {
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
