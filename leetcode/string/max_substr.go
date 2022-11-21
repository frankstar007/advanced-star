/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 * @desc 无重复字符的最长子串 id:3
 **/

package string

func lengthOfLongestSubstring(s string) int {
	var i, j int
	max := 0
	for i = 0 ; i < len(s); {
		exist := make(map[rune]bool)
		exist[rune(s[i])] = true
		for j = i + 1; j < len(s); j++ {
			if _, ok := exist[rune(s[j])]; ok {
				gap := j - i
				if gap > max {
					max = gap
				}
				i++
				break
			} else  {
				exist[rune(s[j])] = true
			}
		}
		if j >= len(s)  {
			gap := j - i
			if gap > max {
				max = gap
			}
			break
		}

	}

	return max
}
