package models

import (
	"github.com/google/uuid"
)

type Product struct {
	//ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Code        uuid.UUID `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}
