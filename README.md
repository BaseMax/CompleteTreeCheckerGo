# Complete Tree Checker (Go)

This is a simple Go program that checks if a given **tree** is complete or not. A tree is complete if all levels are completely filled except possibly the last level and the last level has all keys as left as possible.

## Examples

In the following example, the tree is complete:

```go
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
```

And another non-complete tree example:

```go
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
```

Copyright (c) 2022, Max Base
