package services

import (
	"github.com/kunalpareek/golang-rest-api/models"
	"gopkg.in/mgo.v2/bson"
)

// bookDAO specifies the interface of the book DAO needed by BookService.
type bookDAO interface {
	// Get returns the book with the specified book ID.
	Get(id bson.ObjectId) (*models.Book, error)
	// Count returns the number of books.
	Count() (int, error)
	// Query returns the list of books with the given offset and limit.
	Query(offset, limit int) ([]models.Book, error)
	// Create saves a new book in the storage.
	Create(book *models.Book) error
	// Update updates the book with given ID in the storage.
	Update(id bson.ObjectId, book *models.Book) error
	// Delete removes the book with given ID from the storage.
	Delete(id bson.ObjectId) error
}

// BookService provides services related with books.
type BookService struct {
	dao bookDAO
}

// NewBookService creates a new BookService with the given book DAO.
func NewBookService(dao bookDAO) *BookService {
	return &BookService{dao}
}

// Get returns the book with the specified the book ID.
func (s *BookService) Get(id bson.ObjectId) (*models.Book, error) {
	return s.dao.Get(id)
}

// Create creates a new book.
func (s *BookService) Create(model *models.Book) (*models.Book, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(model); err != nil {
		return nil, err
	}
	return s.dao.Get(model.ID)
}

// Update updates the book with the specified ID.
func (s *BookService) Update(id bson.ObjectId, model *models.Book) (*models.Book, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(id)
}

// Delete deletes the book with the specified ID.
func (s *BookService) Delete(id bson.ObjectId) (*models.Book, error) {
	book, err := s.dao.Get(id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(id)
	return book, err
}

// Count returns the number of books.
func (s *BookService) Count() (int, error) {
	return s.dao.Count()
}

// Query returns the books with the specified offset and limit.
func (s *BookService) Query(offset, limit int) ([]models.Book, error) {
	return s.dao.Query(offset, limit)
}
