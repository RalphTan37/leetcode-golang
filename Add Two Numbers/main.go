package main

import (
	"fmt"
	"strings"
)

// Defintion for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{} //dummy node, essentially a placeholder - start a new linked list
	curr := dummyHead        //current pointer that points to dummyhead, traverse thru linked list
	carry := 0               //store the carry for the sum of two nums, initialized to 0

	for l1 != nil || l2 != nil { //as long as there is a digit in the list(s)
		var v1, v2 int //represent values of current nodes of l1 & l2
		if l1 != nil {
			v1 = l1.Val  //assign value of v1 as current node of l1
			l1 = l1.Next //move to next node in l1
		}
		if l2 != nil {
			v2 = l2.Val  //assign value of v2 as current node of l2
			l2 = l2.Next //move to next node in l1
		}
		sum := v1 + v2 + carry               //sum of two values & carry
		carry = sum / 10                     //if the sum is > 10 then carry will be 1 - ex. 8 + 9 = 17 -> carry = 1
		curr.Next = &ListNode{Val: sum % 10} //new node added to linked list
		curr = curr.Next                     //moves the next node in the result list
	}

	if carry > 0 { //if there is a remaining carry
		curr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next //returns result
}

// declares a slice, appends each value in the list as a string and prints the linked list
func printLinkedList(l *ListNode) {
	var sum []string
	for l != nil {
		sum = append(sum, fmt.Sprintf("%d", l.Val))
		l = l.Next
	}
	fmt.Println((strings.Join(sum, "")))
}

// run tests
func main() {
	//Test Case #1
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	//l1 = [2,4,3]

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	//l2 = [5,6,4]

	result := addTwoNumbers(l1, l2) //342 + 465 = 807
	printLinkedList(result)         //[7,0,8] -> 708

	/*
		//Test Case #2
		l1 := &ListNode{Val: 0}
		//l1 = [0]

		l2 := &ListNode{Val: 0}
		//l2 = [0]

		result := addTwoNumbers(l2, l2) //0 + 0 = 0
		printLinkedList(result) //[0] -> 0
	*/

	/*
		//Test Case #3
		l1 := &ListNode{Val: 9}
		l1.Next := &ListNode{Val: 9}
		l1.Next.Next := &ListNode{Val: 9}
		l1.Next.Next.Next := &ListNode{Val: 9}
		l1.Next.Next.Next.Next := &ListNode{Val: 9}
		l1.Next.Next.Next.Next.Next := &ListNode{Val: 9}
		l1.Next.Next.Next.Next.Next.Next := &ListNode{Val: 9}
		//l1 = [9,9,9,9,9,9,9]

		l2 := &ListNode{Val: 9}
		l2.Next := &ListNode{Val: 9}
		l2.Next.Next := &ListNode{Val: 9}
		l2.Next.Next.Next := &ListNode{Val: 9}
		//l2 = [9,9,9,9]

		result := addTwoNumbers(l1, l2) //9999999 + 9999 = 89990001
		printLinkedList(result) //[1,0,0,0,9,9,9,8] -> 10009998
	*/
}
