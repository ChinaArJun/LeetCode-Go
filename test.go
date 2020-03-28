package main

import "fmt"

func rotate(nums []int, k int)  {
	//itemsOne := nums[:k]
	//itemsTwo := nums[k:]
	//nums = append(itemsTwo,itemsOne...)
	// itemsOne := nums[:k]
	// itemsTwo := nums[k:]
	l := len(nums)
	j := k % l
	// nums = append(itemsTwo,itemsOne...)
	nums = append(nums[l-j:l],nums[:l-j]...)
	fmt.Println(nums)
}

func main() {
	rotate([]int{1,2,3,4,5,6,7}, 2)
	fmt.Println()
}