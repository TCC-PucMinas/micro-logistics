package controller

import (
	"context"
	"errors"
	"micro-logistic/communicate"
	model "micro-logistic/models"
	"micro-logistic/service"
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
		courier.IdDeposit = c.Deposit.Id
		courier.IdClient = c.Client.Id
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
		return res, errors.New("Courier not found!")
	}

	courierGet := &communicate.Courier{}
	courierGet.Id = courier.Id
	courierGet.IdDriver = courier.Driver.Id
	courierGet.IdProduct = courier.Product.Id
	courierGet.IdClient = courier.Client.Id
	courierGet.IdDeposit = courier.Deposit.Id

	res.Courier = courierGet

	return res, nil
}

func (s *CourierServer) CreateCourier(ctx context.Context, request *communicate.CreateCourierRequest) (*communicate.CreateCourierResponse, error) {
	res := &communicate.CreateCourierResponse{}
	courier := model.Courier{}

	if err := courier.GetOneByIdDriverAndIdProduct(request.IdDriver, request.IdProduct); err != nil || courier.Id != 0 {
		return res, errors.New("Not duplicated courier")
	}

	courier.Driver = model.Driver{Id: request.IdDriver}
	courier.Product = model.Product{Id: request.IdProduct}
	courier.Client = model.Client{Id: request.IdClient}
	courier.Deposit = model.Deposit{Id: request.IdDeposit}

	id, err := courier.Create()

	if err != nil {
		return res, errors.New("Error creating courier!")
	}

	destiny, err := service.ListOneDestinationByIdProduct(request.IdProduct)

	if err != nil {
		return res, err
	}

	if err := courier.Deposit.GetById(request.IdDeposit); err != nil {
		return res, err
	}

	courierRoute := model.CourierRoute{
		Courier:   model.Courier{Id: id},
		Order:     0,
		LatInit:   model.LatAndLng{Lat: courier.Deposit.Lat, Lng: courier.Deposit.Lng},
		LatFinish: model.LatAndLng{Lat: destiny.Destination.Lat, Lng: destiny.Destination.Lng},
	}

	if err := courierRoute.CreateCourierRoute(); err != nil {
		return res, err
	}

	route := service.NewRoutes(&courierRoute)

	go route.TracingRoutes()

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
		Id:        request.Id,
		Driver:    model.Driver{Id: request.IdDriver},
		Product:   model.Product{Id: request.IdProduct},
		Client:    model.Client{Id: request.IdClient},
		Deposit:   model.Deposit{Id: request.IdDeposit},
		Doc:       model.Documentation{Type: request.Doc.Type, Value: request.Doc.Value},
		Delivered: request.Delivered,
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

func (s *CourierServer) ValidateCourier(ctx context.Context, request *communicate.CourierValidateRequest) (*communicate.CourierValidateResponse, error) {
	res := &communicate.CourierValidateResponse{
		Valid: false,
	}

	courier := model.Courier{}

	if err := courier.GetOneByIdDriverAndIdProduct(request.IdDriver, request.IdProduct); err != nil || courier.Id != 0 {
		return res, nil
	}

	res.Valid = true

	return res, nil
}
