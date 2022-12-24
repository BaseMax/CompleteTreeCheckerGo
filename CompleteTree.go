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

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

/**
 * Create a new node
 */
func newNode(data int) *Node {
	return &Node{data: data}
}

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
