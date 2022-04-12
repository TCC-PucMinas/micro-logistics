package model

import (
	"encoding/json"
	"fmt"
	"micro-logistic/db"
	"micro-logistic/helpers"
)

type Documentation struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (l *Documentation) StructToString() string {
	bytes, _ := json.Marshal(l)
	return string(bytes)
}

func (l *Documentation) StringToStruct(val string) {
	_ = json.Unmarshal([]byte(val), l)
}

type Courier struct {
	Id        int64         `json:"id"`
	Deposit   Deposit       `json:"deposit"`
	Client    Client        `json:"client"`
	Product   Product       `json:"product"`
	Delivered bool          `json:"delivered"`
	Doc       Documentation `json:"doc"`
}

func (c *Courier) GetOneByIdDriverAndIdProduct(idProduct int64) error {

	sql := db.ConnectDatabase()

	query := `select id, id_product from couriers where id_product = ? limit 1;`

	requestConfig, err := sql.Query(query, idProduct)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var idProduct, id int64
		_ = requestConfig.Scan(&id, &idProduct)
		if id != 0 {
			c.Id = id
			c.Product.Id = idProduct
		}
	}

	return nil
}

func (c *Courier) GetById(id int64) error {

	sql := db.ConnectDatabase()

	query := `select id, id_deposit, id_client, id_product, delivered, doc from couriers where id = ? limit 1;`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var delivered bool
		var doc string
		var idProduct, idClient, idDeposit, id int64
		_ = requestConfig.Scan(&id, &idDeposit, &idClient, &idProduct, &delivered, &doc)
		if id != 0 {
			c.Id = id
			c.Client.Id = idClient
			c.Deposit.Id = idDeposit
			c.Product.Id = idProduct
			c.Delivered = delivered
			c.Doc.StringToStruct(doc)
		}
	}

	return nil
}

func (c *Courier) DeleteById() error {
	sql := db.ConnectDatabase()

	query := "delete from couriers where id = ?"

	deleteDeposit, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := deleteDeposit.Exec(c.Id)

	if e != nil {
		return e
	}

	return nil
}

func (c *Courier) UpdateById() error {
	sql := db.ConnectDatabase()

	query := "update couriers set id_deposit = ?, id_client = ?, id_product = ?, delivered = ?, doc = ? where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(c.Deposit.Id, c.Client.Id, c.Product.Id, c.Delivered, c.Doc.StructToString(), c.Id)

	if e != nil {
		return e
	}

	return nil
}

func (c *Courier) Create() (int64, error) {
	sql := db.ConnectDatabase()

	query := "insert into couriers (id, id_deposit, id_client, id_product) values (?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return 0, err
	}

	result, e := createDestination.Exec(c.Id, c.Deposit.Id, c.Client.Id, c.Product.Id)

	if e != nil {
		return 0, e
	}

	return result.LastInsertId()
}

func (c *Courier) GetCouriers() ([]Courier, error) {
	var courierArray []Courier

	sql := db.ConnectDatabase()
	query := "select id, id_deposit, id_client, id_product from couriers"

	requestConfig, err := sql.Query(query)

	if err != nil {
		return courierArray, err
	}

	for requestConfig.Next() {
		courierGet := Courier{}
		var idProduct, idClient, idDeposit, id int64
		_ = requestConfig.Scan(&id, &idDeposit, &idClient, &idProduct)
		if id != 0 {
			courierGet.Id = id
			courierGet.Product.Id = idProduct
			courierGet.Deposit.Id = idDeposit
			courierGet.Client.Id = idClient
			courierArray = append(courierArray, courierGet)
		}
	}

	return courierArray, nil
}

func (c *Courier) GetCourierPaginate(page, limit int64) ([]Courier, int64, error) {
	var courierArray []Courier
	var total int64

	sql := db.ConnectDatabase()

	paginate := helpers.Paginate{
		Page:  page,
		Limit: limit,
	}

	paginate.PaginateMounted()
	paginate.MountedQuery("deposits")

	query := fmt.Sprintf(`select id, id_deposit, id_client, id_product, %v  from couriers c LIMIT ? OFFSET ?;`, paginate.Query)

	requestConfig, err := sql.Query(query, paginate.Limit, paginate.Page)

	if err != nil {
		return courierArray, total, err
	}

	for requestConfig.Next() {
		courierGet := Courier{}
		var idProduct, idClient, idDeposit, id int64
		_ = requestConfig.Scan(&id, &idDeposit, &idClient, &idProduct, &total)
		if id != 0 {
			courierGet.Id = id
			courierGet.Product.Id = idProduct
			courierGet.Deposit.Id = idDeposit
			courierGet.Client.Id = idClient
			courierArray = append(courierArray, courierGet)
		}
	}

	return courierArray, total, nil
}
