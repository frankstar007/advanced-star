/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/15
 * @contact frankstarye@tencent.com
 **/

package collection

import "sync"

type Queue struct {
	element []Element
	lock    *sync.RWMutex
}



func NewQueue() *Queue {
	return &Queue{
		element: []Element{},
		lock:    &sync.RWMutex{},
	}
}

func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.element = append(q.element, item)
}

func (q *Queue) Pop() Element {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.Empty() {
		return nil
	}
	e := q.element[0]
	q.element = q.element[1:]
	return e
}


func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	return len(q.element)
}

