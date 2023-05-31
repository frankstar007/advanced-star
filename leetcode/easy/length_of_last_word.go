/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 * id: 58
 **/

package easy

func lengthOfLastWord(s string) int {

	i, end := len(s) - 1, len(s) - 1
	//从末尾找到不为空的位置
	for  i >= 0 && s[i] - ' ' == 0 {
		i --
	}
	end = i
	result := 0
	for end >= 0 && s[end] - ' ' != 0 {
		end --
		result ++
	}

	return result


}