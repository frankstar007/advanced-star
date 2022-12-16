/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 **/

package stack

import (
	"container/list"
	"fmt"
	"sync"
)

type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func NewStack() *Stack {
	data := list.New()
	lock := &sync.RWMutex{}
	return &Stack{
		list: data,
		lock: lock,
	}
}


func (s *Stack) Push(val interface{}) {
	defer s.lock.Unlock()
	s.lock.Lock()
	s.list.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	defer s.lock.Unlock()
	s.lock.Lock()
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Print() {
	for e := s.list.Front(); e != nil ; e = e.Next() {
		fmt.Print(e.Value, "")
	}
	fmt.Println()
}