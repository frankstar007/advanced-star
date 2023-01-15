/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/8
 * @contact frankstarye@tencent.com
 * @desc Palindrome Number id : 9
 **/

package arrays

//isPalindrome 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	m := x
	arr := make([]int, 0)
	for m > 0 {
		reminder := m % 10
		arr = append(arr, reminder)
		m = m / 10
	}

	num := 0
	for _, ar := range arr {
		num = num * 10 + ar
	}
	if num == x {
		return true
	}

	return false

}