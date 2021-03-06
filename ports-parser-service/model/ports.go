package model

import (
	pp "ports-parser-service/grpc"
)

// Port type
type Port struct {
	PortCode    string        `json:"port_code"`
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

// ProcessPortsRequest type
type ProcessPortsRequest struct {
	File string `json:"file"`
	Size int    `json:"size`
}

// ProcessPortsResponse type
type ProcessPortsResponse struct {
	File            string `json:"file"`
	NumberOfRecords int    `json:"number_of_records"`
}

// ToPortResponse converts from gRPC model to REST client facing type
func ToPortResponse(p *pp.Port) Port {
	return Port{City: p.City, Country: p.Country,
		Name: p.Name, Coordinates: []float64{p.Longitude, p.Latitude}, Timezone: p.Timezone, Province: p.Province, PortCode: p.PortCode, Code: p.Code}
}

// ToPbRequest convets from REST client Request to gRPC request
func ToPbRequest(p Port) *pp.PortRequest {
	var longtitude, latitude float64
	if len(p.Coordinates) > 0 {
		longtitude = p.Coordinates[0]
		latitude = p.Coordinates[1]
	}
	return &pp.PortRequest{PortCode: p.PortCode, City: p.City, Longitude: longtitude, Latitude: latitude, Province: p.Province, Timezone: p.Timezone, Code: p.Code, Country: p.Country, Name: p.Name}
}
