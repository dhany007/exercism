package lsproduct

import (
	"errors"
	"strconv"
)

func isValidProduct(digits string, span int) error {
	if digits == "" || span < 0 {
		return errors.New("invalid")
	}

	if span > len(digits) {
		return errors.New("invalid")
	}

	return nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if err := isValidProduct(digits, span); err != nil {
		return 0, err
	}

	var series []string
	i := 0
	for {
		if i+span > len(digits) {
			break
		}
		a := digits[i : i+span]
		series = append(series, a)

		i++
	}

	var large int64 = 0
	for _, s := range series {
		val, err := calculateSeries(s)
		if err != nil {
			return 0, err
		}
		if val > large {
			large = val
		}
	}
	return large, nil
}

func calculateSeries(serie string) (int64, error) {
	var sum int64 = 1

	for _, s := range serie {
		val, err := strconv.ParseInt(string(s), 10, 64)
		if err != nil {
			return 0, errors.New("digits input must only contain digits")
		}

		sum *= val
	}

	return sum, nil
}
