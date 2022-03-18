package service

import (
	"context"
	"micro-logistic/communicate"

	"google.golang.org/grpc"
)

const attemptRetryClient = 2

func integrationClientById(c *communicate.ValidateClientByIdRequest, retry int) (*communicate.ValidateClientByIdResponse, error) {
	ctx := context.Background()
	connGeolocation, err := grpc.Dial(":6000", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer connGeolocation.Close()

	serviceLocation := communicate.NewClientCommunicateClient(connGeolocation)

	client, err := serviceLocation.ValidateClientById(ctx, c)

	if err != nil {
		return attempRetryLatencyClient(retry, client, err, c)
	}

	return client, nil
}

func attempRetryLatencyClient(retry int, requestClient *communicate.ValidateClientByIdResponse, err error, c *communicate.ValidateClientByIdRequest) (*communicate.ValidateClientByIdResponse, error) {
	retry += 1
	if retry <= attemptRetryClient {
		return integrationClientById(c, retry)
	}
	return requestClient, err
}

func ValidateClientById(idClient int64) error {

	validateClientById := &communicate.ValidateClientByIdRequest{
		IdClient: idClient,
	}

	_, err := integrationClientById(validateClientById, 1)

	return err
}
