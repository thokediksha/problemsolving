package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Traverse the first linked list and store the memory addresses of all nodes in a map
	nodeMap := make(map[*ListNode]bool)
	for node := headA; node != nil; node = node.Next {
		nodeMap[node] = true
	}

	// Traverse the second linked list and check if each node's memory address is in the map
	for node := headB; node != nil; node = node.Next {
		if nodeMap[node] {
			return node
		}
	}

	// If we reach here, there is no common element
	return nil
}

func main() {
	// Initialize the two linked lists
	list1 := &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}}
	list2 := &ListNode{11, &ListNode{12, &ListNode{13, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}}}

	// Find the point of common element
	commonNode := getIntersectionNode(list1, list2)

	if commonNode != nil {
		fmt.Printf("The point of common element is %v\n", commonNode.Val)
	} else {
		fmt.Println("The two linked lists do not intersect")
	}

}





// Traverse the first linked list and store the memory addresses of all the nodes i
// n a map or hash table. This will take O(n1) time and O(n1) space, where n1 is the length 
// of the first linked list.

// Traverse the second linked list and for each node, check if its memory address is already present
// in the map or hash table. If it is, then we have found the point of common element. If not, continue
//   traversing the second linked list. This will take O(n2) time and O(1) space, where n2 is the length 
//   of the second linked list.

// If we reach the end of the second linked list without finding a common element,
//  then the two linked lists do not intersect.


// The time complexity of this algorithm is O(n1 + n2), where n1 and n2 are the lengths of the 
// two linked lists respectively. The space complexity is O(n1), because we need to store the
//  memory addresses of all nodes in the first linked list.