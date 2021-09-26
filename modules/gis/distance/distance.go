package main

import (
	"fmt"
	"github.com/asmarques/geodist"
	geo "github.com/kellydunn/golang-geo"
	"gitlab.com/psns/geocalc"
)

type Dis struct {
	Start    geocalc.Point
	End      geocalc.Point
	Distance float64
}

// Distance calculation
func Distance() {
	for _, dis := range []Dis{
		{
			//41.53786, 121.5665	40.99264, 123.1513	145.661公里 / 90.509英里
			Start:    geocalc.Point{Lat: 41.53786, Lon: 121.5665},
			End:      geocalc.Point{Lat: 40.99264, Lon: 123.1513},
			Distance: 145661,
		},
		{
			Start:    geocalc.Point{Lat: 30.551589, Lon: 114.202234},
			End:      geocalc.Point{Lat: 30.557181, Lon: 114.503553},
			Distance: 28841,
		},
	} {
		d1, _ := geodist.VincentyDistance(geodist.Point{Lat: dis.Start.Lat, Long: dis.Start.Lon}, geodist.Point{Lat: dis.End.Lat, Long: dis.End.Lon})
		d2 := geodist.HaversineDistance(geodist.Point{Lat: dis.Start.Lat, Long: dis.Start.Lon}, geodist.Point{Lat: dis.End.Lat, Long: dis.End.Lon})
		d3 := D2(dis.Start.Lat, dis.Start.Lon, dis.End.Lat, dis.End.Lon)
		d4 := geocalc.Distance(dis.Start, dis.End)
		d5 := geo.NewPoint(dis.Start.Lat, dis.Start.Lon).GreatCircleDistance(geo.NewPoint(dis.End.Lat, dis.End.Lon))
		fmt.Printf("d1:%.4f,d2:%.4f,d3:%.4f,d4:%.4f,d5:%.4f\n", d1, d2, d3/1000, d4/1000, d5)
	}
}

func main() {
	Distance()
}
