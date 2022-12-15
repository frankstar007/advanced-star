/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
 * @desc 摇摆排序 id:324
 **/

package heap

import "fmt"

func wiggleSort(nums []int)  {
	length := len(nums)
	//堆排序
	heap_tree_sort(nums, length)
	fmt.Printf("sorted:%v", nums)
	fmt.Println()
	arr := append([]int{}, nums...)

	x := (length + 1) / 2
	for i, j, k := 0, x-1, length-1; i < length; i += 2 {
		nums[i] = arr[j]
		if i+1 < length {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}

}

//heap_node 单个节点 从小到大排序
func heap_node(nums []int, n , node int) {
	if node >= n {
		return
	}
	max := node
	left := 2 * node + 1
	right := 2 * node + 2
	if left < n && nums[left] > nums[max] {
		max = left
	}
	if right < n && nums[right] > nums[max] {
		max = right
	}
	if max != node {
		swap(nums, max, node)
		heap_node(nums, n, max)
	}
}

//heap_tree 构造堆 二叉完全树
func heap_tree(nums []int, n int) {
	last := n - 1
	parent := (last - 1) / 2
	for i := parent; i >= 0; i-- {
		heap_node(nums, n, i)
	}
}

//heap_tree_sort 二叉完全树 排序
func heap_tree_sort(nums []int, n int) {
	heap_tree(nums, n)
	last := n - 1
	for i := last; i >= 0; i-- {
		swap(nums, i, 0)
		heapify(nums, i, 0)
	}
}