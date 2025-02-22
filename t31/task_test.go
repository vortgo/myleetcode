package t31

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{nums: []int{1, 2, 3}},
			[]int{1, 3, 2},
		},

		{
			"1342",
			args{nums: []int{1, 3, 4, 2}},
			[]int{1, 4, 2, 3},
		},

		{
			"132",
			args{nums: []int{1, 3, 2}},
			[]int{2, 1, 3},
		},

		{
			"round",
			args{nums: []int{3, 2, 1}},
			[]int{1, 2, 3},
		},

		{
			"round large",
			args{nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
