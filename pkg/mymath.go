package mymath

import "errors"

func Add(a int64, b int64) int64 {
	return (a + b)
}

func Subtract(a int64, b int64) int64 {
	return (a - b)
}

func Divide(a int64, b int64) (float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero not possible")
	}
	return float32(a) / float32(b), nil
}
