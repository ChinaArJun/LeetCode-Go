package cn

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *cn.ListNode
	}
	list := cn.ListNode{Val: 1, Next: &cn.ListNode{Val: 2, Next: &cn.ListNode{Val: 3, Next: &cn.ListNode{Val: 4, Next:nil}}}}
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
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}