package controller

import (
	"context"
	"errors"
	"micro-logistic/communicate"
	model "micro-logistic/models"
)

type DriverServer struct{}

func (d *DriverServer) ValidateDriverExistByNameAndIdCarry(ctx context.Context, request *communicate.ValidateDriverExistByNameAndIdCarryRequest) (*communicate.ValidateDriverExistByNameAndIdCarryResponse, error) {
	res := &communicate.ValidateDriverExistByNameAndIdCarryResponse{Valid: false}

	driver := model.Driver{}

	if err := driver.GetByNameAndIdCarry(request.Name, request.IdCarry); err != nil || driver.Id > 0 {
		return res, nil
	}

	res.Valid = true

	return res, nil
}

func (d *DriverServer) DriverListAll(ctx context.Context, request *communicate.DriverListAllRequest) (*communicate.DriverListAllResponse, error) {
	res := &communicate.DriverListAllResponse{}

	var driver model.Driver

	drivers, total, err := driver.GetDriverPaginateByNameAndIdCarry(request.Name, request.IdCarry, request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.DataDriver{}
	for _, c := range drivers {
		driver := &communicate.Driver{}
		driver.Id = c.Id
		driver.Name = c.Name
		driver.Image = c.Image
		driver.IdCarry = c.Carrying.Id
		driver.IdTruck = c.Truck.Id

		data.Driver = append(data.Driver, driver)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}

func (d *DriverServer) ListOneDriverById(ctx context.Context, request *communicate.ListOneDriverByIdRequest) (*communicate.ListOneDriverByIdResponse, error) {
	res := &communicate.ListOneDriverByIdResponse{}

	var driver model.Driver

	if err := driver.GetById(request.Id); err != nil || driver.Id == 0 {
		return res, errors.New("Driver Id invalid!")
	}

	driverResponse := &communicate.Driver{}
	driverResponse.Id = driver.Id
	driverResponse.Name = driver.Name
	driverResponse.Image = driver.Image
	driverResponse.IdCarry = driver.Carrying.Id
	driverResponse.IdTruck = driver.Truck.Id

	res.Driver = driverResponse
	return res, nil
}

func (d *DriverServer) CreateDriver(ctx context.Context, request *communicate.CreateDriverRequest) (*communicate.CreateDriverResponse, error) {
	res := &communicate.CreateDriverResponse{}

	driver := model.Driver{
		Name:     request.Name,
		Image:    request.Image,
		Carrying: model.Carrying{Id: request.IdCarry},
		Truck:    model.Truck{Id: request.IdTruck},
	}

	if err := driver.CreateDriver(); err != nil {
		return res, errors.New("Error creating driver!")
	}

	res.Created = true

	return res, nil
}

func (d *DriverServer) UpdateDriverById(ctx context.Context, request *communicate.UpdateDriverByIdRequest) (*communicate.UpdateDriverByIdResponse, error) {
	res := &communicate.UpdateDriverByIdResponse{}

	driver := model.Driver{}

	if err := driver.GetById(request.Id); err != nil || driver.Id == 0 {
		return res, errors.New("Truck not found!")
	}

	driver = model.Driver{
		Id:       request.Id,
		Name:     request.Name,
		Image:    request.Image,
		Carrying: model.Carrying{Id: request.IdCarry},
		Truck:    model.Truck{Id: request.IdTruck},
	}

	if err := driver.UpdateDriver(); err != nil {
		return res, errors.New("Erro updating driver!")
	}

	res.Updated = true

	return res, nil
}

func (d *DriverServer) DeleteDriverById(ctx context.Context, request *communicate.DeleteDriverByIdRequest) (*communicate.DeleteDriverByIdResponse, error) {
	res := &communicate.DeleteDriverByIdResponse{}

	driver := model.Driver{}

	if err := driver.GetById(request.Id); err != nil || driver.Id == 0 {
		return res, errors.New("Driver not found!")
	}

	if err := driver.DeleteById(); err != nil {
		return res, errors.New("Erro deleting driver!")
	}

	res.Deleted = true

	return res, nil
}
