/**
 *
 * Name: Complete Binary Tree
 * Description: A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.
 * Author: Ma Base
 * Date: 2022/12/24
 * Repository: https://github.com/BaseMax/CompleteTreeGo
 * License: GPL-3.0
 *
 **/

package main

/**
 * Node struct
 */
type Node struct {
	data  int
	left  *Node
	right *Node
}

/**
 * Tree struct
 */
type Tree struct {
	root *Node
}

/**
 * Create a new node
 */
func newNode(data int) *Node {
	return &Node{data: data}
}

/**
 * Check a tree is complete or not
 */
func (tree *Tree) isComplete() bool {
	if tree.root == nil {
		return true
	}
	queue := []*Node{tree.root}
	isLeaf := false

	// Level order traversal
	for len(queue) > 0 {
		// Get first node from queue
		node := queue[0]
		// Remove first node from queue
		queue = queue[1:]

		// If leaf node is found, then the tree is not complete
		if isLeaf && (node.left != nil || node.right != nil) {
			return false
		}

		// If left child is not null, add it to queue
		if node.left != nil {
			queue = append(queue, node.left)
			// If right child of current node is not null, then the tree is not complete
		} else if node.right != nil {
			return false
		}
		// If right child of current node is not null, add it to queue
		if node.right != nil {
			queue = append(queue, node.right)
			// Change isLeaf to true, so that the next node is leaf node
		} else {
			isLeaf = true
		}
	}

	// Otherwise, the tree is complete
	return true
}
