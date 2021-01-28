package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}

	for first, v := range nums {
		if v >= target {
			continue
		}
		k := target - v
		second, ok := m[k]
		if !ok {
			continue
		}

		if first == second {
			continue
		}
		return []int{first, second}
	}

	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(twoSum(nums, 12))
}
