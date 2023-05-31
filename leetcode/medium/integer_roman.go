/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/12
 * @contact frankstarye@tencent.com
 * @desc id : 12
 **/

package medium

import "strings"

func intToRoman(num int) string {
	roman := make(map[int]string)
	roman[1] = "I"
	roman[5] = "V"
	roman[10] = "X"
	roman[50] = "L"
	roman[100] = "C"
	roman[500] = "D"
	roman[1000] = "M"
	roman[4] = "IV"
	roman[9] = "IX"
	roman[40] = "XL"
	roman[90] = "XC"
	roman[400] = "CD"
	roman[900] = "CM"

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sb := strings.Builder{}

	for i := 0; i < len(keys) ; i ++ {
		m := num / keys[i]
		if m == 0 {
			continue
		}
		k := m
		if k >= 1 && k <= 5 {
			for k > 0 {
				sb.WriteString(roman[keys[i]])
				k--
			}
			num = num - keys[i] * m
		}
	}

	return sb.String()

}