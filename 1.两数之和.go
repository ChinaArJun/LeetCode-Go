package cn

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表

  
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	// 解题思路：
	// 时间复杂度是 O(n)
	// 由于哈希查找的时间复杂度为 O(1)，所以可以利用哈希容器 map 降低时间复杂度
	// 循环找出每一个元素遍历，将不存在的当前值存入map
	// 通过map查找key找出另外一个值的位置
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tag := target - nums[i]
		if k, ok := m[tag]; ok {
			return []int{k,i}
		}
		m[nums[i]] = i
	}
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
