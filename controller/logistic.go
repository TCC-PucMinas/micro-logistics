package controller

import (
	"context"
	"errors"
	"log"
	"micro-logistic/communicate"
	model "micro-logistic/models"
	"micro-logistic/service"
)

type LogisticServer struct{}

func (s *LogisticServer) CalculateLogistic(ctx context.Context, request *communicate.CalulateRequest) (*communicate.CalculateResponse, error) {

	res := &communicate.CalculateResponse{}

	carry := model.Carrying{}

	if err := carry.GetById(request.IdCarring); err != nil {
		return res, errors.New("Id invalid!")
	}

	if err := service.ValidateClientById(request.IdClient); err != nil {
		log.Println("err", err)
		return res, errors.New("Id Client invalid!")
	}

	destination, err := service.ListOneDestinationById(request.IdDestination)

	if err != nil {
		return res, err
	}

	origin := service.Origin{
		Lat: carry.Lat,
		Lng: carry.Lng,
	}

	destiny := service.Destiny{
		Lat: destination.Destination.Lat,
		Lng: destination.Destination.Lng,
	}

	humanReadable, meters, err := service.CalculateRouter(origin, destiny)

	if err != nil {
		return res, err
	}

	res.Origin = &communicate.LatAndLong{
		Lat: carry.Lat,
		Lng: carry.Lng,
	}
	res.Destiny = &communicate.LatAndLong{
		Lat: destination.Destination.Lat,
		Lng: destination.Destination.Lng,
	}

	res.HumanReadable = humanReadable
	res.Meters = meters

	return res, nil

}

func (s *LogisticServer) ValidateCarringById(ctx context.Context, request *communicate.ValidateCarryingRequest) (*communicate.ValidateCarryingResponse, error) {

	res := &communicate.ValidateCarryingResponse{}

	carry := model.Carrying{}

	if err := carry.GetById(request.IdCarring); err != nil {
		return res, errors.New("Carrying Id invalid!")
	}

	res.Valid = true

	return res, nil
}
