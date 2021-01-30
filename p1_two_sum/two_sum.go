package p1_two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int, len(nums))
	for i, v := range nums {
		if len(m[v]) == 0 {
			m[v] = []int{}
		}
		m[v] = append(m[v], i)
	}

	for first, v := range nums {
		k := target - v
		second, _ := m[k]

		for _, v := range second {
			if first == v {
				continue
			}
			return []int{first, v}
		}
	}

	return []int{}
}
