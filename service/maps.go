package service

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

var keyGoogle = "AIzaSyBPBrahfw9qmxMQAtTtDI54qBpjgF4I6wA"

type LatAndLng struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Calculate struct {
	Origin  LatAndLng `json:"lat_and_lng_origin`
	Destiny LatAndLng `json:"lat_and_lng_destiny`
}

func (calc *Calculate) CalculateRoute() error {

	c, err := maps.NewClient(maps.WithAPIKey(keyGoogle))

	if err != nil {
		return err
	}

	r := &maps.DirectionsRequest{
		Origin:      "Sydney",
		Destination: "Perth",
	}

	route, _, err := c.Directions(context.Background(), r)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fmt.Println(route)
	return nil
}
