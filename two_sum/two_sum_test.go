package main

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5}
	t.Logf("nums: %#v\n", nums)
	t.Run("Target=1", func(t *testing.T) {
		r := twoSum(nums, 1)
		t.Logf("r: %#v\n", r)
		if r[0] != 0 {
			t.Fatal("result not match")
		}
		if r[1] != 1 {
			t.Fatal("result not match")
		}
	})
	t.Run("Target=10", func(t *testing.T) {
		r := twoSum(nums, 10)
		t.Logf("r: %#v\n", r)

		if len(r) != 0 {
			t.Fatal("result should be nil")
		}
	})

	t.Run("Target=11", func(t *testing.T) {
		r := twoSum(nums, 11)
		t.Logf("r: %#v\n", r)

		if len(r) != 0 {
			t.Fatal("result should be nil")
		}
	})

}
