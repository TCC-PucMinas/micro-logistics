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
	keyClientRedisGetById          = "key-client-get-by-id"
	keyClientRedisGetByNameAndPage = "key-client-get-by-name-and-page"
)

type Client struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func setRedisCacheClientGetById(client *Client) error {
	redis := db.ConnectDatabaseRedis()

	marshal, err := json.Marshal(client)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyClientRedisGetById, marshal)

	return redis.Set(key, marshal, 1*time.Hour).Err()
}

func getClientRedisCacheGetOneById(id string) (Client, error) {
	client := Client{}

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v", keyClientRedisGetById, id)

	value, err := redis.Get(key).Result()

	if err != nil {
		return client, err
	}

	if err := json.Unmarshal([]byte(value), &client); err != nil {
		return client, err
	}

	return client, nil
}

func (client *Client) GetById(idClient string) error {

	if c, err := getClientRedisCacheGetOneById(idClient); err == nil {
		client = &c
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id, name, email, phone from clients where id = ? limit 1;`

	requestConfig, err := sql.Query(query, idClient)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, name, email, phone string
		_ = requestConfig.Scan(&id, &name, &email, &phone)
		i64, _ := strconv.ParseInt(id, 10, 64)
		client.Id = i64
		client.Name = name
		client.Email = email
		client.Phone = phone
	}

	if client.Id == 0 {
		return errors.New("Not found key")
	}

	_ = setRedisCacheClientGetById(client)

	return nil
}

func setRedisCacheClientGetByName(name string, page, limit int64, clients []Client) error {
	redis := db.ConnectDatabaseRedis()

	marshal, err := json.Marshal(clients)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%v - %v -%v -%v", keyClientRedisGetByNameAndPage, name, page, limit)

	return redis.Set(key, marshal, 1*time.Hour).Err()
}

func getClientRedisCacheGetOneByName(name string, page, limit int64) ([]Client, error) {
	var clients []Client

	redis := db.ConnectDatabaseRedis()

	key := fmt.Sprintf("%v - %v -%v -%v", keyClientRedisGetByNameAndPage, name, page, limit)

	value, err := redis.Get(key).Result()

	if err != nil {
		return clients, err
	}

	if err := json.Unmarshal([]byte(value), &clients); err != nil {
		return clients, err
	}

	return clients, nil
}

func (client *Client) GetByNameLike(name string, page, limit int64) ([]Client, int64, error) {
	var clientArray []Client
	var total int64

	if c, err := getClientRedisCacheGetOneByName(name, page, limit); err == nil {
		clientArray = c
		return clientArray, total, nil
	}

	sql := db.ConnectDatabase()

	name = "%" + name + "%"

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("clients")

	query := fmt.Sprintf("select id, name, email, phone, %v from clients where name like ? LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, name, paginate.Limit, paginate.Page)

	if err != nil {
		return clientArray, total, err
	}

	for requestConfig.Next() {
		var clientGet Client
		var id, name, email, phone string
		err = requestConfig.Scan(&id, &name, &email, &phone, &total)
		i64, _ := strconv.ParseInt(id, 10, 64)
		clientGet.Id = i64
		clientGet.Name = name
		clientGet.Email = email
		clientGet.Phone = phone

		clientArray = append(clientArray, clientGet)
	}

	_ = setRedisCacheClientGetByName(name, page, limit, clientArray)

	return clientArray, total, nil
}
