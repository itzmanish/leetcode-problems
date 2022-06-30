// 206. Reverse Linked List
// Easy

// 12723

// 218

// Add to List

// Share
// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Example 1:

// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
// Example 2:

// Input: head = [1,2]
// Output: [2,1]
// Example 3:

// Input: head = []
// Output: []

// Constraints:

// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000

// Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

package leetcodepractice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head.Next
	head.Next = nil
	if next == nil {
		return head
	}
	newHead := reverseListRecursive(next)
	next.Next = head
	return newHead
}
