package controller

import (
	"context"
	"micro-logistic/communicate"
	model "micro-logistic/models"
)

type CourierRoutes struct{}

func (s *CourierRoutes) CourierRouteListAll(ctx context.Context, request *communicate.CourieRouteAllRequest) (*communicate.CourierRouteAllResponse, error) {
	res := &communicate.CourierRouteAllResponse{}

	var courierRoute model.CourierRoute

	arrayCourierRoute, total, err := courierRoute.GetCourierRoutesPaginate(request.Page, request.Limit)

	if err != nil {
		return res, err
	}

	data := &communicate.DataCourierRoute{}

	for _, c := range arrayCourierRoute {
		courier := &communicate.CourierRoutes{}
		courier.Route = &communicate.Route{
			IdCourier: c.Courier.Id,
			Init:      &communicate.Coordenates{Lat: c.LatInit.Lat, Lng: c.LatInit.Lng},
			Finish:    &communicate.Coordenates{Lat: c.LatFinish.Lat, Lng: c.LatFinish.Lng},
			Id:        c.Id,
			Order:     c.Order,
		}
		courierGet := model.Courier{}

		if err := courierGet.GetById(c.Courier.Id); err != nil {
			return res, err
		}

		courier.Id = courierGet.Id
		courier.IdClient = courierGet.Client.Id
		courier.IdDeposit = courierGet.Deposit.Id
		courier.IdDriver = courierGet.Driver.Id
		courier.IdProduct = courierGet.Product.Id
		data.CourierRoute = append(data.CourierRoute, courier)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}
