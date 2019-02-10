package daos

import (
	"time"

	"github.com/kunalpareek/golang-rest-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BookDAO persists Book data in database
type BookDAO struct {
	dbName         string
	collectionName string
	dbConnection   *mgo.Session
}

// NewBookDAO creates a new BookDAO
func NewBookDAO(session *mgo.Session) *BookDAO {
	// err := session.DB("lookup").C("books")
	// if err != nil {
	// 	panic(err)
	// }
	return &BookDAO{"lookup", "books", session}
}

// Get reads the Book with the specified ID from the database.
func (dao *BookDAO) Get(id bson.ObjectId) (*models.Book, error) {

	var Book models.Book
	if err := dao.dbConnection.DB(dao.dbName).C(dao.collectionName).FindId(id).One(&Book); err != nil {
		return nil, err
	}
	return &Book, nil
}

// Create saves a new Book record in the database.
// The Book.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BookDAO) Create(Book *models.Book) error {
	Book.ID = bson.NewObjectId()
	currentTime := time.Now()
	Book.CreatedAt = currentTime
	Book.UpdatedAt = currentTime

	return dao.dbConnection.DB(dao.dbName).C(dao.collectionName).Insert(&Book)
}

// Update saves the changes to a Book in the database.
func (dao *BookDAO) Update(id bson.ObjectId, Book *models.Book) error {
	if _, err := dao.Get(id); err != nil {
		return err
	}

	Book.ID = id
	currentTime := time.Now()
	Book.UpdatedAt = currentTime

	return dao.dbConnection.DB(dao.dbName).C(dao.collectionName).Update(bson.M{"_id": id}, Book)
}

// Delete deletes an Book with the specified ID from the database.
func (dao *BookDAO) Delete(id bson.ObjectId) error {
	return dao.dbConnection.DB(dao.dbName).C(dao.collectionName).RemoveId(id)
}

// Count returns the number of the Book records in the database.
func (dao *BookDAO) Count() (int, error) {
	var count int
	count, err := dao.dbConnection.DB(dao.dbName).C(dao.collectionName).Count()
	return count, err
}

// Query retrieves the Book records with the specified offset and limit from the database.
func (dao *BookDAO) Query(queryParam interface{}, offset, limit int) ([]models.Book, error) {
	Books := []models.Book{}
	err := dao.dbConnection.DB(dao.dbName).C(dao.collectionName).Find(queryParam).All(&Books)
	return Books, err
}
