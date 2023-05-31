/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/24
 * @contact frankstarye@tencent.com
 * @desc convert sorted array to binary search tree id : 108
 **/

package easy

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums)

}

func buildBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val:nums[mid]}
	if mid - 1 >= 0 {
		root.Left = buildBST(nums[0:mid])
	}
	if mid + 1 < len(nums) {
		root.Right = buildBST(nums[mid + 1:])
	}
	return root
}