package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/TCC-PucMinas/micro-logistics/db"
)

type Destination struct {
	Id       int64  `json:"id"`
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Number   string `json:"number"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
	Client   Client `json:"client"`
}

var keyDestinationRedisGetOneByClientId = "key-client-get-by-client-id"

func setRedisCacheDestinationGetByClientId(destination *Destination) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(destination)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyDestinationRedisGetOneByClientId, json)

	return db.Set(key, json, 1*time.Hour).Err()
}

func getRedisCacheDestinationGetByClientId(id string) (Destination, error) {
	destination := Destination{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v", keyDestinationRedisGetOneByClientId, id)

	value, err := redis.Get(key).Result()

	if err != nil {
		return destination, err
	}

	if err := json.Unmarshal([]byte(value), &destination); err != nil {
		return destination, err
	}

	return destination, nil
}

func (destination *Destination) DestinationGetByClientId(idClient string) error {

	if c, err := getRedisCacheDestinationGetByClientId(idClient); err == nil {
		destination = &c
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id, street, district, city, country, state, number, lat, lng from destinations where id_client = ? limit 1;`

	requestConfig, err := sql.Query(query, idClient)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, street, district, city, country, state, number, lat, lng string
		_ = requestConfig.Scan(&id, &street, &district, &city, &country, &state, &number, &lat, &lng)
		i64, _ := strconv.ParseInt(id, 10, 64)
		destination.Id = i64
		destination.Street = street
		destination.District = district
		destination.City = city
		destination.Country = country
		destination.State = state
		destination.Number = number
		destination.Lat = lat
		destination.Lng = lng
	}

	if destination.Id == 0 {
		return errors.New("Not found key")
	}

	_ = setRedisCacheDestinationGetByClientId(destination)

	return nil
}
