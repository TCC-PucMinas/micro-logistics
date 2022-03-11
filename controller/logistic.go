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

	var conn *grpc.ClientConn

	defer conn.Close()

	res := &communicate.CalculateResponse{}
	//googleServie := service.Calculate{}
	carry := model.Carrying{}

	if err := carry.GetById(request.IdCarring); err != nil {
		return res, errors.New("Id invalid!")
	}

	serviceClient := communicate.NewClientCommunicateClient(conn)

	validateClientById := &communicate.ValidateClientByIdRequest{
		Id: request.IdClient,
	}

	if _, err := serviceClient.ValidateClientById(ctx, validateClientById); err != nil {
		return res, errors.New("Id Client invalid!")
	}

	if err != nil {
		return res, err
	}

	// comunicação com o micro serviço do client
	//
	//if err := destination(request.IdClient); err != nil {
	//	return res, errors.New("Id client invalid!")
	//}

	origin := &communicate.LatAndLng{
		Lat: carry.Lat,
		Lng: carry.Lng,
	}

	destiny := &communicate.LatAndLng{
		Lat: destination.Lat,
		Lng: destination.Lng,
	}

	requestCommunicateClientGelocation := &communicate.DirectionLocationRequest{
		Origin:  origin,
		Destiny: destiny,
	}

	serviceGeolocation := communicate.NewGelocationCommunicateClient(conn)

	responseRequest, err := serviceGeolocation.DirectionLocation(ctx, requestCommunicateClientGelocation)

	if err != nil {
		return res, err
	}

	res.Origin = &communicate.LatAndLong{
		Lat: responseRequest.Origin.Lat,
		Lng: responseRequest.Origin.Lng,
	}
	res.Destiny = &communicate.LatAndLong{
		Lat: responseRequest.Destiny.Lat,
		Lng: responseRequest.Destiny.Lng,
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
