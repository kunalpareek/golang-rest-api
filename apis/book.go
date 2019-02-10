package apis

import (
	"encoding/json"
	"net/http"

	"github.com/kunalpareek/golang-rest-api/utils"

	"github.com/gorilla/mux"
	"github.com/kunalpareek/golang-rest-api/models"
	"gopkg.in/mgo.v2/bson"
)

type (
	// bookService specifies the interface for the book service needed by bookResource.
	bookService interface {
		Get(id bson.ObjectId) (*models.Book, error)
		Query(queryParam interface{}, offset, limit int) ([]models.Book, error)
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
	r.HandleFunc("/books", resource.get).Methods("GET")
	r.HandleFunc("/books", resource.create).Methods("POST")
	r.HandleFunc("/books", resource.update).Methods("PUT")
	r.HandleFunc("/books", resource.delete).Methods("DELETE")
	return r
}

func (b *bookResource) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	author := r.URL.Query().Get("author")
	if author != "" {
		data, err := b.service.Query(bson.M{"author": author}, 0, 10)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error in fetching data")
		}

		utils.RespondWithJSON(w, http.StatusOK, &data)
	} else {
		data, err := b.service.Get(bson.ObjectIdHex(id))
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error in fetching data")
		}

		utils.RespondWithJSON(w, http.StatusOK, &data)
	}
}

func (b *bookResource) create(w http.ResponseWriter, r *http.Request) {
	var model models.Book
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	data, err := b.service.Create(&model)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error in saving data")
	}

	utils.RespondWithJSON(w, http.StatusOK, &data)
}

func (b *bookResource) update(w http.ResponseWriter, r *http.Request) {
	var model models.Book
	id := r.URL.Query().Get("id")

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	data, err := b.service.Update(bson.ObjectIdHex(id), &model)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error in updating data")
	}

	utils.RespondWithJSON(w, http.StatusOK, &data)
}

func (b *bookResource) delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	data, err := b.service.Delete(bson.ObjectIdHex(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error in deleting data")
	}

	utils.RespondWithJSON(w, http.StatusOK, &data)
}
