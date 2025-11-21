package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func main() {
//	node := ListNode{Val: 9}
//	node.Next = &ListNode{Val: 9}
//	node.Next.Next = &ListNode{Val: 9}
//	node2 := ListNode{Val: 9}
//	node2.Next = &ListNode{Val: 9}
//	node2.Next.Next = &ListNode{Val: 9}
//	node2.Next.Next.Next = &ListNode{Val: 9}
//	numbers := addTwoNumbers(&node, &node2)
//	for numbers != nil {
//		print(numbers.Val)
//		numbers = numbers.Next
//	}
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	numbers1 := addTwoNumbers1(l1, l2, 0)
	if numbers1 == 1 {
		return l1
	}
	return l2

}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode, ten int) int {

	if l1 != nil && l2 != nil {
		v := l1.Val + l2.Val + ten
		l1.Val = v % 10
		l2.Val = v % 10
		if l1.Next == nil && l2.Next == nil {
			if v >= 10 {
				l1.Next = &ListNode{Val: v / 10}
			}
			return 1
		}
		return addTwoNumbers1(l1.Next, l2.Next, v/10)
	} else if l1 != nil {
		v := l1.Val + ten
		l1.Val = v % 10
		if l1.Next == nil {
			if v >= 10 {
				l1.Next = &ListNode{Val: v / 10}
			}
			return 1
		} else {
			return addTwoNumbers1(l1.Next, nil, v/10)
		}

	} else if l2 != nil {
		v := l2.Val + ten
		l2.Val = v % 10
		if l2.Next == nil {
			if v >= 10 {
				l2.Next = &ListNode{Val: v / 10}
			}
			return 2
		} else {
			return addTwoNumbers1(nil, l2.Next, v/10)
		}
	}
	return 1
}
