package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (o Rectangle) CalcPerimeter() float64 {
	return 2 * o.Height + 2 * o.Weight
}

func (o Rectangle) CalcArea() float64 {
	return o.Height * o.Weight
}
