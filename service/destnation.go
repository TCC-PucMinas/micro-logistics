package service

import (
	"context"
	"micro-logistic/communicate"

	"google.golang.org/grpc"
)

const attemptRetryDestination = 2

func integrationDestinationById(c *communicate.ListOneDestinationByIdRequest, retry int) (*communicate.ListOneDestinationByIdResponse, error) {
	ctx := context.Background()
	connGeolocation, err := grpc.Dial(":6000", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer connGeolocation.Close()

	serviceLocation := communicate.NewDestinationCommunicateClient(connGeolocation)

	client, err := serviceLocation.ListOneDestinationById(ctx, c)

	if err != nil {
		return attempRetryLatencyDestination(retry, client, err, c)
	}

	return client, nil
}

func attempRetryLatencyDestination(retry int, requestDestination *communicate.ListOneDestinationByIdResponse, err error, c *communicate.ListOneDestinationByIdRequest) (*communicate.ListOneDestinationByIdResponse, error) {
	retry += 1
	if retry <= attemptRetryDestination {
		return integrationDestinationById(c, retry)
	}
	return requestDestination, err
}

func ListOneDestinationById(idDestination int64) (*communicate.ListOneDestinationByIdResponse, error) {

	validateClientById := &communicate.ListOneDestinationByIdRequest{
		Id: idDestination,
	}

	return integrationDestinationById(validateClientById, 1)
}

func attempRetryLatencyDestinationByProduct(retry int, requestDestination *communicate.ListOneDestinationByIdResponse, err error, c *communicate.ListOneProductByIdProductRequest) (*communicate.ListOneDestinationByIdResponse, error) {
	retry += 1
	if retry <= attemptRetryDestination {
		return integrationDestinationByIdProduct(c, retry)
	}
	return requestDestination, err
}

func integrationDestinationByIdProduct(c *communicate.ListOneProductByIdProductRequest, retry int) (*communicate.ListOneDestinationByIdResponse, error) {
	ctx := context.Background()
	connGeolocation, err := grpc.Dial(":6000", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer connGeolocation.Close()

	serviceLocation := communicate.NewDestinationCommunicateClient(connGeolocation)

	client, err := serviceLocation.ListOneProductByIdProduct(ctx, c)

	if err != nil {
		return attempRetryLatencyDestinationByProduct(retry, client, err, c)
	}

	return client, nil
}

func ListOneDestinationByIdProduct(idProduct int64) (*communicate.ListOneDestinationByIdResponse, error) {

	validateClientById := &communicate.ListOneProductByIdProductRequest{
		IdProduct: idProduct,
	}

	return integrationDestinationByIdProduct(validateClientById, 1)
}
