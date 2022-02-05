package service

import (
	"context"
	"fmt"
	"time"

	"googlemaps.github.io/maps"
)

var keyGoogle = "AIzaSyCeajRVwvBKwxyQRRyMHOx4zVWzk1ETFuU"

type LatAndLng struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Calculate struct {
	Origin        LatAndLng
	Destiny       LatAndLng
	HumanReadable string
	Meters        int
	Duration      time.Duration
}

func (calc *Calculate) CalculateRoute() error {

	c, err := maps.NewClient(maps.WithAPIKey(keyGoogle))

	if err != nil {
		return err
	}

	r := &maps.DirectionsRequest{
		Origin:      fmt.Sprintf("%v, %v", calc.Origin.Lat, calc.Origin.Lng),
		Destination: fmt.Sprintf("%v, %v", calc.Destiny.Lat, calc.Destiny.Lng),
	}

	rout, _, err := c.Directions(context.Background(), r)

	if err != nil {
		return err
	}

	for _, v := range rout {
		for _, a := range v.Legs {
			calc.Meters = a.Meters
			calc.HumanReadable = a.HumanReadable
			calc.Duration = a.Duration
		}
	}
	return nil
}
