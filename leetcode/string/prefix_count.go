/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/8
 * @contact frankstarye@tencent.com
 * @desc counting words with a given prefix
**/

package string

import "fmt"

func prefixCount(words []string, pref string) int {

	prefChar := []rune(pref)
	result := make([]string,0)
	for _, word := range words {
		//如果当前字符串 小于 前缀字符 跳过
		if len(word) < len(pref) {
			continue
		}
		wordChar := []rune(word)
		if wordChar[0] != prefChar[0] {
			continue
		}
		pl := len(prefChar)

		if string(wordChar[:pl]) == pref {
			result = append(result, word)
		}
	}

	fmt.Println(fmt.Sprintf("result:%v", result))

	return len(result)
}
