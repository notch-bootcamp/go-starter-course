package types

import "math"

type (
	Coordinate float32
)

type Vector3D struct {
	Label               string
	Latitude, Longitude Coordinate
	Elevation           float32
}

func (v Vector3D) ToSphericalCoordinates() (float64, float64) {
	lat := (float64(v.Latitude) * math.Pi) / 180
	long := (float64(v.Longitude) * math.Pi) / 180
	return lat, long
}

func (v *Vector3D) ChangeLabel(newLabel string) *Vector3D {
	v.Label = newLabel
	return v
}

func (v *Vector3D) ChangeCoordinates(latitude, longitude Coordinate) *Vector3D {
	v.Latitude = latitude
	v.Longitude = longitude
	return v
}

func (v *Vector3D) ChangeElevation(elevation float32) *Vector3D {
	v.Elevation = elevation
	return v
}
