package test

import (
	"math"
	"math/rand"
)

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Max returns the larger of x or y.
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min returns the smaller of x or y.
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

// RandInt returns a non-negative pseudo-random int from the default Source.
func RandInt() int {
	return rand.Int()
}

// Floor returns the greatest integer value less than or equal to x.
func Floor(x float64) float64 {
	return math.Floor(x)
}
