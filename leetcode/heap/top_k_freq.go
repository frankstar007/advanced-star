/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/4
 * @contact frankstarye@tencent.com
`* @desc top k frequency nums id:347 前top k个高频元素
 **/

package heap

import "fmt"



//topKFrequent 前top k 个高频元素
func topKFrequent(nums []int, k int) []int {
	// 存放 元素->出现次数
	exist := make(map[int]int)
	for _, num := range nums {
		if _, ok := exist[num]; ok {
			exist[num] += 1
		} else {
			exist[num] = 1
		}
	}
	// 存放 出现次数 -> 元素集合
	frequency := make([]int, 0)
	freNum := make(map[int][]int)
	for num, fre := range exist {
		if _, ok := freNum[fre]; ok {
			freNum[fre] = append(freNum[fre], num)
		} else {
			freNum[fre] = []int{num}
			frequency = append(frequency, fre)
		}
	}
	//堆排序 频次数组
	heap_sort(frequency, len(frequency))
	//获取前top K 的频次
	result := make([]int, 0)
	for i := len(frequency)-1; i >= 0 ; i-- {
		if k <= 0 {
			break
		}
		size := len(freNum[frequency[i]])
		if k >= size {
			result = append(result, freNum[frequency[i]]...)
			k = k - size
		} else {
			part := freNum[frequency[i]][:size - k]
			result = append(result, part...)
			k = size - k
		}

	}
	fmt.Printf("top k frequency elements:%v", result)
	return result
}

func swap(frequency []int, i int, j int) {
	frequency[i], frequency[j] = frequency[j], frequency[i]
}

//heap_sort 堆排序
func heap_sort(nums []int, n int) {
	build_heap(nums, n)
	last := n - 1
	for i := last; i >= 0; i-- {
		swap(nums, i, 0)
		heapify(nums, i, 0)
	}
}

//build_heap 构建堆
func build_heap(nums []int, n int) {
	last := n - 1
	parent := (last - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(nums, n, i)
	}
}


func heapify(nums []int, n, index int) {
	if n <= index {
		return
	}
	//计算当前节点的左右节点
	left := 2 * index + 1
	right := 2 * index + 2
	max := index
	if left < n && nums[left] > nums[max] {
		max = left
	}
	if right < n && nums[right] > nums[max] {
		max = right
	}
	if max != index {
		swap(nums, max, index)
		heapify(nums, n, max)
	}

}
