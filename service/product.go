package service

import (
	"context"
	"micro-logistic/communicate"

	"google.golang.org/grpc"
)

const attemptRetryProduct = 2

func integrationProductById(c *communicate.ListOneProductByIdRequest, retry int) (*communicate.ListOneProductByIdResponse, error) {
	ctx := context.Background()
	connGeolocation, err := grpc.Dial(":6000", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer connGeolocation.Close()

	serviceLocation := communicate.NewProductCommunicateClient(connGeolocation)

	client, err := serviceLocation.ListOneProductById(ctx, c)

	if err != nil {
		return attempRetryLatencyProduct(retry, client, err, c)
	}

	return client, nil
}

func attempRetryLatencyProduct(retry int, requestClient *communicate.ListOneProductByIdResponse, err error, c *communicate.ListOneProductByIdRequest) (*communicate.ListOneProductByIdResponse, error) {
	retry += 1
	if retry <= attemptRetryClient {
		return integrationProductById(c, retry)
	}
	return requestClient, err
}

func ListOneProductById(idProduct int64) (*communicate.ListOneProductByIdResponse, error) {

	validateClientById := &communicate.ListOneProductByIdRequest{
		Id: idProduct,
	}

	client, err := integrationProductById(validateClientById, 1)

	return client, err
}
