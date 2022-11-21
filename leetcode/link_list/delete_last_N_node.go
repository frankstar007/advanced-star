/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2022/11/20
 * @contact frankstarye@tencent.com
 * @desc 删除倒数第N个节点 id:19
 **/

package link_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 先获取链表的长度
	length := 0
	p := head
	for p != nil {
		length += 1
		p = p.Next
	}
	//如果删除头结点
	if n == length {
		return head.Next
	}

	// 如果删除中间节点
	q := head
	for i := 1; i < (length - n); i++ {
		q = q.Next
	}
	q.Next = q.Next.Next
	return head
}
