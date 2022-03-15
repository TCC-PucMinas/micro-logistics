package model

type Driver struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Carrying Carrying `json:"Carrying"`
	Truck    Truck    `json:"truck"`
}

