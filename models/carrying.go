package model

import (
	"fmt"
	"micro-logistic/helpers"
	"strconv"

	"micro-logistic/db"
)

type Carrying struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	District string `json:"district"`
	ZipCode  string `json:"zipCode"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Number   string `json:"number"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
}

func (carrying *Carrying) GetById(id int64) error {

	sql := db.ConnectDatabase()

	query := `select id, name, street, district, city, country, state, number, lat, lng, zipCode  from carryings where id = ? limit 1;`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, name, street, district, city, country, state, number, lat, lng, zipCode string
		_ = requestConfig.Scan(&id, &name, &street, &district, &city, &country, &state, &number, &lat, &lng, &zipCode)
		i64, _ := strconv.ParseInt(id, 10, 64)
		carrying.Id = i64
		carrying.Name = name
		carrying.Street = street
		carrying.District = district
		carrying.City = city
		carrying.ZipCode = zipCode
		carrying.Country = country
		carrying.State = state
		carrying.Number = number
		carrying.Lat = lat
		carrying.Lng = lng
	}

	return nil
}

func (carrying *Carrying) GetByName(name string) error {

	sql := db.ConnectDatabase()

	query := `select id, name, street, district, city, country, state, number, zipCode, lat, lng  from carryings where name = ? limit 1;`

	requestConfig, err := sql.Query(query, name)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, name, street, district, city, country, state, number, lat, lng, zipCode string
		_ = requestConfig.Scan(&id, &name, &street, &district, &city, &country, &state, &number, &zipCode, &lat, &lng)
		i64, _ := strconv.ParseInt(id, 10, 64)
		carrying.Id = i64
		carrying.Name = name
		carrying.Street = street
		carrying.District = district
		carrying.City = city
		carrying.ZipCode = zipCode
		carrying.Country = country
		carrying.State = state
		carrying.Number = number
		carrying.Lat = lat
		carrying.Lng = lng
	}

	return nil
}

func (carrying *Carrying) DeleteById() error {
	sql := db.ConnectDatabase()

	query := "delete from carryings where id = ?"

	deleteCarry, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := deleteCarry.Exec(carrying.Id)

	if e != nil {
		return e
	}

	return nil
}

func (carrying *Carrying) UpdateCarryById() error {
	sql := db.ConnectDatabase()

	query := "update carryings set name = ?, street = ?, district = ?, city = ?, country = ?, `state` = ?, `number` = ?, zipCode = ?, lat = ?, lng = ? where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(carrying.Name, carrying.Street, carrying.District, carrying.City, carrying.Country, carrying.State, carrying.Number, carrying.ZipCode, carrying.Lat, carrying.Lng, carrying.Id)

	if e != nil {
		return e
	}

	return nil
}

func (carrying *Carrying) CreateCarry() error {
	sql := db.ConnectDatabase()

	query := "insert into carryings (name, street, district, city, country, `state`, `number`, zipCode, lat, lng) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(carrying.Name, carrying.Street, carrying.District, carrying.City, carrying.Country, carrying.State, carrying.Number, carrying.ZipCode, carrying.Lat, carrying.Lng)

	if e != nil {
		return e
	}

	return nil
}

func (carry *Carrying) GetCarryByNamePaginate(name string, page, limit int64) ([]Carrying, int64, error) {
	var carryArray []Carrying
	var total int64

	name = "%" + name + "%"

	sql := db.ConnectDatabase()

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("carryings")

	query := fmt.Sprintf("select id, name, street, district, city, country, state, number, zipCode, lat, lng , %v from carryings where name like ? LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, name, paginate.Limit, paginate.Page)

	if err != nil {
		return carryArray, total, err
	}

	for requestConfig.Next() {
		carryGet := Carrying{}
		var name, street, zipCode, district, city, country, state, number, lat, lng string
		var id int64
		_ = requestConfig.Scan(&id, &name, &street, &district, &city, &country, &state, &number, &zipCode, &lat, &lng, &total)

		if id != 0 {
			carryGet.Id = id
			carryGet.Name = name
			carryGet.Street = street
			carryGet.District = district
			carryGet.City = city
			carryGet.Country = country
			carryGet.ZipCode = zipCode
			carryGet.State = state
			carryGet.Number = number
			carryGet.Lat = lat
			carryGet.Lng = lng

			carryArray = append(carryArray, carryGet)
		}

	}

	return carryArray, total, nil
}
