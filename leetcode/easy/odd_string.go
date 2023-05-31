/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/5/25
 * @contact frankstarye@tencent.com
 **/

package easy

import "fmt"

func oddString(words []string) string {
	arrs := make([][]int, len(words))
	for i, word := range words {
		arrs[i] = computeArr(word)
	}
	fmt.Println(arrs)
	var j int
	for j = 1; j+1 < len(words); j++ {
		if !equalArr(arrs[j-1], arrs[j]) {
			if equalArr(arrs[j+1], arrs[j]) {
				return words[j-1]
			} else {
				return words[j]
			}

		}
	}
	return words[j]
}

func equalArr(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}

func computeArr(word string) []int {
	arr := make([]int, len(word)-1)
	for i := 0; i+1 < len(word); i++ {
		arr[i] = int(word[i+1] - word[i])
	}
	return arr
}
