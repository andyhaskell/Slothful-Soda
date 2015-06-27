package main

type Location struct {
	Id   int
	Name string
	Lat  float64
	Lng  float64
}

func initLocation(name string, lat, lng float64) *Location {
	return &Location{Name: name, Lat: lat, Lng: lng}
}
