package model

import (
	"encoding/json"
	"fmt"
	"micro-logistic/helpers"
	"strconv"
	"time"

	"micro-logistic/db"
)

var (
	keyCarryingRedisGetById           = "key-carrying-get-by-id"
	keyCarryingRedisGetByName         = "key-carrying-get-by-name"
	keyCarryingRedisGetPaginateByName = "key-carrying-get-paginate-by-name"
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

func setRedisCacheCarryingGetById(carry *Carrying) error {
	db, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	json, err := json.Marshal(carry)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyCarryingRedisGetById, json)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getCarryingRedisCacheGetOneById(id int64) (Carrying, error) {
	carrying := Carrying{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return carrying, err
	}

	key := fmt.Sprintf("%v - %v", keyCarryingRedisGetById, id)

	value, err := redis.Get(key).Result()

	if err != nil {
		return carrying, err
	}

	if err := json.Unmarshal([]byte(value), &carrying); err != nil {
		return carrying, err
	}

	return carrying, nil
}

func getCarryRedisCacheGetOneByName(name string) (Carrying, error) {
	carry := Carrying{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return carry, err
	}

	key := fmt.Sprintf("%v - %v", keyCarryingRedisGetByName, name)

	value, err := redis.Get(key).Result()

	if err != nil {
		return carry, err
	}

	if err := json.Unmarshal([]byte(value), &carry); err != nil {
		return carry, err
	}

	return carry, nil
}

func setRedisCacheCarryGetByNameAndEmail(carry *Carrying) error {
	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	marshal, err := json.Marshal(carry)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyCarryingRedisGetByName, carry.Name)

	return redis.Set(key, marshal, 1*time.Hour).Err()
}

func getCarryRedisCacheGetOneByNamePaginate(name string, page, limit int64) ([]Carrying, error) {
	var carry []Carrying

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return carry, err
	}

	key := fmt.Sprintf("%v - %v -%v -%v", keyCarryingRedisGetPaginateByName, name, page, limit)

	value, err := redis.Get(key).Result()

	if err != nil {
		return carry, err
	}

	if err := json.Unmarshal([]byte(value), &carry); err != nil {
		return carry, err
	}

	return carry, nil
}

func setRedisCacheCarryGetByPaginateByName(name string, page, limit int64, carry []Carrying) error {
	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	marshal, err := json.Marshal(carry)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%v - %v -%v -%v", keyCarryingRedisGetPaginateByName, name, page, limit)

	return redis.Set(key, marshal, 1*time.Hour).Err()
}

func (carrying *Carrying) GetById(id int64) error {

	if c, err := getCarryingRedisCacheGetOneById(id); err == nil {
		carrying.City = c.City
		carrying.Country = c.Country
		carrying.District = c.District
		carrying.Id = c.Id
		carrying.Lat = c.Lat
		carrying.Lng = c.Lng
		carrying.Name = c.Name
		carrying.Number = c.Number
		carrying.State = c.State
		carrying.Street = c.Street
		carrying.ZipCode = c.ZipCode
		return nil
	}

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

	_ = setRedisCacheCarryingGetById(carrying)

	return nil
}

func (carrying *Carrying) GetByName(name string) error {

	if c, err := getCarryRedisCacheGetOneByName(carrying.Name); err == nil {
		carrying = &c
		return nil
	}

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

	_ = setRedisCacheCarryGetByNameAndEmail(carrying)

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

	if c, err := getCarryRedisCacheGetOneByNamePaginate(name, page, limit); err == nil {
		carryArray = c
		return carryArray, total, nil
	}

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

	_ = setRedisCacheCarryGetByPaginateByName(name, page, limit, carryArray)

	return carryArray, total, nil
}
