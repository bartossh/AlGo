package geomath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJson(t *testing.T) {
	j, err := ReadJsonToGeoCollection("countries.geojson")
	assert.Nil(t, err)
	assert.NotNil(t, j)
}
