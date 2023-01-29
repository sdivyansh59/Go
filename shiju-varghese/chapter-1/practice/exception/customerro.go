package exception

import "errors"

func Divide(x, y int) (int, error) {
	if y == 0 {
		return -1,errors.New("divide by 0")
	} else {
		return (x / y), nil
	}
}