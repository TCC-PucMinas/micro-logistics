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
		routeAp.IdCourier = v.Courier.Id
		routeAp.Order = v.Order
		routeAp.Origin = LatAndLng(v.LatInit)
		routeAp.Destiny = LatAndLng(v.LatFinish)
		arrayRoutes = append(arrayRoutes, routeAp)
	}

	location, err := OrderRoutes(arrayRoutes)

	if err != nil {
		return err
	}

	for _, v := range location.Routes {
		routeAp := NewRoutes(r.courierRoute)
		routeAp.courierRoute.Courier.Id = v.IdCourier
		routeAp.courierRoute.Order = v.Order
		routeAp.courierRoute.LatInit.Lat = v.Origin.Lat
		routeAp.courierRoute.LatInit.Lng = v.Origin.Lng
		routeAp.courierRoute.LatFinish.Lat = v.Destiny.Lat
		routeAp.courierRoute.LatFinish.Lng = v.Destiny.Lng
		go routeAp.courierRoute.UpdateByCourierId()
	}

	return nil
}
