package model

type Destination struct {
	Id       int64  `json:"id"`
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Number   string `json:"number"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
	Client   Client `json:"client"`
}
