package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Next.Val < head.Val {
			head.Val, head.Next.Val = head.Next.Val, head.Val
		}
		return head
	}
	// 快慢指针
	r, mid := head.Next, head
	for r.Next != nil {
		mid, r = mid.Next, r.Next
		if r.Next != nil {
			r = r.Next
		}
	}
	// 归并排序
	rLink := sortList(mid.Next)
	mid.Next = nil
	lLink := sortList(head)
	head = nil
	var tail *ListNode
	for rLink != nil || lLink != nil {
		if rLink == nil {
			tail.Next = lLink
			break
		}
		if lLink == nil {
			tail.Next = rLink
			break
		}
		// 取小数
		if lLink.Val < rLink.Val {
			if head == nil {
				head = &ListNode{
					Val: lLink.Val,
				}
				tail = head
			} else {
				tail.Next = &ListNode{
					Val: lLink.Val,
				}
				tail = tail.Next
			}
			lLink = lLink.Next
		} else {
			if head == nil {
				head = &ListNode{
					Val: rLink.Val,
				}
				tail = head
			} else {
				tail.Next = &ListNode{
					Val: rLink.Val,
				}
				tail = tail.Next
			}
			rLink = rLink.Next
		}
	}
	return head
}
