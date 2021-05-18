package model

import pm "ports-manager-service/grpc"

// Port type
type Port struct {
	ID        int64
	PortCode  string
	Name      string
	City      string
	Country   string
	Alias     []interface{}
	Regions   []interface{}
	Latitude  float64
	Longitude float64
	Province  string
	Timezone  string
	Unlocs    []string
	Code      string
}

// FromPort converts from Port DAO to grpc Port type
func FromPort(p *Port) *pm.Port {
	return &pm.Port{PortCode: p.PortCode, Latitude: p.Latitude, Longitude: p.Longitude,
		Name: p.Name, City: p.City, Country: p.Country, Province: p.Province, Timezone: p.Timezone, Code: p.Code}
}

// ToPort converts from gRPC PortRequest to Port DAO
func ToPort(p *pm.PortRequest) Port {
	return Port{PortCode: p.PortCode, Latitude: p.Latitude, Longitude: p.Longitude,
		Name: p.Name, City: p.City, Country: p.Country, Timezone: p.Timezone, Code: p.Code, Province: p.Province}
}
