package badkitty

import "math"

// Round will well Round
func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

// ToFixed ..
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundAlt(num*output)) / output
}

// RoundAlt ..
func RoundAlt(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
