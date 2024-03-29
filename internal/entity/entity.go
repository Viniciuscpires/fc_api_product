package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string
	Name        string
	Description string
	CategoryID  string
	ImageUrl    string
	Price       float64
}

func NewProduct(name, description, categoryId, imageUrl string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CategoryID:  categoryId,
		ImageUrl:    imageUrl,
		Price:       price,
	}
}
