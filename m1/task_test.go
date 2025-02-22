package m1

import (
	"reflect"
	"testing"
)

func Test_revert(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		{
			"1",
			args{head: &Node{v: 1, next: &Node{v: 2, next: &Node{v: 3}}}},
			Node{v: 3, next: &Node{v: 2, next: &Node{v: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revert(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revert() = %v, want %v", got, tt.want)
			}
		})
	}
}
