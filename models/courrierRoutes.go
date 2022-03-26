package model

import (
	"encoding/json"
	"micro-logistic/db"
)

type LatAndLng struct {
	Lat string `json:"Lat"`
	Lng string `json:"lng`
}

func (l *LatAndLng) StructToString() string {
	bytes, _ := json.Marshal(l)
	return string(bytes)
}

func (l *LatAndLng) StringToStruct(val string) {
	_ = json.Unmarshal([]byte(val), l)
}

type CourierRoute struct {
	Id        int64     `json:"id"`
	Courier   Courier   `json:"courier"`
	Order     int64     `json:"order"`
	LatInit   LatAndLng `json:"lat_init"`
	LatFinish LatAndLng `json:"lat_finish"`
}

func (courierRoute *CourierRoute) CreateCourierRoute() error {
	sql := db.ConnectDatabase()

	query := "insert into courier_routes (id_courier, `order`, `latInit`, `latFinish`) values (?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(courierRoute.Courier.Id, courierRoute.Order, courierRoute.LatInit.StructToString(), courierRoute.LatFinish.StructToString())

	if e != nil {
		return e
	}

	return nil
}

func (courierRoute *CourierRoute) GetCourierRoutes() ([]CourierRoute, error) {
	var courierRoutesArray []CourierRoute

	sql := db.ConnectDatabase()
	query := "select id, id_courier, `order`, `latInit`, `latFinish` from courier_routes"

	requestConfig, err := sql.Query(query)

	if err != nil {
		return courierRoutesArray, err
	}

	for requestConfig.Next() {
		courierRouteGet := CourierRoute{}
		var latInit, latFinish string
		var IdCourier, order, id int64
		_ = requestConfig.Scan(&id, &IdCourier, &order, &latInit, &latFinish)
		if id != 0 {
			courierRouteGet.Id = id
			courierRouteGet.LatInit.StringToStruct(latInit)
			courierRouteGet.LatFinish.StringToStruct(latFinish)
			courierRouteGet.Order = order
			courierRouteGet.Courier.Id = IdCourier
			courierRoutesArray = append(courierRoutesArray, courierRouteGet)
		}
	}
	return courierRoutesArray, nil
}
