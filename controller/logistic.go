package controller

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/TCC-PucMinas/micro-logistics/communicate"
	model "github.com/TCC-PucMinas/micro-logistics/models"
	"github.com/TCC-PucMinas/micro-logistics/service"
)

type LogisticServer struct {
}

func (s *LogisticServer) CalculateLogistic(ctx context.Context, request *communicate.CalulateRequest) (*communicate.CalculateResponse, error) {

	res := &communicate.CalculateResponse{}

	googleServie := service.Calculate{}

	carry := model.Carrying{}

	if err := carry.GetById(request.IdCarring); err != nil {
		return res, errors.New("Id invalid!")
	}

	client := model.Client{}

	if err := client.GetById(request.IdClient); err != nil {
		return res, errors.New("Id Client invalid!")
	}

	destination := model.Destination{}

	if err := destination.DestinationGetByClientId(request.IdClient); err != nil {
		log.Println("err", err)
		return res, errors.New("Id client invalid!")
	}

	origin := service.LatAndLng{
		Lat: carry.Lat,
		Lng: carry.Lng,
	}

	destiny := service.LatAndLng{
		Lat: destination.Lat,
		Lng: destination.Lng,
	}

	googleServie.Origin = origin
	googleServie.Destiny = destiny

	if err := googleServie.CalculateRoute(); err != nil {
		return res, errors.New("Error in calculate route google maps!")
	}

	res.Origin = &communicate.LatAndLong{
		Lat: googleServie.Origin.Lat,
		Lng: googleServie.Origin.Lng,
	}
	res.Destiny = &communicate.LatAndLong{
		Lat: googleServie.Destiny.Lat,
		Lng: googleServie.Destiny.Lng,
	}

	res.Duration = int64(googleServie.Duration)

	h, _ := time.ParseDuration(googleServie.Duration.String())
	res.Duration = int64(h.Minutes())
	res.HumanReadable = googleServie.HumanReadable
	res.Meters = int64(googleServie.Meters)

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

func (s *LogisticServer) ValidateClientById(ctx context.Context, request *communicate.ValidateClientRequest) (*communicate.ValidateClientResponse, error) {

	res := &communicate.ValidateClientResponse{}

	client := model.Client{}

	if err := client.GetById(request.IdClient); err != nil {
		return res, errors.New("Client Id invalid!")
	}

	res.Valid = true

	return res, nil
}
