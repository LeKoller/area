package area

import "math"

// Pi is the math constant
const Pi = 3.1416

// Circ is responsible for calculating tha area of a circle
func Circ(axis float64) float64 {
	return math.Pow(axis, 2) * Pi
}

// Rect is responsible for calculating tha area of a rectangle
func Rect(basis, height float64) float64 {
	return basis * height
}

func _EqTriangle(basis, height float64) float64 {
	return (basis * height) / 2
}
