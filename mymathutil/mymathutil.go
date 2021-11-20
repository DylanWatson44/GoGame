package mymathutil

func ClampFloat64(val, low, high float64) float64 {
	if val < low {
		return low
	}
	if val > high {
		return high
	}
	return val
}
