package models

type Product struct {
	//ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
