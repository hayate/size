package size

// Size returns description of the passed size
func Size(s int) string {
	switch {
	case s < 0:
		return "negative"
	case s == 0:
		return "zero"
	case s < 10:
		return "small"
	case s < 100:
		return "big"
	case s < 1000:
		return "huge"
	}
	return "enormous"
}
