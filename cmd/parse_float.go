package cmd

import "strconv"

func parseTwoFloats(a, b string) (float64, float64, error) {
	x, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}
