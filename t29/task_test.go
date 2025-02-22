package t29

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"+",
			args{dividend: 8, divisor: 3},
			2,
		},
		{
			"+1",
			args{dividend: 96, divisor: 3},
			32,
		},
		{
			"+2",
			args{dividend: 971923, divisor: 3},
			323974,
		},
		{
			"-1",
			args{dividend: -8, divisor: 2},
			-4,
		},
		{
			"-2",
			args{dividend: 8, divisor: -4},
			-2,
		},
		{
			"-2",
			args{dividend: 8, divisor: -4},
			-2,
		},
		{
			"-3",
			args{dividend: -8, divisor: -5},
			1,
		},
		{
			"-4",
			args{dividend: -1021989372, divisor: -82778243},
			12,
		},
		{
			"5",
			args{dividend: 2147483647, divisor: 2},
			1073741823,
		},
		{
			"5",
			args{dividend: 1060849722, divisor: 99958928},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
