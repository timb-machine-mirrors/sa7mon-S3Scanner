package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRegionsDO(t *testing.T) {
	r, err := GetRegionsDO()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(r), 1)
	assert.Contains(t, r, "nyc3")
}

func TestGetRegionsLinode(t *testing.T) {
	r, err := GetRegionsLinode()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(r), 1)
	assert.Contains(t, r, "us-east-1")
}

func TestGetRegionsDreamhost(t *testing.T) {
	dor, err := GetRegionsDreamhost()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(dor), 1)
	assert.Contains(t, dor, "us-east-1")
}