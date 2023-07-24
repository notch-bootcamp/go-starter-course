package main

import "math"

type Measurable interface {
	Length() float64
}

type MesurableFloat float64
type Point struct{ X, Y float64 }
type Vector struct{ X, Y, Z float64 }

func (f *MesurableFloat) Length() float64 {
	if *f < 0 {
		return float64(-*f)
	}
	return float64(*f)
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func DoubleTheDistance(m Measurable) float64 {
	return m.Length() * 2
}

func main() {
	f := MesurableFloat(42)
	p := Point{2, 3}
	v := Vector{1, 2, 3}
	DoubleTheDistance(&p)
	DoubleTheDistance(&f)
	DoubleTheDistance(&v)
}
