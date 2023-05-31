/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/23
 * @contact frankstarye@tencent.com
 **/

package medium

import "sort"

func largestValFromLabels(values []int, labels []int, numWanted int, useLimit int) int {

	var sum int

	//转换一下value数组 以及对应的labels
	indexArr := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		indexArr[i] = i
	}
	sort.Slice(indexArr, func(i, j int) bool {
		return values[indexArr[i]] > values[indexArr[j]]
	})

	labelCnt := make(map[int]int) //存放每个label 使用了多少次
	for i := 0; i < len(values); i++ {

		if numWanted <= 0 {
			break
		}
		cnt := labelCnt[labels[indexArr[i]]]
		if cnt >= useLimit {
			continue
		}
		sum += values[indexArr[i]]
		labelCnt[labels[indexArr[i]]] = cnt + 1
		numWanted--
	}
	return sum

}
