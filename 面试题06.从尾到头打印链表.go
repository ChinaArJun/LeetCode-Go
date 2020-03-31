package cn

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
// 
//
// 示例 1： 
//
// 输入：head = [1,3,2]
//输出：[2,3,1] 
//
// 
//
// 限制： 
//
// 0 <= 链表长度 <= 10000 
//

  
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*type ListNode struct {
	Val int
	Next *ListNode
}*/
func reversePrint(head *ListNode) []int {
	// 解题思路：取值，反转
	// 老是报错 ： panic: runtime error: invalid memory address or nil pointer dereference
	//			[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x4bb046]
	//
	if head.Next == nil {
		return nil
	}
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i ++ {
		res[i] = arr[len(arr)- 1 - i]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
