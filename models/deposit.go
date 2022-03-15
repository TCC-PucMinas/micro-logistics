package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"micro-logistic/helpers"
	"strconv"
	"time"

	"micro-logistic/db"
)

var (
	keyDepositRedisGetById           = "key-deposit-get-by-id"
	keyDepositRedisGetByName         = "key-deposit-get-by-name"
	keyDepositRedisGetPaginateByName = "key-deposit-get-paginate-by-name"
)

type Deposit struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Street   string   `json:"street"`
	District string   `json:"district"`
	ZipCode  string   `json:"zipCode"`
	City     string   `json:"city"`
	Country  string   `json:"country"`
	State    string   `json:"state"`
	Number   string   `json:"number"`
	Lat      string   `json:"lat"`
	Lng      string   `json:"lng"`
	Carrying Carrying `json:"carrying"`
}

func setRedisCacheDepositgGetById(deposit *Deposit) error {
	db, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	json, err := json.Marshal(deposit)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyDepositRedisGetById, json)

	return db.Set(key, json, 1*time.Minute).Err()
}

func getDepositRedisCacheGetOneById(id int64) (Deposit, error) {
	deposit := Deposit{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return deposit, err
	}

	key := fmt.Sprintf("%v - %v", keyDepositRedisGetById, id)

	value, err := redis.Get(key).Result()

	if err != nil {
		return deposit, err
	}

	if err := json.Unmarshal([]byte(value), &deposit); err != nil {
		return deposit, err
	}

	return deposit, nil
}

func getDepositRedisCacheGetOneByNameAndCarry(name string, idCarry int64) (Deposit, error) {
	deposit := Deposit{}

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return deposit, err
	}

	key := fmt.Sprintf("%v - %v - %v", keyDepositRedisGetByName, name, idCarry)

	value, err := redis.Get(key).Result()

	if err != nil {
		return deposit, err
	}

	if err := json.Unmarshal([]byte(value), &deposit); err != nil {
		return deposit, err
	}

	return deposit, nil
}

func setRedisCacheDepositGetByNameAndEmail(deposit *Deposit) error {
	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	marshal, err := json.Marshal(deposit)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v -%v", keyDepositRedisGetByName, deposit.Name, deposit.Carrying.Id)

	return redis.Set(key, marshal, 1*time.Minute).Err()
}

func getDepositRedisCacheGetOneByNamePaginate(name string, page, limit int64) ([]Deposit, error) {
	var deposit []Deposit

	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return deposit, err
	}

	key := fmt.Sprintf("%v - %v -%v -%v", keyDepositRedisGetPaginateByName, name, page, limit)

	value, err := redis.Get(key).Result()

	if err != nil {
		return deposit, err
	}

	if err := json.Unmarshal([]byte(value), &deposit); err != nil {
		return deposit, err
	}

	return deposit, nil
}

func setRedisCacheDepositGetByPaginateByName(name string, page, limit int64, deposit []Deposit) error {
	redis, err := db.ConnectDatabaseRedis()

	if err != nil {
		return err
	}

	marshal, err := json.Marshal(deposit)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%v - %v -%v -%v", keyDepositRedisGetPaginateByName, name, page, limit)

	return redis.Set(key, marshal, 1*time.Minute).Err()
}

func (deposit *Deposit) GetById(id int64) error {

	if c, err := getDepositRedisCacheGetOneById(id); err == nil {
		deposit = &c
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id, name, street, district, city, country, state, number, lat, lng, id_carry  from deposits where id = ? limit 1;`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var name, street, district, city, country, state, number, lat, lng string
		var idCarry, id int64
		_ = requestConfig.Scan(&id, &name, &street, &district, &city, &country, &state, &number, &lat, &lng, &idCarry)
		deposit.Id = id
		deposit.Name = name
		deposit.Street = street
		deposit.District = district
		deposit.City = city
		deposit.Country = country
		deposit.State = state
		deposit.Number = number
		deposit.Lat = lat
		deposit.Lng = lng
		deposit.Carrying.Id = idCarry
	}

	if deposit.Id == 0 {
		return errors.New("Not found key")
	}

	_ = setRedisCacheDepositgGetById(deposit)

	return nil
}

func (deposit *Deposit) GetByNameAndIdCarry(name string, idCarry int64) error {

	if c, err := getDepositRedisCacheGetOneByNameAndCarry(name, idCarry); err == nil {
		deposit = &c
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id, name, street, district, city, country, state, number, lat, lng, id_carry  from deposits where name = ? and id_carry = ? limit 1;`

	requestConfig, err := sql.Query(query, name, idCarry)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, name, street, district, city, country, state, number, lat, lng string
		var idCarry int64
		_ = requestConfig.Scan(&id, &name, &street, &district, &city, &country, &state, &number, &lat, &lng, &idCarry)
		i64, _ := strconv.ParseInt(id, 10, 64)
		deposit.Id = i64
		deposit.Name = name
		deposit.Street = street
		deposit.District = district
		deposit.City = city
		deposit.Country = country
		deposit.State = state
		deposit.Number = number
		deposit.Lat = lat
		deposit.Lng = lng
		deposit.Carrying.Id = idCarry
	}

	_ = setRedisCacheDepositGetByNameAndEmail(deposit)

	return nil
}

func (deposit *Deposit) DeleteById() error {
	sql := db.ConnectDatabase()

	query := "delete from deposits where id = ?"

	deleteDeposit, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := deleteDeposit.Exec(deposit.Id)

	if e != nil {
		return e
	}

	return nil
}

func (deposit *Deposit) UpdateDepositById() error {
	sql := db.ConnectDatabase()

	query := "update deposits set `name` = ?, street = ?, district = ?, city = ?, country = ?, `state` = ?, `number` = ?, lat = ?, lng = ?, id_carry = ? where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(deposit.Name, deposit.Street, deposit.District, deposit.City, deposit.Country, deposit.State, deposit.Number, deposit.Lat, deposit.Lng, deposit.Carrying.Id, deposit.Id)

	if e != nil {
		return e
	}

	return nil
}

func (deposit *Deposit) CreateDeposit() error {
	sql := db.ConnectDatabase()

	query := "insert into deposits (`name`, street, district, city, country, `state`, `number`, lat, lng, id_carry) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := createDestination.Exec(deposit.Name, deposit.Street, deposit.District, deposit.City, deposit.Country, deposit.State, deposit.Number, deposit.Lat, deposit.Lng, deposit.Carrying.Id)

	if e != nil {
		return e
	}

	return nil
}

func (deposit *Deposit) GetDepositByNamePaginate(name string, page, limit int64) ([]Deposit, int64, error) {
	var depositArray []Deposit
	var total int64

	if c, err := getDepositRedisCacheGetOneByNamePaginate(name, page, limit); err == nil {
		depositArray = c
		return depositArray, total, nil
	}

	name = "%" + name + "%"

	sql := db.ConnectDatabase()

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("deposits")

	query := fmt.Sprintf("select id, name, street, district, city, country, state, number, lat, lng, id_carry, %v from deposits where name like ? LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, name, paginate.Limit, paginate.Page)

	if err != nil {
		return depositArray, total, err
	}

	for requestConfig.Next() {
		depositGet := Deposit{}
		var name, street, district, city, country, state, number, lat, lng string
		var id, idCarry int64
		_ = requestConfig.Scan(&id, &name, &street, &district, &city, &country, &state, &number, &lat, &lng, &idCarry, &total)
		if id != 0 {
			depositGet.Id = id
			depositGet.Name = name
			depositGet.Street = street
			depositGet.District = district
			depositGet.City = city
			depositGet.Country = country
			depositGet.State = state
			depositGet.Number = number
			depositGet.Lat = lat
			depositGet.Lng = lng
			depositGet.Carrying.Id = idCarry
		}

		depositArray = append(depositArray, depositGet)
	}

	_ = setRedisCacheDepositGetByPaginateByName(name, page, limit, depositArray)

	return depositArray, total, nil
}
