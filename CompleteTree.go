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

import "fmt"

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

/**
 * Main function
 */
func main() {
	// Create a complete tree
	tree := Tree{}
	tree.root = newNode(1)
	tree.root.left = newNode(2)
	tree.root.right = newNode(3)
	tree.root.left.left = newNode(4)
	tree.root.left.right = newNode(5)
	tree.root.right.left = newNode(6)
	tree.root.right.right = newNode(7)
	tree.root.left.left.left = newNode(8)
	tree.root.left.left.right = newNode(9)
	tree.root.left.right.left = newNode(10)
	tree.root.left.right.right = newNode(11)
	tree.root.right.left.left = newNode(12)
	tree.root.right.left.right = newNode(13)
	tree.root.right.right.left = newNode(14)
	tree.root.right.right.right = newNode(15)

	if tree.isComplete() {
		fmt.Println("The tree is complete")
	} else {
		fmt.Println("The tree is not complete")
	}

	// Create a non-complete tree
	tree = Tree{}
	tree.root = newNode(1)
	tree.root.left = newNode(2)
	tree.root.right = newNode(3)
	tree.root.left.left = newNode(4)
	tree.root.left.right = newNode(5)
	tree.root.right.left = newNode(6)
	tree.root.right.right = newNode(7)
	tree.root.left.left.right = newNode(9)

	if tree.isComplete() {
		fmt.Println("The tree is complete")
	} else {
		fmt.Println("The tree is not complete")
	}
}
