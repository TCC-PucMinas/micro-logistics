package controller

import (
	"context"
	"log"
	"micro-logistic/communicate"
	model "micro-logistic/models"
	"micro-logistic/service"
)

type CourierRoutes struct{}

func (s *CourierRoutes) CourierRouteListAll(ctx context.Context, request *communicate.CourieRouteAllRequest) (*communicate.CourierRouteAllResponse, error) {
	res := &communicate.CourierRouteAllResponse{}

	var courierRoute model.CourierRoute

	arrayCourierRoute, total, err := courierRoute.GetCourierRoutesPaginate(request.Delivered, request.Page, request.Limit)

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

		deposit := model.Deposit{}
		_ = deposit.GetById(courierGet.Deposit.Id)
		client, _ := service.ListOneClientById(courierGet.Client.Id)
		product, _ := service.ListOneProductById(courierGet.Product.Id)

		//
		courier.Id = courierGet.Id
		courier.Client = &communicate.ClientCr{Id: client.Client.Id, Name: client.Client.Name}
		courier.Deposit = &communicate.DepositCr{Id: deposit.Id, Name: deposit.Name}
		courier.Product = &communicate.ProductCr{Id: product.Product.Id, Name: product.Product.Name}
		courier.Delivered = courierGet.Delivered
		courier.Doc = &communicate.Doc{Type: courierGet.Doc.Type, Value: courierGet.Doc.Value}
		data.CourierRoute = append(data.CourierRoute, courier)
	}

	res.Total = total
	res.Page = request.Page
	res.Limit = request.Limit

	res.Data = data
	return res, nil
}

func (s *CourierRoutes) CreateCorrierRoute(ctx context.Context, request *communicate.CreateCourierRoutesRequest) (*communicate.CreateCourierRoutesResponse, error) {

	log.Println(request.IdDriver)
	res := &communicate.CreateCourierRoutesResponse{}

	for _, v := range request.Courriers {

		courrier := model.Courier{}

		if err := courrier.GetById(v.Id); err != nil {
			return res, err
		}

		destiny, err := service.ListOneDestinationByIdProduct(courrier.Product.Id)

		if err != nil {
			return res, err
		}

		if err := courrier.Deposit.GetById(courrier.Deposit.Id); err != nil {
			return res, err
		}

		courierRoute := model.CourierRoute{
			Courier: model.Courier{Id: v.Id},
		}

		if err := courierRoute.GetCourierRouteByIdCourier(); err != nil || courierRoute.Id > 0 {
			continue
		}

		courierRoute = model.CourierRoute{
			Courier:   model.Courier{Id: v.Id},
			Order:     0,
			Driver:    model.Driver{Id: request.IdDriver},
			LatInit:   model.LatAndLng{Lat: courrier.Deposit.Lat, Lng: courrier.Deposit.Lng},
			LatFinish: model.LatAndLng{Lat: destiny.Destination.Lat, Lng: destiny.Destination.Lng},
		}

		if err := courierRoute.CreateCourierRoute(); err != nil {
			return res, err
		}

		route := service.NewRoutes(&courierRoute)

		_ = route.TracingRoutes()
	}

	res.Created = true
	return res, nil
}

func (s *CourierRoutes) CourierRouteListOne(ctx context.Context, request *communicate.CourierRouteListOneRequest) (*communicate.CourierRouteListOneResponse, error) {

	res := &communicate.CourierRouteListOneResponse{}

	courierRoute := model.CourierRoute{}

	if err := courierRoute.GetCourierRouteById(request.Id); err != nil {
		return res, err
	}

	courier := model.Courier{}

	if err := courier.GetById(courierRoute.Courier.Id); err != nil {
		return res, err
	}

	deposit := model.Deposit{}
	_ = deposit.GetById(courier.Deposit.Id)
	client, _ := service.ListOneClientById(courier.Client.Id)
	product, _ := service.ListOneProductById(courier.Product.Id)

	returnResponse := &communicate.CourierRoutes{}
	returnResponse.Id = courier.Id
	returnResponse.Client = &communicate.ClientCr{Id: client.Client.Id, Name: client.Client.Name}
	returnResponse.Deposit = &communicate.DepositCr{Id: deposit.Id, Name: deposit.Name}
	returnResponse.Product = &communicate.ProductCr{Id: product.Product.Id, Name: product.Product.Name}
	returnResponse.Delivered = courier.Delivered
	returnResponse.Doc = &communicate.Doc{Type: courier.Doc.Type, Value: courier.Doc.Value}
	returnResponse.Route = &communicate.Route{
		IdCourier: courierRoute.Courier.Id,
		Init:      &communicate.Coordenates{Lat: courierRoute.LatInit.Lat, Lng: courierRoute.LatInit.Lng},
		Finish:    &communicate.Coordenates{Lat: courierRoute.LatFinish.Lat, Lng: courierRoute.LatFinish.Lng},
		Id:        courierRoute.Id,
		Order:     courierRoute.Order,
	}

	res.CourierRoute = returnResponse

	return res, nil
}
