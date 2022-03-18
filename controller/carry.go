package controller

import (
	"context"
	"errors"
	"micro-logistic/communicate"
	model "micro-logistic/models"
	"micro-logistic/service"
)

type CarryServer struct{}

func (s *CarryServer) CarryListAll(ctx context.Context, request *communicate.CarryListAllRequest) (*communicate.CarryListAllResponse, error) {
	res := &communicate.CarryListAllResponse{}

	var carry model.Carrying

	carrings, total, err := carry.GetCarryByNamePaginate(request.Name, request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.DataCarry{}
	for _, c := range carrings {
		carry := &communicate.Carry{}
		carry.Id = c.Id
		carry.Name = c.Name
		carry.City = c.City
		carry.Country = c.Country
		carry.State = c.State
		carry.Street = c.Street
		carry.District = c.District
		carry.ZipCode = c.ZipCode
		carry.Number = c.Number
		carry.Lat = c.Lat
		carry.Lng = c.Lng
		data.Carry = append(data.Carry, carry)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}

func (s *CarryServer) ListOneCarryById(ctx context.Context, request *communicate.ListOneCarryByIdRequest) (*communicate.ListOneCarryByIdResponse, error) {
	res := &communicate.ListOneCarryByIdResponse{}

	var carry model.Carrying

	if err := carry.GetById(request.Id); err != nil || carry.Id == 0 {
		return res, errors.New("Carrying Id invalid!")
	}

	carryResponse := &communicate.Carry{}
	carryResponse.Id = carry.Id
	carryResponse.Name = carry.Name
	carryResponse.City = carry.City
	carryResponse.Country = carry.Country
	carryResponse.State = carry.State
	carryResponse.Street = carry.Street
	carryResponse.District = carry.District
	carryResponse.ZipCode = carry.ZipCode
	carryResponse.Number = carry.Number
	carryResponse.Lat = carry.Lat
	carryResponse.Lng = carry.Lng

	res.Carry = carryResponse
	return res, nil
}

func (s *CarryServer) CreateCarry(ctx context.Context, request *communicate.CreateCarryRequest) (*communicate.CreateCarryResponse, error) {
	res := &communicate.CreateCarryResponse{}

	carry := model.Carrying{
		Name:     request.Name,
		City:     request.City,
		Country:  request.Country,
		State:    request.State,
		Street:   request.Street,
		ZipCode:  request.ZipCode,
		District: request.District,
		Number:   request.Number,
	}

	if err := carry.GetByName(request.Name); err != nil || carry.Id != 0 {
		return res, errors.New("carry not duplicated!")
	}

	lat, lng, err := service.GetLocationCarrying(carry)

	if err != nil {
		return res, err
	}

	carry.Lat = lat
	carry.Lng = lng

	if err := carry.CreateCarry(); err != nil {
		return res, errors.New("Error carrying client!")
	}

	res.Created = true

	return res, nil
}

func (s *CarryServer) UpdateCarryById(ctx context.Context, request *communicate.UpdateCarryByIdRequest) (*communicate.UpdateCarryByIdResponse, error) {
	res := &communicate.UpdateCarryByIdResponse{}

	carry := model.Carrying{
		Id:       request.Id,
		Name:     request.Name,
		City:     request.City,
		Country:  request.Country,
		State:    request.State,
		Street:   request.Street,
		ZipCode:  request.ZipCode,
		District: request.District,
		Number:   request.Number,
	}

	lat, lng, err := service.GetLocationCarrying(carry)

	if err != nil {
		return res, err
	}

	carry.Lat = lat
	carry.Lng = lng

	if err := carry.GetById(request.Id); err != nil || carry.Id == 0 {
		return res, errors.New("Carry not found!")
	}

	if err := carry.UpdateCarryById(); err != nil {
		return res, errors.New("Erro updating carry!")
	}

	res.Updated = true

	return res, nil
}

func (s *CarryServer) DeleteCarryById(ctx context.Context, request *communicate.DeleteCarryByIdRequest) (*communicate.DeleteCarryByIdResponse, error) {
	res := &communicate.DeleteCarryByIdResponse{}

	carry := model.Carrying{}

	if err := carry.GetById(request.Id); err != nil || carry.Id == 0 {
		return res, errors.New("Carrying not found!")
	}

	if err := carry.DeleteById(); err != nil {
		return res, errors.New("Erro deleting carrying!")
	}

	res.Deleted = true

	return res, nil
}

func (s *CarryServer) ValidateCarryExist(ctx context.Context, request *communicate.ValidateCarryCreateRequest) (*communicate.ValidateCarryCreateResponse, error) {
	res := &communicate.ValidateCarryCreateResponse{}

	carry := model.Carrying{}

	if err := carry.GetByName(request.Name); err != nil || carry.Id != 0 {
		return res, errors.New("Client duplicated!")
	}

	res.Valid = true

	return res, nil
}

func (s *CarryServer) ValidateCarryById(ctx context.Context, request *communicate.ValidateCarryByIdRequest) (*communicate.ValidateCarryByIdResponse, error) {

	res := &communicate.ValidateCarryByIdResponse{}

	carry := model.Carrying{}

	if err := carry.GetById(request.IdCarry); err != nil {
		return res, errors.New("Carrying Id invalid!")
	}

	res.Valid = true

	return res, nil
}
