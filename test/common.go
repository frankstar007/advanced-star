/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/3/12
 * @contact frankstarye@tencent.com
 **/

package main

import "fmt"

func main() {
	arr := make([]string, 0)
	arr = append(arr, "1","2","3","4")
	fmt.Println(arr[:3])
}
