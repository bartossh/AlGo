// Package geomath provides basic geo math functionalities.
// In external geojson Coordinates are in form of type is in form of [][][][]float64.
// To avoid unnecessary type transformation and value copying
// Simplified types are defined, such as Point, Line and Polygon as a slice.
// This allow to pass slices to functions (pointer to underlining arrays),
// no values are copied so all the calculations are very fast.
// This package aims to have maximum performance.
package geomath

import (
	"errors"
	"fmt"

	geo "github.com/kellydunn/golang-geo"
)

type orientation int8

const (
	collinear        orientation = iota
	clockwise        orientation = iota
	counterclockwise orientation = iota
)

// Point is 2d point representing edge of polygon or end of a line
// order is {lon, lat}
type Point []float64

// Line is 2d line representing side of polygon or end of a line
type Line [][]float64

// Polygon is a multi vertices object
// order is {lon, lat}
type Polygon [][][]float64

// ToPrimitive returns polygon as a slice of primitive type
func (p Polygon) ToPrimitive() [][][]float64 {
	return p
}

func (p Point) IsInDistance(cp Point, d float64) (bool, error) {
	if len(p) != 2 {
		return false, errors.New("invalid data format")
	}
	pnt0 := geo.NewPoint(p[1], p[0])
	pnt1 := geo.NewPoint(cp[1], cp[0])

	calcDist := pnt0.GreatCircleDistance(pnt1) * 1000 // [ km ] to [ m ]
	return calcDist < d, nil
}

// GeneratePointsInDist generates given n number of points evenly spaced in circle with d [ m ] distance.
// If point is not within the Polygon then it is not included in the result.
func (p Polygon) GeneratePointsInDist(d float64, n int, pn Point) ([]Point, error) {
	if len(pn) != 2 {
		return nil, fmt.Errorf("expected point has two values, lon and lat, given point has %v value(s)", len(p))
	}
	d = d / 1000 // m to km
	pnt := geo.NewPoint(pn[1], pn[0])
	step := float64(360 / n)

	pnts := make([]Point, 0, n)

	for b := 0.0; b < 360.0; b += step {
		point := pnt.PointAtDistanceAndBearing(d, b)
		pp := []float64{point.Lng(), point.Lat()}
		isWithin, err := p.Within([][][]float64{{pp}})
		if err != nil {
			return nil, err
		}
		if isWithin {
			pnts = append(pnts, pp)
		}
	}

	return pnts, nil
}

// Within verifies if given q Polygon is within the Polygon.
// Will not work for Arctic and Antarctic as it makes some simplifications.
// Intersection is calculated in latitude direction,
// assuming that extreme value of latitude equals always 90 deg.
func (p Polygon) Within(q Polygon) (bool, error) {
	if len(p) == 0 || len(q) == 0 {
		return false, errors.New("no polygons present")
	}
	if len(p[0]) < 3 {
		return false, errors.New("receiver struct is not a polygon")
	}
	if len(q[0]) == 0 {
		return false, errors.New("given polygon is empty")
	}

	intersections := 0
	for _, qPoint := range q[0] {
		if len(qPoint) != 2 {
			return false, errors.New("point should have to coordinates")
		}
		polygon := p[0]
		for i := 0; i < len(polygon)-1; i++ {
			j := i + 1
			l1 := Line{polygon[i], polygon[j]}
			l2 := Line{qPoint, {qPoint[0], 90}}
			ok, err := doIntersect(l1, l2)
			if err != nil {
				return false, err
			}
			if ok {
				intersections++
			}
		}
	}
	if intersections%2 != 0 {
		return true, nil
	}

	return false, nil
}

func onLine(p, q, r Point) bool {
	if len(p) != 2 || len(q) != 2 || len(r) != 2 {
		return false
	}

	distLon := r[0] - p[0]
	distLat := r[1] - p[1]

	dLon := distLon / (q[0] - p[0])
	dLat := distLat / (q[1] - p[1])

	absDistLon := distLon
	if distLon < 0 {
		absDistLon = -distLon
	}
	absDistLat := distLat
	if distLat < 0 {
		absDistLat = -distLat
	}

	isOn := dLon == dLat

	betweenLon := 0.0 <= dLon && dLon <= absDistLon
	betweenLat := 0.0 <= dLat && dLat <= absDistLat

	return isOn && betweenLon && betweenLat
}

func calculateOrientation(p, q, r Point) (orientation, error) {
	if len(p) != 2 || len(q) != 2 || len(r) != 2 {
		return collinear, errors.New("line should have two points, longitude and latitude")
	}
	v := (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
	if v == 0 {
		return collinear, nil
	}
	if v > 0 {
		return clockwise, nil
	}
	return counterclockwise, nil
}

func doIntersect(l1, l2 [][]float64) (bool, error) {
	if len(l1) != 2 || len(l2) != 2 {
		return false, errors.New("line should have two points, longitude and latitude")
	}

	o1, err := calculateOrientation(l1[0], l1[1], l2[0])
	if err != nil {
		return false, err
	}
	o2, err := calculateOrientation(l1[0], l1[1], l2[1])
	if err != nil {
		return false, err
	}
	o3, err := calculateOrientation(l2[0], l2[1], l1[0])
	if err != nil {
		return false, err
	}
	o4, err := calculateOrientation(l2[0], l2[1], l1[1])
	if err != nil {
		return false, err
	}

	if o1 != o2 && o3 != o4 {
		return true, nil
	}

	if o1 == 0 && onLine(l1[0], l2[0], l1[1]) {
		return true, nil
	}
	if o2 == 0 && onLine(l1[0], l2[1], l1[1]) {
		return true, nil
	}
	if o3 == 0 && onLine(l2[0], l1[0], l2[1]) {
		return true, nil
	}
	if o4 == 0 && onLine(l2[0], l1[1], l2[1]) {
		return true, nil
	}

	return false, nil
}
