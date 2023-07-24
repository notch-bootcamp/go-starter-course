package types

type Position struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
	Elevation int     `json:"elev,omitempty"`
}
