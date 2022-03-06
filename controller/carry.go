package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
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
		District: request.District,
		Number:   request.Number,
	}

	if err := carry.GetByName(request.Name); err != nil || carry.Id != 0 {
		return res, errors.New("carry not duplicated!")
	}

	latAndLng := service.LatAndLng{}
	address := fmt.Sprintf("%v, %v, %v, %v, %v, %v", carry.Street, carry.Number, carry.District, carry.City, carry.State, carry.Country)
	if err := latAndLng.GetLatAndLngByAddress(address); err != nil {
		return res, err
	}

	carry.Lat = latAndLng.Lat
	carry.Lng = latAndLng.Lng

	if err := carry.CreateCarry(); err != nil {
		log.Println("err", err)
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
		District: request.District,
		Number:   request.Number,
	}

	latAndLng := service.LatAndLng{}
	address := fmt.Sprintf("%v, %v, %v, %v, %v, %v", carry.Street, carry.Number, carry.District, carry.City, carry.State, carry.Country)
	if err := latAndLng.GetLatAndLngByAddress(address); err != nil {
		return res, err
	}

	carry.Lat = latAndLng.Lat
	carry.Lng = latAndLng.Lng

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

	client := model.Carrying{}

	if err := client.GetById(request.IdClient); err != nil {
		return res, errors.New("Carrying Id invalid!")
	}

	res.Valid = true

	return res, nil
}
