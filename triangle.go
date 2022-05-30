package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (o Triangle) CalcPerimeter() float64 {
	return 3 * o.Side
}

func (o Triangle) CalcArea() float64 {
	return o.Side * o.Side * math.Sqrt(3) / 4
}
