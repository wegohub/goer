package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	head := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	cur := sortList(head)
	PrintLink(cur)
	return "Hello World!", nil
}

func sortList(head *ListNode) *ListNode {
	N := 0
	cur := head
	for cur != nil {
		N++
		cur = cur.Next
	}
	// 抓住返回头
	h := head
	// 每一组的头节点
	teamFirst := head
	// 前驱指针
	var pre *ListNode

	for step := 1; step < N; step <<= 1 {
		for teamFirst != nil {
			team := TeamInfo(teamFirst, step)
			teamHead, teamTail := merge(team[0], team[1], team[2], team[3])
			if h == teamHead || pre == nil {
				h = teamHead
				pre = teamTail
			} else {
				pre.Next = teamHead
				pre = teamTail
			}
			// 来到下一组的开始位置
			teamFirst = team[4]
		}
		// 为下一组左准备
		teamFirst = h
		pre = nil
	}
	return h
}

// 可以用快慢指针
// 传入分组的开始指针和步长，返回左边的头和尾，右边的头和尾，以下一组的开始节点(teamFirst)
func TeamInfo(teamFist *ListNode, step int) [5]*ListNode {
	var (
		// 左组的头
		ls = teamFist
		// 左组的尾
		le = teamFist
		// 右组的头
		rs *ListNode
		// 右组的尾
		re *ListNode
		// 下一组的开始节点
		next *ListNode
	)

	pass := 0
	for teamFist != nil {
		pass++
		if pass <= step {
			le = teamFist
		}
		if pass == step+1 {
			rs = teamFist
		}
		if pass > step {
			re = teamFist
		}
		// 终止
		if pass == (step << 1) {
			break
		}
		teamFist = teamFist.Next
	}

	// 断开左边
	le.Next = nil

	// 设置下一组的开始
	if re != nil {
		next = re.Next
		// 断开右组
		re.Next = nil
	}
	return [5]*ListNode{ls, le, rs, re, next}
}

// 合并左右两个组，入参 左右组的头和尾
func merge(ls, le, rs, re *ListNode) (*ListNode, *ListNode) {
	if rs == nil {
		return ls, le
	}
	var (
		head *ListNode
		pre  *ListNode
		cur  *ListNode
		tail *ListNode
	)
	for ls != le.Next && rs != re.Next {
		if ls.Val <= rs.Val {
			cur = ls
			ls = ls.Next
		} else {
			cur = rs
			rs = rs.Next
		}
		if pre == nil {
			head = cur
			pre = cur
		} else {
			pre.Next = cur
			pre = cur
		}
	}

	// 到这里无论如何都会剩下节点，所以在这里设置tail指针
	if ls != le.Next {
		for ls != le.Next {
			pre.Next = ls
			pre = ls
			tail = ls
			ls = ls.Next
		}
	} else {
		for rs != re.Next {
			pre.Next = rs
			pre = rs
			tail = rs
			rs = rs.Next
		}

	}

	return head, tail
}
