/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/9
 * @contact frankstarye@tencent.com

 * @desc id : 6 zigzag conversion
 **/

package medium

import (
	"strings"
)

//convert zigzag
func convert(s string, numRows int) string {

	if len(s) <= 0 {
		return s
	}

	flag := true
	items := make([]string, 0)
	//遍历元素
	for i := 0; i < len(s); {
		//如果是奇数列
		if flag {
			odd := strings.Builder{}
			times := 0
			for times < numRows {
				if i < len(s) {
					odd.WriteRune(rune(s[i]))
				} else {
					odd.WriteRune(' ')
				}
				times ++
				i++
			}
			flag = false

			items = append(items, odd.String())
		} else {
			//如果是偶数列
			//分几列
			for l := 1; l < numRows - 1; l ++ {
				even := strings.Builder{}
				for z := 0; z < numRows; z++ {
					if z == numRows - 1 - l  {
						if i < len(s) {
							even.WriteRune(rune(s[i]))
						} else {
							even.WriteRune(' ')
						}

						i ++
					} else {
						even.WriteRune(' ')
					}
				}

				items = append(items, even.String())
			}
			flag = true
		}
	}

	//最后遍历每个item
	result := strings.Builder{}
	start := 0
	for r := 0; r < numRows; r ++ {
		for _, item := range items {
			//因为上述操作保证个元素长度都是一样的 所以直接横向取元素拼接即可
			if rune(item[start]) == ' ' {
				continue
			}
			result.WriteRune(rune(item[start]))
		}
		start ++
	}

	return result.String()

}