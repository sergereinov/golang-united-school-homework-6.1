package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var capacityError error = fmt.Errorf("out of capacity")
var indexError error = fmt.Errorf("index out of range")

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes: make([]Shape, 0, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("AddShape %w", capacityError)
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if (i >= len(b.shapes)) {
		return nil, fmt.Errorf("GetByIndex %w", indexError)
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if (i >= len(b.shapes)) {
		return nil, fmt.Errorf("ExtractByIndex %w", indexError)
	}

	s := b.shapes[i]
	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes = b.shapes[:len(b.shapes)-1]
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if (i >= len(b.shapes)) {
		return nil, fmt.Errorf("ReplaceByIndex %w", indexError)
	}

	s := b.shapes[i]
	b.shapes[i] = shape
	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var acc float64
	for _,v := range b.shapes {
		acc += v.CalcPerimeter()
	}
	return acc
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var acc float64
	for _,v := range b.shapes {
		acc += v.CalcArea()
	}
	return acc
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	oldLen := len(b.shapes)
	for i:=0; i<len(b.shapes); {
		s := b.shapes[i]
		if _, ok := s.(*Circle); ok {
			b.ExtractByIndex(i)
		} else {
			i++
		}
	}

	if oldLen == len(b.shapes) {
		return fmt.Errorf("RemoveAllCircles circle not found")
	}

	return nil //happy path
}
