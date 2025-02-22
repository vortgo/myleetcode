package t1942

import "testing"

func Test_smallestChair(t *testing.T) {
	type args struct {
		times        [][]int
		targetFriend int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"zero",
			args{times: [][]int{
				{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310},
				{78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856},
				{98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821},
				{48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467},
			}, targetFriend: 6},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestChair(tt.args.times, tt.args.targetFriend); got != tt.want {
				t.Errorf("smallestChair() = %v, want %v", got, tt.want)
			}
		})
	}
}
