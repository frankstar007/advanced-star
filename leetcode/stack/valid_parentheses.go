/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/9
 * @contact frankstarye@tencent.com
 * @desc 有效的括号 id:20
**/

package stack

import "fmt"

func isValid(s string) bool {

	stack := make([]byte, 0)

	for i := 0; i < len(s); i ++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) <= 0 {
				return false
			}
			fmt.Println(fmt.Sprintf("%v", stack))
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if s[i] == ')' && top != '(' {
				return false
			} else if s[i] == '}' && top != '{' {
				return false
			} else if s[i] == '[' && top != '[' {
				return false
			}
		}
	}
	return true


}