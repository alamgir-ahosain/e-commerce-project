package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

var productList = []models.Product{
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

// Handle CORS error
func HandleCORSFunc(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                  //set header: (*) anyone can access th
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Alamgir-CustomHeader") //allow content type
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS") //allow content type
	w.Header().Set("Content-Type", "application/json")
}

// make JSON format
func MakeJSONFormatFunc(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// make JSON format
func MakeJSONFormatThreeFunc(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// create product function:post request
func CreateProductFunc(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give  the valid json format", 400)
		return
	}

	productList = append(productList, newProduct)
	MakeJSONFormatFunc(w, 201)
}

// GET id
func GetID(w http.ResponseWriter, r *http.Request) (int, error) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.SendError(w, r, http.StatusBadRequest, "Invalid ID Type")
		return 0, err
	}

	return id, nil
}

// Get prdduct by id
func GetProductByIdFunc(id int) (models.Product, error) {
	for i, val := range productList {
		if val.ID == id {
			return productList[i], nil
		}

	}
	return models.Product{}, fmt.Errorf("produnct with id=%d not found", id)

}

// delete  product by ID
func DeleteProductByIdFunc(id int) (models.Product, error) {
	for i, val := range productList {
		if val.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			return productList[i], nil
		}
	}
	return models.Product{}, fmt.Errorf("error delete product ")
}

// update product by Id
func UpdateProductByIdFunc(w http.ResponseWriter, r *http.Request, id int) {

	var newProduct models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give  the valid json format", 400)
		return
	}
	for i, val := range productList {
		if val.ID == id {
			newProduct.ID = id
			productList[i] = newProduct
			MakeJSONFormatThreeFunc(w, 201, newProduct)
			return
		}
	}

}

// PATCH request to update product by ID (partial update)
func UpdateProductByIdPutFunc(w http.ResponseWriter, r *http.Request, id int) {
	var patchData map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&patchData)
	if err != nil {
		http.Error(w, "Give the valid JSON format", http.StatusBadRequest)
		return
	}

	for i, val := range productList {
		if val.ID == id {
			// Update fields if present
			if title, ok := patchData["title"].(string); ok {
				productList[i].Title = title
			}
			if description, ok := patchData["description"].(string); ok {
				productList[i].Description = description
			}
			if price, ok := patchData["price"].(float64); ok {
				productList[i].Price = price
			}
			if imageUrl, ok := patchData["imageUrl"].(string); ok {
				productList[i].ImgUrl = imageUrl
			}

			// Send updated product as JSON response
			MakeJSONFormatThreeFunc(w, 200, patchData)
			return
		}
	}

	// If product with given ID is not found
	http.Error(w, "Product not found", http.StatusNotFound)
}
