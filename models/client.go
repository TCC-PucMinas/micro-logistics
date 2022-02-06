package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/TCC-PucMinas/micro-logistics/db"
)

var keyClientRedisGetById = "key-client-get-by-id"

type Client struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func setRedisCacheClientGetById(client *Client) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(client)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyClientRedisGetById, json)

	return db.Set(key, json, 1*time.Hour).Err()
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
