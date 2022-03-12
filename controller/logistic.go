package controller

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"micro-logistic/communicate"
	model "micro-logistic/models"
)

type LogisticServer struct {
}

func (s *LogisticServer) CalculateLogistic(ctx context.Context, request *communicate.CalulateRequest) (*communicate.CalculateResponse, error) {

	res := &communicate.CalculateResponse{}

	conn, err := grpc.Dial(":6000", grpc.WithInsecure())

	if err != nil {
		return res, err
	}

	defer conn.Close()
	carry := model.Carrying{}

	if err := carry.GetById(request.IdCarring); err != nil {
		return res, errors.New("Id invalid!")
	}

	serviceClient := communicate.NewClientCommunicateClient(conn)

	validateClientById := &communicate.ValidateClientByIdRequest{
		IdClient: request.IdClient,
	}

	if _, err := serviceClient.ValidateClientById(ctx, validateClientById); err != nil {
		return res, errors.New("Id Client invalid!")
	}

	serviceDestination := communicate.NewDestinationCommunicateClient(conn)

	requestDestinationById := &communicate.ListOneDestinationByIdRequest{
		Id: request.IdDestination,
	}

	destination, err := serviceDestination.ListOneDestinationById(ctx, requestDestinationById)

	if err != nil {
		return res, err
	}

	origin := &communicate.LatAndLng{
		Lat: carry.Lat,
		Lng: carry.Lng,
	}

	destiny := &communicate.LatAndLng{
		Lat: destination.Destination.Lat,
		Lng: destination.Destination.Lng,
	}

	requestCommunicateClientGelocation := &communicate.DirectionLocationRequest{
		Origin:  origin,
		Destiny: destiny,
	}

	connGeolocation, err := grpc.Dial(":7000", grpc.WithInsecure())

	if err != nil {
		return res, err
	}

	defer connGeolocation.Close()

	serviceGeolocation := communicate.NewGelocationCommunicateClient(connGeolocation)

	responseRequest, err := serviceGeolocation.DirectionLocation(ctx, requestCommunicateClientGelocation)

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

	res.HumanReadable = responseRequest.HumanReadable
	res.Meters = responseRequest.Meters

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
