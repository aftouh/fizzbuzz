package core

import (
	"errors"
	"strconv"
)

func Fizzbuzz(int1, int2, limit int, str1, str2 string) ([]string, error) {
	if int1 <= 0 || int2 <= 0 {
		return nil, errors.New("int1 and int2 must be greater than 0")
	}
	if limit < 1 {
		return nil, errors.New("limit must be great or equal to 1")
	}

	var res []string
	res = append(res, "1")
	mutiple := int1 * int2

	for i := 2; i <= limit; i++ {
		switch {
		case i%mutiple == 0:
			res = append(res, str1+str2)
		case i%int1 == 0:
			res = append(res, str1)
		case i%int2 == 0:
			res = append(res, str2)
		default:
			res = append(res, strconv.Itoa(i))
		}
	}

	return res, nil
}
