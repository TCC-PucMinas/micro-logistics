package model

type Product struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Client Client `json:"Client"`
}
