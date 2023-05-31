/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/20
 * @contact frankstarye@tencent.com
 **/

package medium

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	k := 1
	str := "1"
	for k < n {
		strs := count(str)
		str = say(strs)
		k ++
	}

	return str
}

func say(strs []string) string {
	sb := strings.Builder{}
	for _, s := range strs {
		sb.WriteString(fmt.Sprintf("%d%d", len(s), s[0] - '0'))
	}
	return sb.String()

}

func count(str string) []string {

	items := make([]string, 0)
	for i := 0; i < len(str); {
		j := i + 1
		for j < len(str) && str[j] == str[i]{
			j ++
		}
		//
		if j < len(str) {
			items = append(items, str[i : j])
			i = j
		} else {
			items = append(items, str[i:])
			break
		}


	}
	fmt.Println(items)
	return items
}

