/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 **/

package stack

import "testing"

func Test_Stack(t *testing.T) {
	sk := NewStack()

	sk.Push(1)
	sk.Push(2)
	sk.Pop()
	sk.Push(3)
	sk.Push(5)
	sk.Push(6)
	sk.Pop()


	sk.Print()

}
