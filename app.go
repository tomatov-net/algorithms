package main

import (
	"fmt"
	"math/rand"
	"time"
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
	isBalanced := binaryTree.IsBST(badTree)
	fmt.Println(isBalanced)

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Perm(10))
	for _, val := range rand.Perm(100000000) {
		binaryTree.Insert(val)
		fmt.Println(val)
	}

}
