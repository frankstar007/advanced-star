/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 * @desc lru缓存 id:146
 * requirement: 最久未使用的key 剔除 保证put get 是O(1)
 **/

package hash


type LinkedNode struct {
	key, value int
	pre, next *LinkedNode
}


type LRUCache struct {
     cap int
     size int
     cache map[int]*LinkedNode
     head, tail *LinkedNode
}


func initLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key: key,
		value: value,
	}
}


func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cap: capacity,
		cache: map[int]*LinkedNode{},
		head: initLinkedNode(0, 0),
		tail: initLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node

}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) removeTail() *LinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}



func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; ok {
		//如果存在 更新其值
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	} else {
		node := initLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size ++
		if this.size > this.cap {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size --
		}
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
