package t5

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"first",
			args{s: "babad"},
			"bab",
		},

		{
			"my",
			args{s: "123sdfghjhgfdsoooo"},
			"sdfghjhgfds",
		},

		{
			"odd",
			args{s: "cbbd"},
			"bb",
		},

		{
			"bb",
			args{s: "bb"},
			"bb",
		},
		{
			"ccc",
			args{s: "ccc"},
			"ccc",
		},
		{
			"aaaa",
			args{s: "aaaa"},
			"aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
