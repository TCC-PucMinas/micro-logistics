package model

import (
	"fmt"
	"micro-logistic/db"
	"micro-logistic/helpers"
)

type Truck struct {
	Id       int64    `json:"id"`
	Brand    string   `json:"brand"`
	Model    string   `json:"model"`
	Year     string   `json:"year"`
	Plate    string   `json:"plate"`
	Carrying Carrying `json:"carrying"`
}

func (truck *Truck) GetById(id int64) error {

	sql := db.ConnectDatabase()

	query := `select id, brand, model, year, plate, id_carry from trucks where id = ? limit 1;`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var brand, model, year, plate string
		var id, idCarry int64
		_ = requestConfig.Scan(&id, &brand, &model, &year, &plate, &idCarry)
		if id != 0 {
			truck.Id = id
			truck.Brand = brand
			truck.Model = model
			truck.Plate = plate
			truck.Year = year
			truck.Carrying.Id = idCarry
		}
	}

	return nil
}

func (truck *Truck) DeleteById() error {
	sql := db.ConnectDatabase()

	query := "delete from trucks where id = ?"

	deleteCarry, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := deleteCarry.Exec(truck.Id)

	if e != nil {
		return e
	}

	return nil
}

func (truck *Truck) UpdateTruck() error {
	sql := db.ConnectDatabase()

	query := "update trucks set brand = ?, model = ?, year = ?, plate = ?, id_carry = ? where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(truck.Brand, truck.Model, truck.Year, truck.Plate, truck.Carrying.Id, truck.Id)

	if e != nil {
		return e
	}

	return nil
}

func (truck *Truck) CreateTruck() error {
	sql := db.ConnectDatabase()

	query := "insert into trucks (brand, model, year, plate, id_carry) values (?, ?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(truck.Brand, truck.Model, truck.Year, truck.Plate, truck.Carrying.Id)

	if e != nil {
		return e
	}

	return nil
}

func (truck *Truck) GetTruckByIdCarryPaginate(plate string, idCarry, page, limit int64) ([]Truck, int64, error) {
	var truckArray []Truck
	var total int64

	sql := db.ConnectDatabase()

	plate = "%" + plate + "%"

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("trucks")

	query := fmt.Sprintf("select id, brand, model, year, plate, id_carry, %v from trucks where id_carry = ? and plate like ? LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, idCarry, plate, paginate.Limit, paginate.Page)

	if err != nil {
		return truckArray, total, err
	}

	for requestConfig.Next() {
		truckGet := Truck{}
		var brand, model, year, plate string
		var id, idCarry int64
		_ = requestConfig.Scan(&id, &brand, &model, &year, &plate, &idCarry, &total)

		if id != 0 {
			truckGet.Id = id
			truckGet.Brand = brand
			truckGet.Model = model
			truckGet.Plate = plate
			truckGet.Year = year
			truckGet.Carrying.Id = idCarry

			truckArray = append(truckArray, truckGet)
		}

	}

	return truckArray, total, nil
}
