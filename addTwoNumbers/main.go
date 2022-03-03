package addTwoNumbers

import "fmt"

type Node struct {
	data int
	node *Node
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1temp := l1
	l2temp := l2
	var LenL1, LenL2 int
	LenL1 = 0
	LenL2 = 0

	for {
		LenL1 = LenL1 + 1
		if l1temp.Next == nil {
			break
		} else {
			l1temp = l1temp.Next
		}
	}

	for {
		LenL2 = LenL2 + 1
		if l2temp.Next == nil {
			break
		}
		l2temp = l2temp.Next
	}
	var MaxLen int
	if LenL2 >= LenL1 {
		MaxLen = LenL2
	} else {
		MaxLen = LenL1
	}

	var ListNodeSlice []*ListNode
	ListNodeSlice = make([]*ListNode, 0, MaxLen+1)
	for i := 0; i <= MaxLen; i++ {
		ListNodeSlice = append(ListNodeSlice, new(ListNode))
	}
	HeadNode := ListNodeSlice[0]
	CurrentNode := HeadNode
	tmp := 0
	for i := 0; i <= MaxLen; i++ {

		CurrentNode.Val = (tmp + l1.Val + l2.Val) % 10

		tmp = (tmp + l1.Val + l2.Val) / 10
		if i+1 >= len(ListNodeSlice) {
			CurrentNode.Next = nil
		} else {
			CurrentNode.Next = ListNodeSlice[i+1]
		}

		if l1.Next == nil {
			l1.Next = &ListNode{
				0, nil,
			}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{
				0, nil,
			}
		}

		l1 = l1.Next

		l2 = l2.Next
		CurrentNode = CurrentNode.Next
	}
	var PreNode *ListNode
	var tempHead *ListNode
	tempHead = HeadNode
	for {

		if HeadNode.Next == nil {
			if HeadNode.Val == 0 {
				PreNode.Next = nil
			}
			break
		} else {
			PreNode = HeadNode
			HeadNode = HeadNode.Next
		}
	}
	return tempHead
}

func testmain() {

	list_a := []int{9, 9}
	list_b := []int{9}

	l1 := makeNode(list_a)
	l2 := makeNode(list_b)

	ccc := addTwoNumbers(l1, l2)
	for {
		if ccc.Next == nil {
			fmt.Println(ccc.Val)
			break
		} else {
			fmt.Println(ccc.Val)
			ccc = ccc.Next
		}
	}
	//list_b := []int{4, 5, 6}

	//输入：l1 = [9,9,9,9,9,9,9],
	//           l2 = [9,9,9,9]

	//输出：[8,9,9,9,0,0,0,1]
	//
	//var SecendOne = new(Node)
	//SecendOne.node = LastOne
	//SecendOne.data = 4
	//var firstOne = new(Node)
	//firstOne.data = 2
	//firstOne.node = SecendOne
}

func makeNode(list_a []int) *ListNode {

	node_list := make([]*ListNode, 0, len(list_a))
	for _ = range list_a {
		node_list = append(node_list, new(ListNode))
	}
	var preNode = new(ListNode)
	for i := len(list_a) - 1; i >= 0; i = i - 1 {
		last_data := list_a[i]

		node_list[i].Val = last_data
		if i == len(list_a)-1 {
			node_list[i].Next = nil
		} else {
			node_list[i].Next = preNode
		}
		preNode = node_list[i]
	}

	headNode := node_list[0]

	return headNode

}
