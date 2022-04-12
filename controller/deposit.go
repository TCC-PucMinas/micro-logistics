package controller

import (
	"context"
	"errors"
	"micro-logistic/communicate"
	model "micro-logistic/models"
	"micro-logistic/service"
)

type DepositServer struct{}

func (s *DepositServer) DepositListAll(ctx context.Context, request *communicate.DepositListAllRequest) (*communicate.DepositListAllResponse, error) {
	res := &communicate.DepositListAllResponse{}

	var deposit model.Deposit

	deposits, total, err := deposit.GetDepositByNamePaginate(request.Name, request.IdCarry, request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.DataDeposit{}
	for _, c := range deposits {
		deposit := &communicate.Deposit{}
		deposit.Id = c.Id
		deposit.Name = c.Name
		deposit.City = c.City
		deposit.Country = c.Country
		deposit.State = c.State
		deposit.Street = c.Street
		deposit.District = c.District
		deposit.Number = c.Number
		deposit.Lat = c.Lat
		deposit.Lng = c.Lng
		deposit.IdCarry = c.Carrying.Id
		data.Deposit = append(data.Deposit, deposit)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}

func (s *DepositServer) ListOneDepositById(ctx context.Context, request *communicate.ListOneDepositByIdRequest) (*communicate.ListOneDepositByIdResponse, error) {
	res := &communicate.ListOneDepositByIdResponse{}

	var deposit model.Deposit

	if err := deposit.GetById(request.Id); err != nil || deposit.Id == 0 {
		return res, err
	}

	depositGet := &communicate.Deposit{}
	depositGet.Id = deposit.Id
	depositGet.Name = deposit.Name
	depositGet.City = deposit.City
	depositGet.Country = deposit.Country
	depositGet.State = deposit.State
	depositGet.Street = deposit.Street
	depositGet.District = deposit.District
	depositGet.Number = deposit.Number
	depositGet.Lat = deposit.Lat
	depositGet.Lng = deposit.Lng
	depositGet.ZipCode = deposit.ZipCode
	depositGet.IdCarry = deposit.Carrying.Id

	res.Deposit = depositGet

	return res, nil
}

func (s *DepositServer) CreateDeposit(ctx context.Context, request *communicate.CreateDepositRequest) (*communicate.CreateDepositResponse, error) {
	res := &communicate.CreateDepositResponse{}

	deposit := model.Deposit{
		Name:     request.Name,
		City:     request.City,
		Country:  request.Country,
		State:    request.State,
		Street:   request.Street,
		ZipCode:  request.ZipCode,
		District: request.District,
		Number:   request.Number,
		Carrying: model.Carrying{Id: request.IdCarry},
	}

	if err := deposit.GetByNameAndIdCarry(request.Name, request.IdCarry); err != nil || deposit.Id != 0 {
		return res, errors.New("deposit not duplicated!")
	}

	lat, lng, err := service.GetLocationDeposit(deposit)

	if err != nil {
		return res, err
	}

	deposit.Lat = lat
	deposit.Lng = lng

	if err := deposit.CreateDeposit(); err != nil {
		return res, errors.New("Error creating deposit!")
	}

	res.Created = true

	return res, nil
}

func (s *DepositServer) UpdateDepositById(ctx context.Context, request *communicate.UpdateDepositByIdRequest) (*communicate.UpdateDepositByIdResponse, error) {
	res := &communicate.UpdateDepositByIdResponse{}

	deposit := model.Deposit{}

	if err := deposit.GetById(request.Id); err != nil || deposit.Id == 0 {
		return res, errors.New("Deposit not found!")
	}

	deposit = model.Deposit{
		Id:       request.Id,
		Name:     request.Name,
		City:     request.City,
		Country:  request.Country,
		State:    request.State,
		ZipCode:  request.ZipCode,
		Street:   request.Street,
		District: request.District,
		Number:   request.Number,
		Carrying: model.Carrying{Id: request.IdCarry},
	}

	lat, lng, err := service.GetLocationDeposit(deposit)

	if err != nil {
		return res, err
	}

	deposit.Lat = lat
	deposit.Lng = lng

	if err := deposit.UpdateDepositById(); err != nil {
		return res, errors.New("Erro updating deposit!")
	}

	res.Updated = true

	return res, nil
}

func (s *DepositServer) DeleteDepositById(ctx context.Context, request *communicate.DeleteDepositByIdRequest) (*communicate.DeleteDepositByIdResponse, error) {
	res := &communicate.DeleteDepositByIdResponse{}

	deposit := model.Deposit{}

	if err := deposit.GetById(request.Id); err != nil || deposit.Id == 0 {
		return res, errors.New("Deposit not found!")
	}

	if err := deposit.DeleteById(); err != nil {
		return res, errors.New("Erro deleting deposit!")
	}

	res.Deleted = true

	return res, nil
}

func (s *DepositServer) ValidateDepositExist(ctx context.Context, request *communicate.ValidateDepositCreateRequest) (*communicate.ValidateDepositCreateResponse, error) {
	res := &communicate.ValidateDepositCreateResponse{}

	deposit := model.Deposit{}

	if err := deposit.GetByNameAndIdCarry(request.Name, request.IdCarry); err != nil || deposit.Id != 0 {
		return res, errors.New("Deposit duplicated!")
	}

	res.Valid = true

	return res, nil
}

func (s *DepositServer) ValidateDepositById(ctx context.Context, request *communicate.ValidateDepositByIdRequest) (*communicate.ValidateDepositByIdResponse, error) {

	res := &communicate.ValidateDepositByIdResponse{}

	deposit := model.Deposit{}

	if err := deposit.GetById(request.IdClient); err != nil {
		return res, errors.New("Deposit Id invalid!")
	}

	res.Valid = true

	return res, nil
}
