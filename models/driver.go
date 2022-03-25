package model

import (
	"fmt"
	"log"
	"micro-logistic/db"
	"micro-logistic/helpers"
)

type Driver struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Carrying Carrying `json:"Carrying"`
	Truck    Truck    `json:"truck"`
}

func (driver *Driver) GetById(id int64) error {

	sql := db.ConnectDatabase()

	query := `select id, name, image, id_carring, id_truck from drivers where id = ? limit 1;`

	log.Println("id", id)
	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var name, image string
		var id, idCarry, idTruck int64
		_ = requestConfig.Scan(&id, &name, &image, &idCarry, &idTruck)
		if id != 0 {
			driver.Id = id
			driver.Name = name
			driver.Image = image
			driver.Carrying.Id = idCarry
			driver.Truck.Id = idTruck
		}
	}

	return nil
}

func (driver *Driver) GetByNameAndIdCarry(name string, idCarry int64) error {

	sql := db.ConnectDatabase()

	query := `select id, name, image, id_carring, id_truck from drivers where name = ? and id_carring = ?  limit 1;`

	requestConfig, err := sql.Query(query, name, idCarry)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var name, image string
		var id, idCarry, idTruck int64
		_ = requestConfig.Scan(&id, &name, &image, &idCarry, &idTruck)
		if id != 0 {
			driver.Id = id
			driver.Name = name
			driver.Image = image
			driver.Carrying.Id = idCarry
			driver.Truck.Id = idTruck
		}
	}

	return nil
}

func (driver *Driver) DeleteById() error {
	sql := db.ConnectDatabase()

	query := "delete from drivers where id = ?"

	deleteCarry, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := deleteCarry.Exec(driver.Id)

	if e != nil {
		return e
	}

	return nil
}

func (driver *Driver) UpdateDriver() error {
	sql := db.ConnectDatabase()

	query := "update drivers set name = ?, image = ?, id_carring = ?, id_truck = ? where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(driver.Name, driver.Image, driver.Carrying.Id, driver.Truck.Id, driver.Id)

	if e != nil {
		return e
	}

	return nil
}

func (driver *Driver) CreateDriver() error {
	sql := db.ConnectDatabase()

	query := "insert into drivers (name, image, id_carring, id_truck) values (?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(driver.Name, driver.Image, driver.Carrying.Id, driver.Truck.Id)

	if e != nil {
		return e
	}

	return nil
}

func (driver *Driver) GetDriverPaginateByNameAndIdCarry(name string, idCarry int64, page, limit int64) ([]Driver, int64, error) {
	var driverArray []Driver
	var total int64

	sql := db.ConnectDatabase()

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	name = "%" + name + "%"
	paginate.PaginateMounted()
	paginate.MountedQuery("drivers")

	query := fmt.Sprintf("select id, name, image, id_carring, id_truck, %v from drivers where name like ? and id_carring = ? LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, name, idCarry, paginate.Limit, paginate.Page)

	if err != nil {
		return driverArray, total, err
	}

	for requestConfig.Next() {
		driverGet := Driver{}
		var name, image string
		var id, idCarry, idTruck int64
		_ = requestConfig.Scan(&id, &name, &image, &idCarry, &idTruck, &total)
		if id != 0 {
			driverGet.Id = id
			driverGet.Name = name
			driverGet.Image = image
			driverGet.Carrying.Id = idCarry
			driverGet.Truck.Id = idTruck
			driverArray = append(driverArray, driverGet)
		}
	}

	return driverArray, total, nil
}
