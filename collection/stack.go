/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 * @desc 栈的实现
 **/

package collection

import "sync"

type Element interface{}

type Stack struct {
	element []Element
	lock    *sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		element: []Element{},
		lock:    &sync.RWMutex{},
	}
}

func (s *Stack) Elements() []Element {
	result := make([]Element, 0)
	for _, e := range s.element {
		result = append(result, e)
	}
	return result
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.element = append(s.element, item)
}

func (s *Stack) Pop() Element {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Empty() {
		return nil
	}
	e := s.element[s.Len() - 1]
	s.element = s.element[:s.Len() - 1]
	return e
}

func (s *Stack) Top() Element {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Empty() {
		return nil
	}
	e := s.element[s.Len() - 1]
	return e
}


func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Len() int {
	return len(s.element)
}
