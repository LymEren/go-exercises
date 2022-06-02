package main

import "fmt"

func numFreq(numbers []int) int {

	var countMax int
	var result int
	arrMap := map[int]int{}

	for _, value := range numbers {
		arrMap[value]++
		if arrMap[value] > countMax {
			countMax = arrMap[value]
			result = value
		}
	}
	return result
}

func main() {
	arr := []int{2, 2, 6, 7, 7, 9, 9, 9}
	fmt.Println(numFreq(arr))
}
