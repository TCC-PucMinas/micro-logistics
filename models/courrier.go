package model

type Courier struct {
	Id     int64  `json:"id"`
	Driver Driver `json:"driver"`
}
