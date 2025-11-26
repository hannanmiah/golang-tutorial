package functions

import "math"

func Sqrt(a float64) (float64,error) {
	if a < 0 {
		return 0, Error{"negative number"}
	}
	return math.Sqrt(a), nil
}

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}