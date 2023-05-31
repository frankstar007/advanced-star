/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/4/5
 * @contact frankstarye@tencent.com
 **/

package main

import "fmt"

type Object struct {
	A int
	B int
	C int
}

func main() {
	os := make([]Object, 0)
	ob1 := Object{
		C: 1,
	}
	os = append(os, ob1)

	ob2 := Object{
		C: 2,
	}
	os = append(os, ob2)

	ob3 := Object{
		C: 3,
	}
	os = append(os, ob3)

	ob4 := Object{
		C: 4,
	}
	os = append(os, ob4)

	build(100, 1000, &os)

	for _, o := range os {
		fmt.Println(fmt.Sprintf("Object:%v", o))
	}
}

func build(a, b int, objects *[]Object) {

	for index := range *objects {
		o := (*objects)[index]

		o.A = a
		o.B = b

		(*objects)[index] = o
	}

}
