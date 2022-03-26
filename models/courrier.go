package model

import (
	"fmt"
	"micro-logistic/db"
	"micro-logistic/helpers"
)

type Courier struct {
	Id        int64   `json:"id"`
	Driver    Driver  `json:"driver"`
	Deposit   Deposit `json:"deposit"`
	Client    Client  `json:"client"`
	Product   Product `json:"product"`
	Delivered bool    `json:"delivered"`
}

func (c *Courier) GetOneByIdDriverAndIdProduct(idDriver, idProduct int64) error {

	sql := db.ConnectDatabase()

	query := `select id, id_driver, id_product from couriers where id_driver = ? and id_product = ? limit 1;`

	requestConfig, err := sql.Query(query, idDriver, idProduct)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var idDriver, idProduct, id int64
		_ = requestConfig.Scan(&id, &idDriver, &idProduct)
		if id != 0 {
			c.Id = id
			c.Driver.Id = idDriver
			c.Product.Id = idProduct
		}
	}

	return nil
}

func (c *Courier) GetById(id int64) error {

	sql := db.ConnectDatabase()

	query := `select id, id_driver, id_deposit, id_client, id_product from couriers where id = ? limit 1;`

	requestConfig, err := sql.Query(query, id)

	if err != nil {
		return err
	}

	for requestConfig.Next() {
		var idDriver, idProduct, idClient, idDeposit, id int64
		_ = requestConfig.Scan(&id, &idDriver, &idDeposit, &idClient, &idProduct)
		if id != 0 {
			c.Id = id
			c.Driver.Id = idDriver
			c.Client.Id = idClient
			c.Deposit.Id = idDeposit
			c.Product.Id = idProduct
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

	query := "update couriers set id_driver = ?, id_deposit = ?, id_client = ?, id_product = ? where id = ?"

	destinationUpdate, err := sql.Prepare(query)

	if err != nil {
		return err
	}

	_, e := destinationUpdate.Exec(c.Driver.Id, c.Deposit.Id, c.Client.Id, c.Product.Id, c.Id)

	if e != nil {
		return e
	}

	return nil
}

func (c *Courier) Create() (int64, error) {
	sql := db.ConnectDatabase()

	query := "insert into couriers (id, id_driver, id_deposit, id_client, id_product) values (?, ?, ?, ?, ?)"

	createDestination, err := sql.Prepare(query)

	if err != nil {
		return 0, err
	}

	result, e := createDestination.Exec(c.Id, c.Driver.Id, c.Deposit.Id, c.Client.Id, c.Product.Id)

	if e != nil {
		return 0, e
	}

	return result.LastInsertId()
}

func (c *Courier) GetCouriers() ([]Courier, error) {
	var courierArray []Courier

	sql := db.ConnectDatabase()
	query := "select id, id_driver, id_deposit, id_client, id_product from couriers"

	requestConfig, err := sql.Query(query)

	if err != nil {
		return courierArray, err
	}

	for requestConfig.Next() {
		courierGet := Courier{}
		var idDriver, idProduct, idClient, idDeposit, id int64
		_ = requestConfig.Scan(&id, &idDriver, &idDeposit, &idClient, &idProduct)
		if id != 0 {
			courierGet.Id = id
			courierGet.Driver.Id = idDriver
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

	query := fmt.Sprintf("select id, id_driver, id_deposit, id_client, id_product, %v from couriers LIMIT ? OFFSET ?;", paginate.Query)

	requestConfig, err := sql.Query(query, paginate.Limit, paginate.Page)

	if err != nil {
		return courierArray, total, err
	}

	for requestConfig.Next() {
		courierGet := Courier{}
		var idDriver, idProduct, idClient, idDeposit, id int64
		_ = requestConfig.Scan(&id, &idDriver, &idDeposit, &idClient, &idProduct, &total)
		if id != 0 {
			courierGet.Id = id
			courierGet.Driver.Id = idDriver
			courierGet.Product.Id = idProduct
			courierGet.Deposit.Id = idDeposit
			courierGet.Client.Id = idClient
			courierArray = append(courierArray, courierGet)
		}
	}

	return courierArray, total, nil
}
