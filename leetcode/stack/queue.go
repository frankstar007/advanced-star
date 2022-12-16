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
	"sync"
)

type Queue struct {
	list *list.List
	lock *sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
		lock: &sync.RWMutex{},
	}
}


func (q *Queue) Push(val interface{}) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.list.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) Peek() interface{} {
	e := q.list.Front()
	if e != nil {
		return e.Value
	}
	return nil
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}