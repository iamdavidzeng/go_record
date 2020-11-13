package main

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	} else {
		less := []int{}
		greater := []int{}
		pivot := nums[0]

		for _, value := range nums[1:] {
			if value <= pivot {
				less = append(less, value)
			} else {
				greater = append(greater, value)
			}
		}

		result := append(quickSort(less), pivot)
		result = append(result, quickSort(greater)...)

		return result
	}
}

func main() {
	nums := []int{5, 2, 21, 4, 1}
	fmt.Println(quickSort(nums))
}
