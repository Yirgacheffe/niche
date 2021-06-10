package fuzz

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertDollarsToPennies(amount string) (int64, error) {
	v, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	if v > 1000 && v < 1100 {
		panic("invalid range")
	}

	grps := strings.Split(amount, ".")
	result := grps[0]

	// after "."
	r := ""
	if len(grps) == 2 {
		if len(grps[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = grps[1]
		if len(r) > 2 {
			r = r[:2]
		}
	}

	for len(r) < 2 {
		r += "0"
	}

	result += r
	return strconv.ParseInt(result, 10, 64)
}
