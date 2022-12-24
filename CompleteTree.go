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
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if isLeaf && (node.left != nil || node.right != nil) {
			return false
		}
		if node.left != nil {
			queue = append(queue, node.left)
		} else if node.right != nil {
			return false
		}
		if node.right != nil {
			queue = append(queue, node.right)
		} else {
			isLeaf = true
		}
	}
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

}
