package service

import (
	"context"
	"google.golang.org/grpc"
	"micro-logistic/communicate"
	model "micro-logistic/models"
	"time"
)

const attemptRetry = 20

type Origin struct {
	Lat string
	Lng string
}

type Destiny struct {
	Lat string
	Lng string
}

func integrationDirectionLocation(c *communicate.DirectionLocationRequest, retry int) (*communicate.DirectionLocationResponse, error) {
	ctx := context.Background()
	connGeolocation, err := grpc.Dial(":7000", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer connGeolocation.Close()

	serviceLocation := communicate.NewGelocationCommunicateClient(connGeolocation)

	location, err := serviceLocation.DirectionLocation(ctx, c)

	if err != nil {
		return attempRetryLatency(retry, location, err, c)
	}

	return location, nil
}

func integrationGeolocation(c *communicate.GelocationRequest, retry int) (*communicate.GelocationResponse, error) {
	ctx := context.Background()
	connGeolocation, err := grpc.Dial(":7000", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer connGeolocation.Close()

	serviceLocation := communicate.NewGelocationCommunicateClient(connGeolocation)

	location, err := serviceLocation.GetLocation(ctx, c)

	if err != nil {
		return attempRetryLatencyGeolocation(retry, location, err, c)
	}

	return location, nil
}

func attempRetryLatency(retry int, location *communicate.DirectionLocationResponse, err error, c *communicate.DirectionLocationRequest) (*communicate.DirectionLocationResponse, error) {
	retry += 1
	if retry <= attemptRetry {
		time.Sleep(1 * time.Second)
		return integrationDirectionLocation(c, retry)
	}
	return location, err
}

func attempRetryLatencyGeolocation(retry int, location *communicate.GelocationResponse, err error, c *communicate.GelocationRequest) (*communicate.GelocationResponse, error) {
	retry += 1
	if retry <= attemptRetry {
		time.Sleep(1 * time.Second)
		return integrationGeolocation(c, retry)
	}
	return location, err
}

func CalculateRouter(o Origin, d Destiny) (string, int64, error) {
	origin := &communicate.LatAndLng{
		Lat: o.Lat,
		Lng: o.Lng,
	}

	destiny := &communicate.LatAndLng{
		Lat: d.Lat,
		Lng: d.Lng,
	}

	requestLocation := &communicate.DirectionLocationRequest{
		Origin:  origin,
		Destiny: destiny,
	}

	location, err := integrationDirectionLocation(requestLocation, 1)

	if err != nil {
		return "", 0, err
	}

	return location.HumanReadable, location.Meters, nil
}

func GetLocationCarrying(request model.Carrying) (string, string, error) {
	requestLocation := &communicate.GelocationRequest{
		Street:   request.Street,
		District: request.District,
		City:     request.City,
		Country:  request.Country,
		ZipCode:  request.ZipCode,
		State:    request.State,
		Number:   request.Number,
	}

	location, err := integrationGeolocation(requestLocation, 1)

	if err != nil {
		return "", "", err
	}

	return location.Lat, location.Lng, nil
}

func GetLocationDeposit(request model.Deposit) (string, string, error) {
	requestLocation := &communicate.GelocationRequest{
		Street:   request.Street,
		District: request.District,
		City:     request.City,
		Country:  request.Country,
		ZipCode:  request.ZipCode,
		State:    request.State,
		Number:   request.Number,
	}

	location, err := integrationGeolocation(requestLocation, 1)

	if err != nil {
		return "", "", err
	}

	return location.Lat, location.Lng, nil
}
