/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/19
 * @contact frankstarye@tencent.com
 * @Desc remove duplicates form sorted list
**/

package easy


func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil {
		val := p.Val
		k := p.Next
		if k == nil {
			break
		}
		for k != nil && k.Val == val {
			k = k.Next
		}
		p.Next = k
		p = p.Next
	}
	return head
}