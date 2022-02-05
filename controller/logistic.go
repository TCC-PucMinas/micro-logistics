package controller

import (
	"context"

	"github.com/TCC-PucMinas/micro-logistics/communicate"
)

type LogisticServer struct {
}

func (s *LogisticServer) CalculateLogistic(ctx context.Context, request *communicate.CalulateRequest) (*communicate.CalculateResponse, error) {

	res := &communicate.CalculateResponse{}

	return res, nil
}
