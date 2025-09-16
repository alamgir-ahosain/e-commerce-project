package repository

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
)

func initialProduct(r *productRepo) {
	r.productList = []models.Product{
		{
			ID:          1,
			Title:       "Orange",
			Description: "orange is a fruit.",
			Price:       100,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQEZeYvkPorIuW7yfCQQ-cM_I0L0UbP0gOMyA&s",
		},
		{
			ID:          2,
			Title:       "Mango",
			Description: "Mangifera indica, commonly known as Mango.",
			Price:       200,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTFT-bteAg0wbO0yBMfyM8fLq0vG5At3wwLtQ&s",
		},
		{
			ID:          3,
			Title:       "Apple",
			Description: "Apple is a green.",
			Price:       150,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJJfODaTyBw4581VyPy5wQHvq4yfAIzGRHVA&s",
		},
		{
			ID:          4,
			Title:       "Jack Fruit",
			Description: "National Fruit.",
			Price:       250,
			ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQe3E9guFDcTsvr64CfPpM6pYbssqXWmbiZ6w&s",
		},
	}
}

type ProductRepo interface {
	CreateProductFunc(product models.Product) (*models.Product, error)
	GetProductByIdFunc() (*[]models.Product, error)
	GetByIdFunc(id int) (*models.Product, error)
	UpdateProductByIdFunc(product models.Product, id int) (*models.Product, error)
	UpdateProductByIdPatchFunc(patchData map[string]interface{}, id int) (*models.Product, error)
	DeleteProductByIdFunc(id int) (*models.Product, error)
}

type productRepo struct {
	productList []models.Product
}

// Public constructor
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	initialProduct(repo)
	return repo
}

// CreateProductFunc implements ProductRepo.
func (pr *productRepo) CreateProductFunc(p models.Product) (*models.Product, error) {

	p.ID = len(pr.productList) + 1
	pr.productList = append(pr.productList, p)
	return &p, nil

}

// GetFunc implements ProductRepo.
func (pr *productRepo) GetProductByIdFunc() (*[]models.Product, error) {
	return &pr.productList, nil
}

// GetByIdFunc implements ProductRepo.
func (pr *productRepo) GetByIdFunc(id int) (*models.Product, error) {
	for i, val := range pr.productList {
		if val.ID == id {
			return &pr.productList[i], nil
		}

	}
	return &models.Product{}, fmt.Errorf("produnct with id=%d not found", id)

}

// UpdateProductByIdPutFunc implements ProductRepo.
func (pr *productRepo) UpdateProductByIdFunc(p models.Product, id int) (*models.Product, error) {
	for i, val := range pr.productList {
		if val.ID == id {
			p.ID = id
			pr.productList[i] = p
			return &pr.productList[i], nil
		}
	}
	return &models.Product{}, fmt.Errorf("product with id=%d not found", id)
}

// UpdateProductByIdPatchFunc implements ProductRepo.
func (pr *productRepo) UpdateProductByIdPatchFunc(patchData map[string]interface{}, id int) (*models.Product, error) {
	for i, val := range pr.productList {
		if val.ID == id {
			// Update fields if present
			if title, ok := patchData["title"].(string); ok {
				pr.productList[i].Title = title
			}
			if description, ok := patchData["description"].(string); ok {
				pr.productList[i].Description = description
			}
			if price, ok := patchData["price"].(float64); ok {
				pr.productList[i].Price = price
			}
			if imageUrl, ok := patchData["imageUrl"].(string); ok {
				pr.productList[i].ImgUrl = imageUrl
			}

			return &pr.productList[i], nil
		}
	}
	return &models.Product{}, fmt.Errorf("product with id=%d not found", id)

}

// DeleteProductByIdFunc implements ProductRepo.
func (pr *productRepo) DeleteProductByIdFunc(id int) (*models.Product, error) {
	for i, val := range pr.productList {
		if val.ID == id {
			deletedProduct := pr.productList[i]
			pr.productList = append(pr.productList[:i], pr.productList[i+1:]...)
			return &deletedProduct, nil
		}
	}
	return &models.Product{}, fmt.Errorf("error deleting product with id=%d", id)
}
