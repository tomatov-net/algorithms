package main

import (
	"fmt"
	binary_tree "tomatov.net/algorithms/binary-tree"
)

func main() {
	binaryTree := &binary_tree.Node{Data: 100}
	binaryTree.Insert(12)
	binaryTree.Insert(33)
	binaryTree.Insert(17)
	binaryTree.Insert(21)
	binaryTree.Insert(200)
	badTree := &binary_tree.Node{
		Data:  100,
		Left:  &binary_tree.Node{Data: 1000},
		Right: &binary_tree.Node{Data: 99},
	}
	isBalanced := binaryTree.IsBalanced(badTree)
	fmt.Println(isBalanced)
}
