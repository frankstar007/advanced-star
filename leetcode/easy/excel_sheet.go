/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/2
 * @contact frankstarye@tencent.com
 * @desc id :=171
 **/

package easy

func titleToNumber(columnTitle string) int {


	sum := 0
	j := 0
	for  i := len(columnTitle) - 1; i >= 0; i-- {
		k := j
		multiply := 1
		for k > 0 {
			multiply *= 26
			k--
		}
		val := rune(columnTitle[i])
		l := int(val - 'A')
		sum += (l  + 1) * multiply
		j ++
	}

	return sum

}