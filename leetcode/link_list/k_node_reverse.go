/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/12/8
 * @contact frankstarye@tencent.com
 * @desc k个有翻转链表
 **/

package link_list

import "fmt"

//reverseKGroup per k node reverse
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	lengthPoint := head
	//step 1 获取链表的长度
	for lengthPoint != nil {
		n++
		lengthPoint = lengthPoint.Next
	}
	if n < k {
		return head
	}
	group := n / k
	//计算需要翻转的组个数
	lists := make([]*ListNode, 0)
	start := head
	i := 0
	last := 0
	for start != nil && last < group{
		last ++
		tmp := start
		var mid *ListNode
		var end *ListNode
		for tmp != nil {
			if i == k - 1 {
				mid = tmp
			}
			if i == k {
				end = tmp
				break
			}
			tmp = tmp.Next
			i++
		}

		mid.Next = nil
		mid = nil
		lists = append(lists, start)
		start = end
		i = 0
	}
	left := start
	fmt.Printf("list:%v", lists)
	//step3 每个组内进行链表的翻转
	reverseList := make([]*ListNode, 0)
	for _, p := range lists {
		newP := reverse(p)
		reverseList = append(reverseList, newP)
	}
	newHead := &ListNode{Val: -1}
	t := newHead
	for _, rl := range reverseList {
		t.Next = rl
		for t.Next != nil {
			t = t.Next
		}
	}
	t.Next = left
	return newHead.Next

}

func reverse(head *ListNode) *ListNode{
	//递归出口
	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead

}