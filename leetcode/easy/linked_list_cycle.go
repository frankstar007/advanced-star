/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/1/27
 * @contact frankstarye@tencent.com
 * @desc linked list cycle id:141
 **/

package easy

//hasCycle 是否有环
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	p, q := head, head.Next

	for p != nil && q != nil {
		p = p.Next

		if q.Next == nil {
			return false
		}
		q = q.Next.Next
		if p == q {
			return true
		}
	}
	return false

}