package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kunalpareek/golang-rest-api/models"
	"gopkg.in/mgo.v2/bson"
)

type (
	// bookService specifies the interface for the book service needed by bookResource.
	bookService interface {
		Get(id bson.ObjectId) (*models.Book, error)
		Query(offset, limit int) ([]models.Book, error)
		Count() (int, error)
		Create(model *models.Book) (*models.Book, error)
		Update(id bson.ObjectId, model *models.Book) (*models.Book, error)
		Delete(id bson.ObjectId) (*models.Book, error)
	}

	// bookResource defines the handlers for the CRUD APIs.
	bookResource struct {
		service bookService
	}
)

// ServeBookResource sets up the routing of book endpoints and the corresponding handlers.
func ServeBookResource(r *mux.Router, service bookService) *mux.Router {
	resource := &bookResource{service}
	// r.Get("/books", resource.get)
	r.HandleFunc("/books", resource.get).Methods("GET")
	r.HandleFunc("/books", resource.create).Methods("POST")
	// r.Get("/books", resource.query)
	// r.Post("/books", resource.create)
	// r.Put("/books/<id>", resource.update)
	// r.Delete("/books/<id>", resource.delete)
	return r
}

func (b *bookResource) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	data, err := b.service.Get(bson.ObjectIdHex(id))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error in fetching data")
	}

	respondWithJson(w, 200, &data)
}

func (b *bookResource) query(w http.ResponseWriter, r *http.Request) {

}

func (b *bookResource) create(w http.ResponseWriter, r *http.Request) {
	var model models.Book
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	data, err := b.service.Create(&model)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error in saving data")
	}

	respondWithJson(w, 200, &data)
}

func (b *bookResource) update(w http.ResponseWriter, r *http.Request) {

}

func (b *bookResource) delete() {

}
