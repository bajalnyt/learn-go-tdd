package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberstoSum ...[]int) []int {
	result := []int{}
	for _, v := range numberstoSum {
		result = append(result, Sum(v))
	}

	return result
}

func SumAllTails(numberstoSum ...[]int) []int {
	result := []int{}
	for _, v := range numberstoSum {
		if len(v) == 0 {
			result = append(result, 0)
		} else {
			tail := v[1:]
			result = append(result, Sum(tail))
		}
	}

	return result
}
