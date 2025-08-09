package services

import (
	"encoding/json"
	"net/http"
	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
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
