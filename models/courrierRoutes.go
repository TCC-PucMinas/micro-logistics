package model

import (
	"encoding/json"
	"fmt"
	"micro-logistic/db"
	"micro-logistic/helpers"
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
	Id        int64   `json:"id"`
	Courier   Courier `json:"courier"`
	Order     int64   `json:"order"`
	Driver    Driver
	LatInit   LatAndLng `json:"lat_init"`
	LatFinish LatAndLng `json:"lat_finish"`
}

func (courierRoute *CourierRoute) CreateCourierRoute() error {
	sql := db.ConnectDatabase()

	query := "insert into courier_routes (id_courier, `order`, `latInit`, `latFinish`, id_driver) values (?, ?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(courierRoute.Courier.Id, courierRoute.Order, courierRoute.LatInit.StructToString(), courierRoute.LatFinish.StructToString(), courierRoute.Driver.Id)

	if e != nil {
		return e
	}

	return nil
}

func (courierRoute *CourierRoute) GetCourierRouteById(id int64) error {

	sql := db.ConnectDatabase()
	query := "select id, id_courier, `order`, `latInit`, `latFinish`, id_driver from courier_routes where id = ?"

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var latInit, latFinish string
		var IdCourier, order, id, idDriver int64
		_ = requestConfig.Scan(&id, &IdCourier, &order, &latInit, &latFinish, &idDriver)
		if id != 0 {
			courierRoute.Id = id
			courierRoute.LatInit.StringToStruct(latInit)
			courierRoute.LatFinish.StringToStruct(latFinish)
			courierRoute.Order = order
			courierRoute.Driver.Id = idDriver
			courierRoute.Courier.Id = IdCourier
		}
	}
	return nil
}

func (courierRoute *CourierRoute) GetCourierRouteByIdCourier() error {

	sql := db.ConnectDatabase()
	query := "select id, id_courier, `order`, `latInit`, `latFinish`, id_driver from courier_routes where id_courier = ?"

	requestConfig, err := sql.Query(query, courierRoute.Courier.Id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var latInit, latFinish string
		var IdCourier, order, id, idDriver int64
		_ = requestConfig.Scan(&id, &IdCourier, &order, &latInit, &latFinish, &idDriver)
		if id != 0 {
			courierRoute.Id = id
			courierRoute.LatInit.StringToStruct(latInit)
			courierRoute.LatFinish.StringToStruct(latFinish)
			courierRoute.Order = order
			courierRoute.Driver.Id = idDriver
			courierRoute.Courier.Id = IdCourier
		}
	}
	return nil
}

func (courierRoute *CourierRoute) GetCourierRoutes(idDriver int64) ([]CourierRoute, error) {
	var courierRoutesArray []CourierRoute

	sql := db.ConnectDatabase()
	query := "select id, id_courier, `order`, `latInit`, `latFinish` from courier_routes where id_driver = ?"

	requestConfig, err := sql.Query(query, idDriver)

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

func (courierRoute *CourierRoute) GetCourierRoutesPaginate(delivered bool, idDriver, page, limit int64) ([]CourierRoute, int64, error) {
	var courierRoutesArray []CourierRoute
	var total int64

	sql := db.ConnectDatabase()

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("courier_routes")

	query := fmt.Sprintf("select cr.id, cr.id_courier, cr.`order`, cr.`latInit`, cr.`latFinish`, id_driver, %v from courier_routes cr inner join couriers c on cr.id_courier = c.id where c.delivered = ? and id_driver = ? order by `order` asc  LIMIT ? OFFSET ?", paginate.Query)

	requestConfig, err := sql.Query(query, delivered, idDriver, paginate.Limit, paginate.Page)

	if err != nil {
		return courierRoutesArray, total, err
	}

	for requestConfig.Next() {
		courierRouteGet := CourierRoute{}
		var latInit, latFinish string
		var IdCourier, order, id, idDriver int64
		_ = requestConfig.Scan(&id, &IdCourier, &order, &latInit, &latFinish, &idDriver, &total)
		if id != 0 {
			courierRouteGet.Id = id
			courierRouteGet.LatInit.StringToStruct(latInit)
			courierRouteGet.LatFinish.StringToStruct(latFinish)
			courierRouteGet.Order = order
			courierRouteGet.Courier.Id = IdCourier
			courierRoutesArray = append(courierRoutesArray, courierRouteGet)
		}
	}

	return courierRoutesArray, total, nil
}

func (courierRoute *CourierRoute) UpdateByCourierId() error {
	sql := db.ConnectDatabase()

	query := "update courier_routes set `order` = ?, `latInit` = ?, `latFinish`= ? where id_courier = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(courierRoute.Order, courierRoute.LatInit.StructToString(), courierRoute.LatFinish.StructToString(), courierRoute.Courier.Id)

	if e != nil {
		return e
	}

	return nil
}

func (courierRoute *CourierRoute) DeleteCourierRouteById(id int64) error {
	sql := db.ConnectDatabase()

	query := "delete from courier_routes where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(id)

	if e != nil {
		return e
	}

	return nil
}
