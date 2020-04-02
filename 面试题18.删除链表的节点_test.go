package cn

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	list := ListNode{Val:-1, Next: &ListNode{Val:-3, Next: &ListNode{Val: -2, Next: &ListNode{Val:4, Next:nil}}}}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.

		{"",args{head:&list}, nil},
		{"",args{head:&list}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, -1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}