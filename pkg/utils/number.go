package utils

import "math"

func RoundToTwoDecimal(f float64) float64 {
	return math.Round(f*100) / 100
}
