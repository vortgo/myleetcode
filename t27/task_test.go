package t27

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"f1",
			args{nums: []int{1, 2, 3, 4, 5, 6, 6, 7, 4, 5, 6, 7, 6, 5, 3, 6, 7}, val: 6},
			12,
		},
		{
			"f2",
			args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			5,
		},
		{
			"f3",
			args{nums: []int{1}, val: 1},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
