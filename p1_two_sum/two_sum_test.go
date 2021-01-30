package p1_two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums=[0,1,2,3,4,5] target=0",
			args: args{nums: []int{0, 1, 2, 3, 4, 5}, target: 0},
			want: []int{},
		},
		{
			name: "nums=[0,1,2,3,4,5] target=1",
			args: args{nums: []int{0, 1, 2, 3, 4, 5}, target: 1},
			want: []int{0, 1},
		},
		{
			name: "nums=[0,3,4,0] target=0",
			args: args{nums: []int{0, 3, 4, 0}, target: 0},
			want: []int{0, 3},
		},
		{
			name: "nums=[-1,-2,-3,-4,-5] target=-8",
			args: args{nums: []int{-1, -2, -3, -4, -5}, target: -8},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Logf("got : %#v", got)
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
