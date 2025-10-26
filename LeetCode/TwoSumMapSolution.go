package main

func twoSumMap2(numbers []int, target int) []int {
	trueTable := make(map[int]int)
	for key, value := range numbers {
		_, ok := trueTable[value]
		if ok {
			return []int{trueTable[value], key}
		}
		trueTable[target-value] = key
	}
	return nil
}

func main() {
	twoSumMap2([]int{2, 7, 11, 15}, 9)
}
