package service

import (
	"context"
	"google.golang.org/grpc"
	"micro-logistic/communicate"
	"time"
)

const attemptRetryDestination = 20

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
		time.Sleep(1 * time.Second)
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
