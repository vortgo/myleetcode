package t15

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{nums: []int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"2",
			args{nums: []int{1, 2, -2, -1}},
			[][]int{},
		},
		{
			"0",
			args{nums: []int{0, 0, 0}},
			[][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
