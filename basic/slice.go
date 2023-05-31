/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/4/7
 * @contact frankstarye@tencent.com
 **/

package main

import (
	"fmt"
	"reflect"
)

func Deduplicate(A []interface{}) []interface{} {
	exist := make(map[string]bool)
	var unique []interface{}
	for _, ele := range A {
		eleVal := reflect.ValueOf(ele)
		if _, ok := exist[eleVal.String()]; !ok {
			exist[eleVal.String()] = true
			unique = append(unique, ele)
		}
	}
	return unique
}

func main() {

	a := []int{2, 3}
	fmt.Println(a[:2])
}
