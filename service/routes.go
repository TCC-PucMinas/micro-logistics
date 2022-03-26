package service

import (
	model "micro-logistic/models"
)

type LatAndLng struct {
	Lat string
	Lng string
}

type Routes struct {
	courierRoute *model.CourierRoute
	Origin       LatAndLng
	Destiny      LatAndLng
	IdCourier    int64
	Order        int64
}

func NewRoutes(courier *model.CourierRoute) *Routes {
	return &Routes{courierRoute: courier}
}

func (r *Routes) TracingRoutes() error {
	var arrayRoutes []Routes

	arrayCourierRoute, err := r.courierRoute.GetCourierRoutes()

	if err != nil {
		return err
	}

	if len(arrayCourierRoute) == 1 {
		return nil
	}

	for _, v := range arrayCourierRoute {
		routeAp := Routes{}
		routeAp.IdCourier = v.Id
		routeAp.Order = v.Order
		routeAp.Origin = LatAndLng(v.LatInit)
		routeAp.Destiny = LatAndLng(v.LatFinish)
		arrayRoutes = append(arrayRoutes, routeAp)
	}

	// enviar para o maps os registro para ser traçados.

	return nil
}
