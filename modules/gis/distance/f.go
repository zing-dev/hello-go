package main

import "math"

func D2(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6378137 // Earth radius in METERS https://en.wikipedia.org/wiki/Earth_radius

	// Distance between lat and longs
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	// convert to radians
	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	// apply formulae
	a := (math.Pow(math.Sin(dLat/2), 2)) + math.Pow(math.Sin(dLon/2), 2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Asin(math.Sqrt(a))

	return r * c
}
