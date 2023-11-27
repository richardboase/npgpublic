package models

import "github.com/google/uuid"

type Product struct {
	Meta         Internals
	User         UserRef
	ID           string  `json:"id" firestore:"id"`
	Name         string  `json:"name" firestore:"name"`
	Description  string  `json:"description" firestore:"description"`
	Price        float64 `json:"price" firestore:"price"`
	Weight       string  `json:"weight" firestore:"weight"`
	Length       string  `json:"length" firestore:"length"`
	Width        string  `json:"width" firestore:"width"`
	Height       string  `json:"height" firestore:"height"`
	Manufacturer string  `json:"manufacturer" firestore:"manufacturer"`
	Model        string  `json:"model" firestore:"model"`
	SerialNumber string  `json:"sn" firestore:"sn"`
	ISBN         string  `json:"isbn" firestore:"isbn"`
	Country      string  `json:"country" firestore:"country"`
}

func (user *User) NewProduct() *Product {
	return &Product{
		Meta: NewInternals("product"),
		User: user.Ref(),
		ID:   uuid.NewString(),
	}
}
