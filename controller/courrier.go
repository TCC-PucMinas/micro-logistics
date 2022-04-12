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
		deposit := model.Deposit{}
		_ = deposit.GetById(c.Deposit.Id)
		client, _ := service.ListOneClientById(c.Client.Id)
		product, _ := service.ListOneProductById(c.Product.Id)
		courier.Id = c.Id
		courier.Product = &communicate.ProductLg{Id: product.Product.Id, Name: product.Product.Name}
		courier.Deposit = &communicate.DepositLg{Id: deposit.Id, Name: deposit.Name}
		courier.Client = &communicate.ClientLg{Id: client.Client.Id, Name: client.Client.Name}
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

	deposit := model.Deposit{}
	_ = deposit.GetById(courier.Deposit.Id)
	client, _ := service.ListOneClientById(courier.Client.Id)

	product, _ := service.ListOneProductById(courier.Product.Id)

	courierGet := &communicate.Courier{}
	courierGet.Id = courier.Id
	courierGet.Product = &communicate.ProductLg{Id: product.Product.Id, Name: product.Product.Name}
	courierGet.Client = &communicate.ClientLg{Id: client.Client.Id, Name: client.Client.Name}
	courierGet.Deposit = &communicate.DepositLg{Id: deposit.Id, Name: deposit.Name}

	res.Courier = courierGet

	return res, nil
}

func (s *CourierServer) CreateCourier(ctx context.Context, request *communicate.CreateCourierRequest) (*communicate.CreateCourierResponse, error) {
	res := &communicate.CreateCourierResponse{}
	courier := model.Courier{}

	if err := courier.GetOneByIdDriverAndIdProduct(request.IdProduct); err != nil || courier.Id != 0 {
		return res, errors.New("Not duplicated courier")
	}

	courier.Product = model.Product{Id: request.IdProduct}
	courier.Client = model.Client{Id: request.IdClient}
	courier.Deposit = model.Deposit{Id: request.IdDeposit}

	_, err := courier.Create()

	if err != nil {
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
		Id:        request.Id,
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

	if err := courier.GetOneByIdDriverAndIdProduct(request.IdProduct); err != nil || courier.Id != 0 {
		return res, nil
	}

	res.Valid = true

	return res, nil
}
