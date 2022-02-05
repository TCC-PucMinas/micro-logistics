package model


var keyCarryingRedisGetById = "key-carrying-get-by-id"

type Carrying struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}


func setRedisCacheCarryingGetById(client Client) error {
	db := db.ConnectDatabaseRedis()

	json, err := json.Marshal(client)

	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v - %v", keyCarryingRedisGetById, json)

	return db.Set(key, json, 1*time.Hour).Err()
}


func getRedisCacheGetOneById(id string) (Carrying, error) {
	carrying := Carrying{}

	redis := db.ConnectDatabaseRedis()

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



func (carrying *Carrying) GetById(id string) error {

	if c, err := getRedisCacheGetOneById(id); err == nil {
		carrying = &c
		return nil
	}

	sql := db.ConnectDatabase()

	query := `select id, name, lat, lng from carryings where id = ? limit 1;`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var id, name, lat, lng, string
		_ = requestConfig.Scan(&id, &name, &lat, &lng)
		i64, _ := strconv.ParseInt(id, 10, 64)
		carrying.Id = i64
		carrying.Name = name
		carrying.Lat = lat
		carrying.Lng = lng
	}

	if carrying.Id == 0 {
		return uerrors.New("Not found key")
	}

	_ = setRedisCacheCarryingGetById(carrying)

	return nil
}
