package model

import (
	"encoding/json"
	"fmt"
	"micro-logistic/db"
	"micro-logistic/helpers"
	"time"
)

type Driver struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Carrying Carrying `json:"Carrying"`
	Truck    Truck    `json:"truck"`
}

var (
	keyDriverRedisGetById           = "key-driver-get-by-id"
	keyDriverRedisGetPaginateByName = "key-driver-get-paginate-by-name"
)

func setRedisCacheDriverGetById(driver *Driver) error {
	db, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	json, err := json.Marshal(driver)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyTruckRedisGetById, json)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getDriverRedisCacheGetOneById(id int64) (Driver, error) {
	driver := Driver{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return driver, err
	}

	key := fmt.Sprintf("%v - %v", keyDriverRedisGetById, id)

	value, err := redis.Get(key).Result()

	if err != nil {
		return driver, err
	}

	if err := json.Unmarshal([]byte(value), &driver); err != nil {
		return driver, err
	}

	return driver, nil
}

func getDriverRedisCacheGetOneByIdDriverPaginate(name string, page, limit int64) ([]Driver, error) {
	var driver []Driver

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return driver, err
	}

	key := fmt.Sprintf("%v - %v - %v - %v", keyDriverRedisGetPaginateByName, name, page, limit)

	value, err := redis.Get(key).Result()

	if err != nil {
		return driver, err
	}

	if err := json.Unmarshal([]byte(value), &driver); err != nil {
		return driver, err
	}

	return driver, nil
}

func setDriverRedisCacheGetOneByIdDriverPaginate(name string, page, limit int64, driver []Driver) error {
	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	marshal, err := json.Marshal(driver)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%v - %v - %v - %v", keyDriverRedisGetPaginateByName, name, page, limit)

	return redis.Set(key, marshal, 1*time.Hour).Err()
}

func (driver *Driver) GetById(id int64) error {

	if t, err := getDriverRedisCacheGetOneById(id); err == nil {
		driver = &t
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id, name, image, id_carring, id_truck from drivers where id = ? limit 1;`

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

	if driver.Id != 0 {
		_ = setRedisCacheDriverGetById(driver)
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

func (driver *Driver) UpdateTruck() error {
	sql := db.ConnectDatabase()

	query := "update trucks set name = ?, image = ?, id_carring = ?, id_truck = ? where id = ?"

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

func (driver *Driver) CreateTruck() error {
	sql := db.ConnectDatabase()

	query := "insert into trucks (name, image, id_carring, id_truck) values (?, ?, ?, ?)"

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

func (driver *Driver) GetTruckByIdDriverPaginateByName(name string, page, limit int64) ([]Driver, int64, error) {
	var driverArray []Driver
	var total int64

	if c, err := getDriverRedisCacheGetOneByIdDriverPaginate(name, page, limit); err == nil {
		driverArray = c
		return driverArray, total, nil
	}

	sql := db.ConnectDatabase()

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	name = "%" + name + "%"
	paginate.PaginateMounted()
	paginate.MountedQuery("drivers")

	query := fmt.Sprintf("select id, name, image, id_carring, id_truck  %v from drivers where name like ? LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, name, paginate.Limit, paginate.Page)

	if err != nil {
		return driverArray, total, err
	}

	for requestConfig.Next() {
		driverGet := Driver{}
		var name, image string
		var id, idCarry, idTruck int64
		_ = requestConfig.Scan(&id, &name, &image, &idCarry, &idTruck)
		if id != 0 {
			driverGet.Id = id
			driverGet.Name = name
			driverGet.Image = image
			driverGet.Carrying.Id = idCarry
			driverGet.Truck.Id = idTruck
			driverArray = append(driverArray, driverGet)
		}
	}

	if len(driverArray) > 0 {
		_ = setDriverRedisCacheGetOneByIdDriverPaginate(name, page, limit, driverArray)
	}

	return driverArray, total, nil
}
