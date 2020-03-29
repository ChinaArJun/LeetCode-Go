package cn
//删除链表中等于给定值 val 的所有节点。 
//
// 示例: 
//
// 输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5
// 
// Related Topics 链表

  
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
func removeElements(head *ListNode, val int) *ListNode {
	// 思路： 循环查找下一个节点，当查找到给定值，将下一个节点指向后一个节点，
	// 当节点最后为空时，所有节点查找完毕
	if head == nil {
		return head
	}
	headNode := &ListNode{Next:head, Val:0}
	curNode := head
	preNode := headNode
	for curNode != nil   {
		if curNode.Val == val {
			// 移除指定节点
			preNode.Next = curNode.Next
		} else {
			preNode = curNode
		}
		curNode = curNode.Next
	}
	return  headNode.Next
}
//leetcode submit region end(Prohibit modification and deletion)
