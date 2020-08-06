package main

import (
	"context"
	"fmt"

	"googlemaps.github.io/maps"
)

func main() {
	client, err := maps.NewClient(maps.WithAPIKey("XXX"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	req := &maps.GeocodingRequest{
		Address:      "XXX",
		Components:   nil,
		Bounds:       nil,
		Region:       "",
		LatLng:       nil,
		ResultType:   nil,
		LocationType: nil,
		PlaceID:      "",
		Language:     "ja",
		Custom:       nil,
	}

	resp, err := client.Geocode(ctx, req)
	if err != nil {
		panic(err)
	}

	if len(resp) == 0 {
		panic("not found")
	}

	result := resp[0]

	for _, r := range result.AddressComponents {
		if r.Types[0] == "postal_code" {
			fmt.Println(r.LongName)
		}
	}

	fmt.Println(result.Geometry.Location.Lat)
	fmt.Println(result.Geometry.Location.Lng)
}
