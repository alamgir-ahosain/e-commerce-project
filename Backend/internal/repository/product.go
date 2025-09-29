package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
)

type ProductRepo interface {
	CreateProductFunc(product models.Product) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
	GetByIdFunc(id int) (*models.Product, error)
	UpdateProductByIdFunc(product models.Product, id int) (*models.Product, error)
	//UpdateProductByIdPatchFunc(patchData map[string]interface{}, id int) (*models.Product, error)
	DeleteProductByIdFunc(id int) error
}

type productRepo struct {
	db *sqlx.DB
}

// Public constructor
func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{db: db}
	return repo
}

// CreateProductFunc implements ProductRepo.
func (pr *productRepo) CreateProductFunc(p models.Product) (*models.Product, error) {
	query := `INSERT INTO products (title, description, price, image_url)
	          values($1,$2,$3,$4) RETURNING id`
	//Executed Query
	row := pr.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	// Scan the returned id into p.ID
	err := row.Scan(&p.ID)
	if err != nil {

		fmt.Println("rows error:", err)
		return nil, err
	}
	return &p, nil

}

// GetFunc implements ProductRepo.
func (pr *productRepo) GetAllProduct() (*[]models.Product, error) {
	var productList []models.Product

	query := `SELECT id,title,description,price,image_url FROM products`
	err := pr.db.Select(&productList, query)
	if err != nil {
		//fmt.Println("rows error:", err)
		return nil, err
	}

	//fmt.Println(productList)
	return &productList, nil
}

// GetByIdFunc implements ProductRepo.
func (pr *productRepo) GetByIdFunc(id int) (*models.Product, error) {

	var prd models.Product
	query := `SELECT id, title,description,price,image_url FROM products where id=$1`
	err := pr.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows Found:", err)
			return nil, nil

		}
	}
	return &prd, err

}

// UpdateProductByIdPutFunc implements ProductRepo.
func (pr *productRepo) UpdateProductByIdFunc(p models.Product, id int) (*models.Product, error) {

	query := `UPDATE products SET title=$1, description=$2, price=$3, image_url=$4 WHERE id=$5 RETURNING id`
	row := pr.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, id)
	err := row.Scan(&p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with id=%d not found", id)
		}
		return nil, err
	}
	return &p, nil

}

// UpdateProductByIdPatchFunc implements ProductRepo.
// func (pr *productRepo) UpdateProductByIdPatchFunc(patchData map[string]interface{}, id int) (*models.Product, error) {
// 	for i, val := range pr.productList {
// 		if val.ID == id {
// 			// Update fields if present
// 			if title, ok := patchData["title"].(string); ok {
// 				pr.productList[i].Title = title
// 			}
// 			if description, ok := patchData["description"].(string); ok {
// 				pr.productList[i].Description = description
// 			}
// 			if price, ok := patchData["price"].(float64); ok {
// 				pr.productList[i].Price = price
// 			}
// 			if imageUrl, ok := patchData["imageUrl"].(string); ok {
// 				pr.productList[i].ImgUrl = imageUrl
// 			}

// 			return &pr.productList[i], nil
// 		}
// 	}
// 	return &models.Product{}, fmt.Errorf("product with id=%d not found", id)

// }

// DeleteProductByIdFunc implements ProductRepo.
func (pr *productRepo) DeleteProductByIdFunc(id int) error {

	q := `DELETE FROM products where id=$1`
	result, err := pr.db.Exec(q, id)
	
	if err != nil {return err}
	cnt, err := result.RowsAffected()
	if err != nil {return err}
	if cnt == 0 {	return fmt.Errorf("product with id=%d not found", id)}
	return nil

}
