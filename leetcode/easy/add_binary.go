/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 * @desc add binary id : 67
**/

package easy

func addBinary(a string, b string) string {
	sb := make([]rune, 0)
	flag := false
	arune, brune := []rune(a), []rune(b)
	i, j := len(arune) - 1, len(brune) - 1
	for i >= 0 && j >= 0 {
		if arune[i] != brune[j] {
			if flag {
				sb = append(sb, '0')
				flag = true
			} else {
				sb = append(sb, '1')
				flag = false
			}
		} else {
			if flag {
				sb = append(sb, '1')
			} else {
				sb = append(sb, '0')
			}
			if arune[i] == '0' {
				flag = false
			} else {
				flag = true
			}
		}
		i--
		j--
	}

	for i >= 0 {
		if !flag {
			sb = append(sb, arune[i])
		} else {
			if arune[i] == '1' {
				sb = append(sb, '0')
				flag = true
			} else {
				sb = append(sb, '1')
				flag = false
			}
		}

		i--
	}

	for j >= 0 {
		if !flag {
			sb = append(sb, brune[j])
		} else {
			if brune[j] == '1' {
				sb = append(sb, '0')
				flag = true
			} else {
				sb = append(sb, '1')
				flag = false
			}
		}

		j--
	}

	if flag {
		sb = append(sb, '1')
	}
	for k, m := 0, len(sb) - 1; k < m; k ,m = k + 1, m-1 {
		sb[k], sb[m] = sb[m], sb[k]
	}

	return string(sb)
}
