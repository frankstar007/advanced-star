/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/15
 * @contact frankstarye@tencent.com
 * @desc group anagrams id: 49
 **/

package sort

func groupAnagrams(strs []string) [][]string {
	strIndex := make(map[string][]string)

	for _, str := range strs {
		//对每个字符串进行排序
		sortStr := selectSort(str)
		if _, ok := strIndex[sortStr]; ok {
			strIndex[sortStr] = append(strIndex[sortStr], str)
		} else {
			strIndex[sortStr] = []string{str}
		}
	}

	result := make([][]string, 0)
	for _, v := range strIndex {
		result = append(result, v)
	}

	return result
}


//选择排序
func selectSort(str string) string {
	chars := []rune(str)
	for i := 0; i < len(chars) - 1; i ++ {
		minIndex := i

		for j := i + 1; j < len(chars); j++ {
			if chars[j] < chars[minIndex] {
				minIndex = j
			}

		}
		if minIndex != i {
			swap(chars, minIndex, i)
		}

	}

	return string(chars)
}

func swap(str []rune, j int, i int) {
	str[j], str[i] = str[i], str[j]
}