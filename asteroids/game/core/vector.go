package core

import "asteroids/game/math"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Vector[Type Number] struct {
	X Type
	Y Type
}

func (vector *Vector[Number]) Normalize() Vector[Number] {
	magnitude := vector.Magnitude()
	return Vector[Number]{
		X: vector.X / magnitude,
		Y: vector.Y / magnitude,
	}
}

func (vector *Vector[Number]) Magnitude() Number {
	return Number(math.Sqrt(float64(vector.X*vector.X + vector.Y*vector.Y)))
}
