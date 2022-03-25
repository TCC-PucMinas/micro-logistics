package controller

import (
	"context"
	"errors"
	"log"
	"micro-logistic/communicate"
	model "micro-logistic/models"
)

type CourierServer struct{}

func (s *CourierServer) CourierListAll(ctx context.Context, request *communicate.CourierListAllRequest) (*communicate.CourierListAllResponse, error) {
	res := &communicate.CourierListAllResponse{}

	var courier model.Courier

	arrayCourier, total, err := courier.GetCourierPaginate(request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.DataCourier{}
	for _, c := range arrayCourier {
		courier := &communicate.Courier{}
		courier.Id = c.Id
		courier.IdDriver = c.Driver.Id
		courier.IdProduct = c.Product.Id
		data.Courier = append(data.Courier, courier)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}

func (s *CourierServer) ListOneCourierById(ctx context.Context, request *communicate.ListOneCourierByIdRequest) (*communicate.ListOneCourierByIdResponse, error) {
	res := &communicate.ListOneCourierByIdResponse{}

	var courier model.Courier

	if err := courier.GetById(request.Id); err != nil || courier.Id == 0 {
		log.Println(err)
		return res, errors.New("Courier not found!")
	}

	courierGet := &communicate.Courier{}
	courierGet.Id = courier.Id
	courierGet.IdDriver = courier.Driver.Id
	courierGet.IdProduct = courier.Product.Id

	res.Courier = courierGet

	return res, nil
}

func (s *CourierServer) CreateCourier(ctx context.Context, request *communicate.CreateCourierRequest) (*communicate.CreateCourierResponse, error) {
	res := &communicate.CreateCourierResponse{}

	courier := model.Courier{
		Driver:  model.Driver{Id: request.IdDriver},
		Product: model.Product{Id: request.IdProduct},
	}

	if err := courier.Create(); err != nil {
		return res, errors.New("Error creating courier!")
	}

	res.Created = true

	return res, nil
}

func (s *CourierServer) UpdateCourierById(ctx context.Context, request *communicate.UpdateCourierByIdRequest) (*communicate.UpdateCourierByIdResponse, error) {
	res := &communicate.UpdateCourierByIdResponse{
		Updated: false,
	}

	courier := model.Courier{}

	if err := courier.GetById(request.Id); err != nil || courier.Id == 0 {
		return res, errors.New("Courier not found!")
	}

	courier = model.Courier{
		Id:      request.Id,
		Driver:  model.Driver{Id: request.IdDriver},
		Product: model.Product{Id: request.IdProduct},
	}

	if err := courier.UpdateById(); err != nil {
		return res, errors.New("Erro updating courier!")
	}

	res.Updated = true

	return res, nil
}

func (s *CourierServer) DeleteCourierById(ctx context.Context, request *communicate.DeleteCourierByIdRequest) (*communicate.DeleteCourierByIdResponse, error) {
	res := &communicate.DeleteCourierByIdResponse{}

	courier := model.Courier{}

	if err := courier.GetById(request.Id); err != nil || courier.Id == 0 {
		return res, errors.New("Courier not found!")
	}

	if err := courier.DeleteById(); err != nil {
		return res, errors.New("Erro deleting courier!")
	}

	res.Deleted = true

	return res, nil
}
