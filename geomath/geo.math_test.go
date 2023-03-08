package geomath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type lineCase struct {
	isOn  bool
	line  [][]float64
	point []float64
	name  string
}

type intersectCase struct {
	isErr        bool
	intersect    bool
	line1, line2 [][]float64
	name         string
}

type withinCase struct {
	isErr            bool
	isWithin         bool
	polygonEnclosing []Polygon
	polygonCoords    Polygon
	name             string
}

func TestOnLine(t *testing.T) {
	cases := []lineCase{
		{
			isOn:  false,
			line:  [][]float64{{1.1, 2.2}, {2.2, 2.2}},
			point: []float64{1.0, 2.2},
			name:  "on vertical line extension, bot not between points, returns false",
		},
		{
			isOn:  false,
			line:  [][]float64{{1.1, 2.2}, {2.2, 1.1}},
			point: []float64{1.1, 2.0},
			name:  "on horizontal  line extension, bot not between points, returns false",
		},
		{
			isOn:  false,
			line:  [][]float64{{1.1, 2.2}, {2.2, 2.2}},
			point: []float64{1.0, 2.2},
			name:  "on line extension, bot not between points, returns false",
		},
		{
			isOn:  false,
			line:  [][]float64{{1.1, 2.2}, {2.2, 2.2}},
			point: []float64{1.1, 2.3},
			name:  "on line extension, bot not between points, returns false",
		},
		{
			isOn:  false,
			line:  [][]float64{{1.1, 2.2}, {2.2, 3.2}},
			point: []float64{1.1, 2.3},
			name:  "not on line, returns false",
		},
		{
			isOn:  false,
			line:  [][]float64{{1.0, 1.0}, {3.0, 2.0}},
			point: []float64{1.1, 2.3},
			name:  "not on line, returns false",
		},
		{
			isOn:  false,
			line:  [][]float64{{1.0, 1.0}, {3.0, 2.0}},
			point: []float64{2.0, 2.6},
			name:  "not on line, returns false",
		},
		{
			isOn:  true,
			line:  [][]float64{{1.0, 1.0}, {3.0, 3.0}},
			point: []float64{2.0, 2.0},
			name:  "on line, returns true",
		},
		{
			isOn:  true,
			line:  [][]float64{{3.0, 3.0}, {1.0, 1.0}},
			point: []float64{2.0, 2.0},
			name:  "on line, returns true",
		},
		{
			isOn:  true,
			line:  [][]float64{{1.0, 1.0}, {2.0, 3.0}},
			point: []float64{2.0, 3.0},
			name:  "on vertex, returns true",
		},
		{
			isOn:  true,
			line:  [][]float64{{0.0, 0.0}, {4.0, 6.0}},
			point: []float64{2.0, 3.0},
			name:  "on line, returns true",
		},
		{
			isOn:  true,
			line:  [][]float64{{4.0, 6.0}, {0.0, 0.0}},
			point: []float64{2.0, 3.0},
			name:  "on line, returns true",
		},
		{
			isOn:  true,
			line:  [][]float64{{4.0, 0.0}, {0.0, 6.0}},
			point: []float64{2.0, 3.0},
			name:  "on line, returns true",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %s-%v", c.name, i), func(t *testing.T) {
			ok := onLine(c.line[0], c.point, c.line[1])
			assert.Equal(t, c.isOn, ok)
		})
	}
}

func TestIntersect(t *testing.T) {
	cases := []intersectCase{
		{
			intersect: false,
			isErr:     true,
			line1:     [][]float64{{1.1, 2.2}, {2.2, 2.2}},
			line2:     [][]float64{},
			name:      "error, corrupted line coordinates",
		},
		{
			intersect: false,
			isErr:     true,
			line1:     [][]float64{},
			line2:     [][]float64{{1.1, 2.2}, {2.2, 2.2}},
			name:      "error, corrupted line coordinates",
		},
		{
			intersect: true,
			isErr:     false,
			line1:     [][]float64{{5.1, 5.0}, {10.0, 10.0}},
			line2:     [][]float64{{0.0, 0.0}, {10.0, 10.0}},
			name:      "do intersect",
		},
		{
			intersect: true,
			isErr:     false,
			line1:     [][]float64{{5.0, 5.1}, {10.0, 10.0}},
			line2:     [][]float64{{0.0, 0.0}, {10.0, 10.0}},
			name:      "do intersect",
		},
		{
			intersect: true,
			isErr:     false,
			line1:     [][]float64{{4.0, 3.0}, {8.0, 10.0}},
			line2:     [][]float64{{0.0, 0.0}, {10.0, 10.0}},
			name:      "do intersect",
		},
		{
			intersect: true,
			isErr:     false,
			line1:     [][]float64{{8.0, 0.0}, {4.0, 6.0}},
			line2:     [][]float64{{0.0, 0.0}, {10.0, 10.0}},
			name:      "do intersect",
		},
		{
			intersect: false,
			isErr:     false,
			line1:     [][]float64{{4.0, 0.0}, {6.0, 3.0}},
			line2:     [][]float64{{0.0, 0.0}, {10.0, 10.0}},
			name:      "do not intersect",
		},
		{
			intersect: false,
			isErr:     false,
			line1:     [][]float64{{4.0, 5.0}, {8.0, 9.0}},
			line2:     [][]float64{{0.0, 0.0}, {10.0, 10.0}},
			name:      "do not intersect",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %s-%v", c.name, i), func(t *testing.T) {
			ok, err := doIntersect(c.line1, c.line2)
			if c.isErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, c.intersect, ok)
			}
		})
	}
}

func TestIsWithin(t *testing.T) {

	polygon0 := []Polygon{{{
		{53.121491, 7.505822},
		{53.839609, 13.041928},
		{52.062728, 14.205001},
		{50.540540, 12.109173},
		{47.961690, 11.950758},
		{48.345232, 8.330028},
		{49.096488, 8.918092},
		{49.574821, 6.684059},
		{53.121491, 7.505822},
	}}}

	polygon1 := []Polygon{{{
		{50.177267, 47.482085},
		{54.059077, 84.923223},
		{42.680974, 87.735015},
		{41.195535, 45.959942},
		{50.177267, 47.482085},
	}}}

	gojsonCollection, err := ReadJsonToGeoCollection("countries.geojson")
	assert.Nil(t, err)

	geojsonPolygon := make([]Polygon, 0)

	for _, feature := range gojsonCollection.Features {
		if feature.Properties.Admin == "Germany" {
			geojsonPolygon, err = feature.ReadCoordinates()
			assert.Nil(t, err)
			fmt.Printf("length of the polygon %v\n", len(geojsonPolygon))
		}
	}

	cases := []withinCase{
		{ // 0
			isErr:            true,
			isWithin:         false,
			polygonEnclosing: geojsonPolygon,
			polygonCoords:    [][][]float64{},
			name:             "error polygon is empty",
		},
		{ // 1
			isErr:            true,
			isWithin:         false,
			polygonEnclosing: geojsonPolygon,
			polygonCoords:    [][][]float64{{{}}},
			name:             "error polygon is empty",
		},
		{ // 2
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: polygon1,
			polygonCoords:    [][][]float64{{{47.272921, 67.421330}}},
			name:             "is Within",
		},
		{ // 3
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: polygon1,
			polygonCoords:    [][][]float64{{{47.972921, 67.221330}}},
			name:             "is Whithin",
		},
		{ // 4
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: polygon1,
			polygonCoords:    [][][]float64{{{47.372921, 67.921330}}},
			name:             "is Whithin",
		},
		{ // 5
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: polygon1,
			polygonCoords:    [][][]float64{{{47.112921, 67.111330}}},
			name:             "is Whithin",
		},
		{ // 6
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: geojsonPolygon,
			polygonCoords:    [][][]float64{{{7.808848, 50.491769}}},
			name:             "is Whithin",
		},
		{ // 7
			isErr:            false,
			isWithin:         false,
			polygonEnclosing: geojsonPolygon,
			polygonCoords:    [][][]float64{{{40.605726, 64.761661}}},
			name:             "isn't Whithin",
		},
		{ // 8
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: geojsonPolygon,
			polygonCoords:    [][][]float64{{{10.401101, 49.406767}}},
			name:             "is Whithin",
		},
		{ // 9
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: polygon0,
			polygonCoords:    [][][]float64{{{50.491769, 7.808848}}},
			name:             "is Whithin",
		},
		{ // 10
			isErr:            false,
			isWithin:         true,
			polygonEnclosing: polygon0,
			polygonCoords:    [][][]float64{{{49.406767, 10.401101}}},
			name:             "is Whithin",
		},
		{ // 11
			isErr:            false,
			isWithin:         false,
			polygonEnclosing: polygon0,
			polygonCoords:    [][][]float64{{{47.899398, 1.553708}}},
			name:             "isn't Whithin",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %s %v", c.name, i), func(t *testing.T) {
			var ok bool
			var err error
			for _, pl := range c.polygonEnclosing {
				ok, err = pl.Within(c.polygonCoords)
				if ok || err != nil {
					break
				}
			}
			if c.isErr {
				assert.NotNil(t, err)
				assert.False(t, ok)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, c.isWithin, ok)
			}
		})
	}
}

func BenchmarkWithIn(b *testing.B) {
	b.ReportAllocs()

	b.StopTimer()
	gojsonCollection, err := ReadJsonToGeoCollection("countries.geojson")
	assert.Nil(b, err)

	geojsonPolygon := make([]Polygon, 0)

	for _, feature := range gojsonCollection.Features {
		if feature.Properties.Admin == "Germany" {
			geojsonPolygon, err = feature.ReadCoordinates()
			assert.Nil(b, err)
		}
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, pl := range geojsonPolygon {
			_, err := pl.Within([][][]float64{{{10.401101, 49.406767}}})
			assert.Nil(b, err)
		}
	}
}

func TestGeneratePointsInDist(t *testing.T) {
	j, err := ReadJsonToGeoCollection("countries.geojson")
	assert.Nil(t, err)

	var geoPolygon []Polygon

	for _, feature := range j.Features {
		if feature.Properties.Admin == "Germany" {
			geoPolygon, err = feature.ReadCoordinates()
			assert.Nil(t, err)
			break
		}
	}

	var found bool
	for _, plg := range geoPolygon {
		pnts, err := plg.GeneratePointsInDist(1000, 24, []float64{13.7251, 53.8709})
		assert.Nil(t, err)
		if len(pnts) > 0 {
			found = true
			assert.Equal(t, 24, len(pnts))
			break
		}
	}
	assert.True(t, found)
}

func BenchmarkGeneratePointsInDist(b *testing.B) {
	b.ReportAllocs()

	b.StopTimer()
	j, err := ReadJsonToGeoCollection("countries.geojson")
	assert.Nil(b, err)

	var geoPolygon []Polygon

	for _, feature := range j.Features {
		if feature.Properties.Admin == "Germany" {
			geoPolygon, err = feature.ReadCoordinates()
			assert.Nil(b, err)
			break
		}
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		for _, plg := range geoPolygon {
			pnts, err := plg.GeneratePointsInDist(1000, 24, []float64{13.7251, 53.8709})
			assert.Nil(b, err)
			if len(pnts) > 0 {
				break
			}
		}
	}
}
