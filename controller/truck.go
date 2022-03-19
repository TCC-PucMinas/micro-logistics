package controller

import (
	"context"
	"errors"
	"log"
	"micro-logistic/communicate"
	model "micro-logistic/models"
)

type TruckServer struct{}

func (s *TruckServer) TruckListAll(ctx context.Context, request *communicate.TruckListAllRequest) (*communicate.TruckListAllResponse, error) {
	res := &communicate.TruckListAllResponse{}

	var truck model.Truck

	trucks, total, err := truck.GetTruckByIdCarryPaginate(request.IdCarry, request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.DataTruck{}
	for _, c := range trucks {
		truck := &communicate.Truck{}
		truck.Id = c.Id
		truck.Brand = c.Brand
		truck.Model = c.Model
		truck.Plate = c.Plate
		truck.Year = c.Year
		truck.IdCarry = c.Carrying.Id

		data.Truck = append(data.Truck, truck)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}

func (s *TruckServer) ListOneTruckById(ctx context.Context, request *communicate.ListOneTruckByIdRequest) (*communicate.ListOneTruckByIdResponse, error) {
	res := &communicate.ListOneTruckByIdResponse{}

	var truck model.Truck

	if err := truck.GetById(request.Id); err != nil || truck.Id == 0 {
		return res, errors.New("Truck Id invalid!")
	}

	truckResponse := &communicate.Truck{}
	truckResponse.Id = truck.Id
	truckResponse.Brand = truck.Brand
	truckResponse.Model = truck.Model
	truckResponse.Plate = truck.Plate
	truckResponse.Year = truck.Year
	truckResponse.IdCarry = truck.Carrying.Id

	res.Truck = truckResponse
	return res, nil
}

func (s *TruckServer) CreateTruck(ctx context.Context, request *communicate.CreateTruckRequest) (*communicate.CreateTruckResponse, error) {
	res := &communicate.CreateTruckResponse{}

	truck := model.Truck{
		Brand:    request.Brand,
		Model:    request.Model,
		Plate:    request.Plate,
		Year:     request.Year,
		Carrying: model.Carrying{Id: request.IdCarry},
	}

	if err := truck.CreateTruck(); err != nil {
		return res, errors.New("Error creating truck!")
	}

	res.Created = true

	return res, nil
}

func (s *TruckServer) UpdateTruckById(ctx context.Context, request *communicate.UpdateTruckByIdRequest) (*communicate.UpdateTruckByIdResponse, error) {
	res := &communicate.UpdateTruckByIdResponse{}

	truck := model.Truck{}

	if err := truck.GetById(request.Id); err != nil || truck.Id == 0 {
		return res, errors.New("Truck not found!")
	}

	truck = model.Truck{
		Id:       request.Id,
		Brand:    request.Brand,
		Model:    request.Model,
		Plate:    request.Plate,
		Year:     request.Year,
		Carrying: model.Carrying{Id: request.IdCarry},
	}

	if err := truck.UpdateTruck(); err != nil {
		log.Println("update", err)
		return res, errors.New("Erro updating truck!")
	}

	res.Updated = true

	return res, nil
}

func (s *TruckServer) DeleteTruckById(ctx context.Context, request *communicate.DeleteTruckByIdRequest) (*communicate.DeleteTruckByIdResponse, error) {
	res := &communicate.DeleteTruckByIdResponse{}

	truck := model.Truck{}

	if err := truck.GetById(request.Id); err != nil || truck.Id == 0 {
		return res, errors.New("Truck not found!")
	}

	if err := truck.DeleteById(); err != nil {
		return res, errors.New("Erro deleting truck!")
	}

	res.Deleted = true

	return res, nil
}
