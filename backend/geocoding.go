package main

import (
	"fmt"

	"github.com/rubenv/opencagedata"
)

func geocode(city, state, country string) (string, string, error) {
	geocoder := opencagedata.NewGeocoder("89d5a7e1287b40b9b8418e3e7775e054")

	query := fmt.Sprintf("%s, %s, %s", city, state, country)
	result, err := geocoder.Geocode(query, nil)
	if err != nil {
		return "", "", err
	}

	if len(result.Results) > 0 {
		f_result := result.Results[0]
		latitude := fmt.Sprintf("%.7f", f_result.Geometry.Latitude)
		longitude := fmt.Sprintf("%.7f", f_result.Geometry.Longitude)
		return latitude, longitude, nil
	}

	return "", "", fmt.Errorf("No results found for query: %s", query)
}
