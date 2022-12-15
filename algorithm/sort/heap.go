/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
 * @desc 堆排序相关的处理
 * 堆的特点： a.完全二叉树  b.父节点大于子节点  c.由于是一个完全二叉树 所以可以使用一个一维数组来表示  d.任意一个节点i 都可以计算出它的
		父节点  和左右子节点  parent_node = (i - 1) / 2 , left_child = 2 * i + 1 , right_child = 2 (i + 1)
 **/

package sort

import "fmt"

//heapify 对 完全 二叉树 的某个节点进行 堆化 的操作
func heapify(tree []int, n, index int) {

	//递归出口
	if index >= n {
		return
	}
	//计算该节点的左右节点
	left := 2 * index + 1
	right := 2 * (index + 1)

	max := index
	//如果左孩子 大于根节点
	if left < n && tree[left] > tree[max] {
		max = left
	}
	//如果右孩子 大于根节点
	if right < n && tree[right] > tree[max] {
		max = right
	}

	//交换
	if max != index {
		swap(tree, max, index)
		//对于交换完后的节点继续进行heapify
		heapify(tree, n, max)
	}

	return

}

//swap 交换数组元素
func swap(tree []int, i, j int) {
	tree[i], tree[j] = tree[j], tree[i]
}

//buildHeap 构建堆
func buildHeap(tree []int, n int) {
	//获取最后一个需要进行堆化的节点
	last := (n - 2) / 2

	for i := last; i >= 0; i-- {
		//从下到上进行heapify
		heapify(tree, n, i)
	}
}

func heapSort(tree []int, n int) {
	//step1 build heap tree
	buildHeap(tree, n)
	last := n - 1
	for i := last; i >= 0 ; i--{
		swap(tree, i, 0)
		heapify(tree, i, 0)
	}

	fmt.Printf("sorted:%v", tree)
}
