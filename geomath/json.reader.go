package geomath

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// ReadJson read json file and unmarshal it to GeoCollection
func ReadJsonToGeoCollection(filepath string) (GeoCollection, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return GeoCollection{}, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return GeoCollection{}, err
	}

	var result GeoCollection
	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

// GeoCollection represents geojson collection of features
type GeoCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

// Feature represents a single feature
type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

// Properties represents geojson feature properties
type Properties struct {
	Admin string `json:"ADMIN"`
	IsoA3 string `json:"ISO_A3"`
}

// Geometry represents geojson feature geometry
type Geometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

// ReadCoordinates reads coordinates based on geometry type
func (feature Feature) ReadCoordinates() ([]Polygon, error) {
	switch feature.Geometry.Type {
	case "MultiPolygon":
		res := []Polygon{}
		if err := json.Unmarshal(feature.Geometry.Coordinates, &res); err != nil {
			return nil, err
		}
		return res, nil
	case "Polygon":
		p := Polygon{}
		if err := json.Unmarshal(feature.Geometry.Coordinates, &p); err != nil {
			return nil, err
		}
		return []Polygon{p}, nil
	default:
		return nil, errors.New("no valid type to unmarshal")
	}
}
