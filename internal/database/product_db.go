package database

import (
	"database/sql"

	"github.com/Viniciuscpires/fc_api_product/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (cd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := cd.db.Query("SELECT id, name, description, category_id, image_url, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var Product entity.Product
		if err := rows.Scan(&Product.ID, &Product.Name, &Product.Description, &Product.CategoryID, &Product.ImageUrl, &Product.Price); err != nil {
			return nil, err
		}
		products = append(products, &Product)
	}
	return products, nil
}

func (cd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	row := cd.db.QueryRow("SELECT id, name, description, category_id, image_url, price FROM products WHERE id = ?", id)
	var Product entity.Product
	if err := row.Scan(&Product.ID, &Product.Name, &Product.Description, &Product.CategoryID, &Product.ImageUrl, &Product.Price); err != nil {
		return nil, err
	}
	return &Product, nil
}

func (cd *ProductDB) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	rows, err := cd.db.Query("SELECT id, name, description, category_id, image_url, price FROM products where category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var Product entity.Product
		if err := rows.Scan(&Product.ID, &Product.Name, &Product.Description, &Product.CategoryID, &Product.ImageUrl, &Product.Price); err != nil {
			return nil, err
		}
		products = append(products, &Product)
	}
	return products, nil
}

func (cd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := cd.db.Exec("INSERT INTO products (id, name, description, category_id, image_url, price) VALUES (? ? ? ? ? ?)", product.ID, product.Name, product.Description, product.CategoryID, product.ImageUrl, product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
}
