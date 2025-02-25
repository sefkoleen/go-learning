package main

func calcSum(values ...int) int {
	sumValue := 0
	for _, value := range values {
		sumValue += value
	}
	return sumValue
}

func calcSub(values ...int) int {
	subValue := values[0]
	for i, value := range values {
		if i == 0 {
			continue
		}

		subValue -= value
	}
	return subValue
}