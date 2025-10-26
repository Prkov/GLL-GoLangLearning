package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for zindex, zvalue := range nums {
			if value+zvalue == target && index != zindex {
				return []int{index, zindex}
			}
		}
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
