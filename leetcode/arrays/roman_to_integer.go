/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/8
 * @contact frankstarye@tencent.com
 * @desc roman to integer
 **/

package arrays


func romanToInt(s string) int {
	symbol := make(map[rune]int)
	symbol['I'] = 1
	symbol['V'] = 5
	symbol['X'] = 10
	symbol['L'] = 50
	symbol['C'] = 100
	symbol['D'] = 500
	symbol['M'] = 1000
	//symbol["IV"] = 4
	//symbol["IX"] = 9
	//symbol["XL"] = 40
	//symbol["XC"] = 90
	//symbol["CD"] = 400
	//symbol["CM"] = 900

	num := 0
	chars := []rune(s)
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == 'V' {
			pre := i - 1
			if pre >= 0 {
				//IV
				if chars[pre] == 'I'{
					num += 4
					i--
					continue
				}
			}
		}
		if chars[i] == 'X' {
			pre := i - 1
			if pre >= 0 {
				//IX
				if chars[pre] == 'I'{
					num += 9
					i--
					continue
				}
			}
		}
		if chars[i] == 'L' {
			pre := i - 1
			if pre >= 0 {
				//XL
				if chars[pre] == 'X' {
					num += 40
					i--
					continue
				}
			}
		}
		if chars[i] == 'C' {
			pre := i - 1
			if pre >= 0 {
				//XC
				if chars[pre] == 'X' {
					num += 90
					i--
					continue
				}
			}
		}

		if chars[i] == 'D' {
			pre := i - 1
			if pre >= 0 {
				//CD
				if chars[pre] == 'C' {
					num += 400
					i--
					continue
				}
			}
		}
		if chars[i] == 'M' {
			pre := i - 1
			if pre >= 0 {
				//CM
				if chars[pre] == 'C' {
					num += 900
					i--
					continue
				}
			}
		}
		num += symbol[chars[i]]
	}


	return num
}

var symbol = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}



func standardRomanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbol[s[i]]
		if i < n - 1 && value < symbol[s[i + 1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}