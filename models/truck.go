package model

type Truck struct {
	Id    int64  `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Plate string `json:"plate"`
}
