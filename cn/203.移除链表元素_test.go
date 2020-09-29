package cn

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *cn.ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *cn.ListNode
		}{
			// TODO: Add test cases.
			//{name:1, args:ar, want:nil},
			//{"1", args{}, val:6}, 12},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}