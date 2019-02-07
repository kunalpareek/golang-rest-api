package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Book struct
type Book struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Author    string        `json:"author" bson:"author"`
	ISBN      string        `json:"isbn" bson:"isbn"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

// Validate validates the Book fields.
func (b Book) Validate() error {
	return nil
}
