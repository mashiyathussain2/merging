package model

type Product struct {
	ProductName string `json:"product_name"`
	ProductCode string `json:"product_code"`
	Price       int    `json:"price"`
}

type Products []Product
