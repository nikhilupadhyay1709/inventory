package services

func divideByPossibleZero(value float32, divider int) float32 {
	var result float32
	if divider == 0 {
		result = 0
	} else {
		result = value / float32(divider)
	}

	return result
}
