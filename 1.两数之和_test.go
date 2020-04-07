package cn

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func Test_twoSum(t *testing.T) {
	testTarget := []int{4,5,6}
	tests := [][]int {
		[]int {1,2,3,4,5},
		[]int {1,2,3,4,5},
		[]int {1,2,3,4,5},
	}
	for index := 0; index < len(tests); index++ {
		got := twoSum(tests[index], testTarget[index])
		fmt.Println("result=",got)
	}
}

func Test_twoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1",args{nums:[]int {1,2,3,4,5}, target:4}, nil},
		{"1",args{nums:[]int {1,2,3,4,5}, target:5}, nil},
		{"1",args{nums:[]int {1,2,3,4,5}, target:6}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}