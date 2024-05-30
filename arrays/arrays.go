package arrays

func buildArray(numbers []int) []int {
	var builtArray []int
	for _, v := range numbers {
		index := v
		builtArray = append(builtArray, numbers[index])
	}
	return builtArray
}

// Knowing the size of your array is better than not knowing it so use make
func buildArrayBestOnLeetCode(numbers []int) []int {
	arr := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		arr[i] = numbers[numbers[i]]
	}
	return arr
}
