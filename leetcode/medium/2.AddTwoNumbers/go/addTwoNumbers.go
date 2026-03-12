package addtwonumbers

import (
	"log"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func ListNodeToInt(l *ListNode) int {
	sum := 0
	power := 0
	curr := l
	for curr != nil {
		sum += curr.Val * (int(math.Pow10(power)))
		curr = curr.Next
		power++
	}
	return sum
}

// FAILS with big numbers
func AddTwoNumbersFails(l1 *ListNode, l2 *ListNode) *ListNode  {
	var head *ListNode = nil
	var curr *ListNode = nil
	if l1 == nil && l2 == nil {
		return nil
	}

	sum1 := ListNodeToInt(l1)
	log.Printf("sum1: %d", sum1)
	sum2 := ListNodeToInt(l2)
	log.Printf("sum2: %d", sum2)
	sum := sum1 + sum2
	log.Printf("Total sum: %d", sum)

	if sum < 10 {
		return &ListNode{
			Val: sum,
			Next: nil,
		}
	}

	for sum >= 10 {
		// get last number
		val := sum % 10
		// cut off last number
		sum = sum / 10

		tempNode := &ListNode{
			Val: val,
			Next: nil,
		}

		if head == nil {
			// first node
			head = tempNode
			// track current node
			curr = head
			continue
		}

		curr.Next = tempNode
		curr = curr.Next
	}

	lastNode := &ListNode{
		Val: sum,
		Next: nil,
	}
	
	curr.Next = lastNode
	return head
}

// attempt 2

func sumCarryHelper(sum *int, carry *int) {
	*sum += *carry
	// if sum excedes 10, then carry will carry the left over
	*carry = *sum / 10
	// sum will only get the last digit
	*sum = *sum % 10
}

func AddTwoNumbersSlow(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var curr *ListNode

	carry := 0
	for l1 != nil && l2 != nil {
		// inti sum 
		sum := l1.Val + l2.Val

		sumCarryHelper(&sum, &carry)

		tempNode := &ListNode{
			Val: sum,
		}
		if head == nil {
			head = tempNode
			curr = tempNode
		} else {
			curr.Next = tempNode
			curr = curr.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val
		sumCarryHelper(&sum, &carry)
		curr.Next = &ListNode{
			Val: sum,
		}
		curr = curr.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val
		sumCarryHelper(&sum, &carry)
		
		curr.Next = &ListNode{
			Val: sum,
		}
		curr = curr.Next
		l2 = l2.Next
	}

	if carry > 0 {
		tempNode := &ListNode{
			Val: carry,
		}
		curr.Next = tempNode
	}
	return head
}


func AddTwoNumbersFast(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var curr *ListNode
	carry := 0
	
	for {
		if (l1 != nil && l2 != nil) {
			sum := l1.Val + l2.Val
			sumCarryHelper(&sum, &carry)

			tempNode := &ListNode{
				Val: sum,
			}

			if head == nil {
				head = tempNode
				curr = head
			} else {
				curr.Next = tempNode
				curr = curr.Next
			}

			l1 = l1.Next
			l2 = l2.Next
		}

		if (l1 != nil && l2 == nil) {
			sum := l1.Val
			sumCarryHelper(&sum, &carry)

			tempNode := &ListNode{
				Val: sum,
			}
			curr.Next = tempNode
			curr = curr.Next

			l1 = l1.Next
		}

		if (l2 != nil && l1 == nil) {
			sum := l2.Val
			sumCarryHelper(&sum, &carry)

			tempNode := &ListNode{
				Val: sum,
			}
			curr.Next = tempNode
			curr = curr.Next

			l2 = l2.Next
		}

		if (l1 == nil && l2 == nil) {
			if carry > 0 {
				tempNode := &ListNode{
					Val: carry,
				}
				curr.Next = tempNode
			}
			break
		}
	}
	return head
}

// MAIN SOLUTION: fastest
func AddTwoNumbersFast2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: 0,
	}
	curr := head

	carry := 0
	
	for {
		if (l1 != nil && l2 != nil) {
			sum := l1.Val + l2.Val
			sumCarryHelper(&sum, &carry)

			curr.Next = &ListNode{
				Val: sum,
			}
			curr = curr.Next

			l1 = l1.Next
			l2 = l2.Next
		}

		if (l1 != nil && l2 == nil) {
			sum := l1.Val
			sumCarryHelper(&sum, &carry)

			curr.Next = &ListNode{
				Val: sum,
			}
			curr = curr.Next

			l1 = l1.Next
		}

		if (l2 != nil && l1 == nil) {
			sum := l2.Val
			sumCarryHelper(&sum, &carry)

			curr.Next = &ListNode{
				Val: sum,
			}
			curr = curr.Next

			l2 = l2.Next
		}

		if (l1 == nil && l2 == nil) {
			if carry > 0 {
				curr.Next = &ListNode{
					Val: carry,
				}
			}
			break
		}
	}
	return head.Next
}