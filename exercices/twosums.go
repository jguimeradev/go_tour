package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sum := 0

	aux := []int{}

	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			sum = v + nums[i]
			if sum == target {
				aux = append(aux, k, i)
				return aux
			}
		}
	}

	return aux
}

func main() {
	nums := []int{5, 11, 15, 7, 3, 2, 2}
	var target int = 4
	fmt.Println(twoSum(nums, target))
}
