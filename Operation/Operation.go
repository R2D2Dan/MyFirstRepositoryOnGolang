package operation

import "errors"

func Pluse(a, b float64) float64 {
	return a + b
}

func Minus(a, b float64) float64 {
	return a - b
}

func Devision(a, b float64) (res float64, err error) {
	err = errors.New("Ошибка деления на ноль")
	if b == 0 {
		return 0, err
	} else {
		return a / b, nil
	}
}

func Multiplication(a, b float64) float64 {
	return a * b
}
