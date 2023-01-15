/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/11
 * @contact frankstarye@tencent.com
 * @desc 2283. Check if Number Has Equal Digit Count and Digit Value
**/

package arrays

func digitCount(num string) bool {
	numbers := []rune(num)
	if len(numbers) <= 0 {
		return true
	}
	vc := make(map[rune]int)
	//count every number
	for _, v := range num {
		vc[v - '0']++
	}
	//check index digit count equal vc
	for i, v := range num {
		if vc[rune(i)] != int(v - '0') {
			return false
		}
	}
	return true




}
