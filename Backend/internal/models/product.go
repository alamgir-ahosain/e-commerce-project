package models

type Product struct {
	ID          int     `json:"id"` //change ID to id in json format
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}
