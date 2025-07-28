package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* ___status code
   200 : ok
   201 : create server resource
   400 : bad request
   404 : not found
   500 : internal server error
*/

type Product struct {
	ID          int     `json:"id"` //change ID to id in json format
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product //list of prduct

// GET->only header
func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*") //set header: (*) anyone can access the responce
	w.Header().Set("Content-Type", "application/json") //body :resource media type

	if r.Method != http.MethodGet { // if r.Method=post,put,patch,delete
		http.Error(w, "only GET request allowed!", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList) //encoded json format

}

func main() {

	mux := http.NewServeMux() //router

	
	mux.HandleFunc("/products", getProducts) //route


	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}
func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "orange is a fruit.",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQEZeYvkPorIuW7yfCQQ-cM_I0L0UbP0gOMyA&s",
	}
	pd2 := Product{
		ID:          2,
		Title:       "Mango",
		Description: "Mangifera indica, commonly known as Mango,.",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTFT-bteAg0wbO0yBMfyM8fLq0vG5At3wwLtQ&s",
	}
	pd3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "Apple is a green.",
		Price:       150,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJJfODaTyBw4581VyPy5wQHvq4yfAIzGRHVA&s",
	}
		pd4 := Product{
		ID:          4,
		Title:       "Jack Fruit",
		Description: "National Fruit.",
		Price:       250,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQe3E9guFDcTsvr64CfPpM6pYbssqXWmbiZ6w&s",
	}

	productList = append(productList, pd1)
	productList = append(productList, pd2)
	productList = append(productList, pd3)
	productList = append(productList, pd4)

}
