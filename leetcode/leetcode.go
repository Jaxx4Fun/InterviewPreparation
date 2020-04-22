package leet

import "container/list"

// 1248. 统计「优美子数组」
// 给你一个整数数组 nums 和一个整数 k。

// 如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。

// 请返回这个数组中「优美子数组」的数目。
func numberOfSubarrays(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	count := 0
	oddPos := make([]int, length+2, length+2)

	for i, v := range nums {
		if v&1 == 1 {
			count++
			oddPos[count] = i
		}
	}

	ans := 0
	oddPos[0] = -1
	count++
	oddPos[count] = len(nums)
	for i := 1; i+k <= count; i++ {
		ans += (oddPos[i] - oddPos[i-1]) * (oddPos[i+k] - oddPos[i+k-1])
	}
	return ans

}

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 199. 二叉树的右视图
// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// 思路：BFS——先放入右节点、再放入左节点
//		每一层次的第一个节点，必定是该层次的最右节点，放入结果中
// 时间复杂度：遍历了所有节点，所以是O(N)
// 空间复杂度：用链表作为队列，队列的最大容量应该是在满二叉树的情况下，最底层的节点数O(2**(log(N)-1))
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			front := queue.Front()
			queue.Remove(front)
			node := front.Value.(*TreeNode)
			if i == 0 {
				ans = append(ans, node.Val)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
	}
	return
}
