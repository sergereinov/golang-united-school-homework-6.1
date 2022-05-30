package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (o Circle) CalcPerimeter() float64 {
	return 2 * o.Radius * math.Pi
}

func (o Circle) CalcArea() float64 {
	return o.Radius * o.Radius * math.Pi
}
