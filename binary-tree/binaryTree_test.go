package binary_tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	node := Node{Data: 666}
	node.Insert(12)

	if node.Left.Data != 12 {
		t.Fail()
	}

	node.Insert(777)

	if node.Right.Data != 777 {
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	node := Node{Data: 666}
	node.Insert(12)
	node.Insert(33)
	node.Insert(17)
	node.Insert(21)
	node.Insert(200)

	if node.Search(12) == nil {
		t.Fail()
	}

	if node.Search(121) != nil {
		t.Fail()
	}
}

func TestIsBST(t *testing.T) {
	normalTree := &Node{Data: 666}
	normalTree.Insert(12)
	normalTree.Insert(33)
	normalTree.Insert(17)
	normalTree.Insert(21)
	normalTree.Insert(200)

	badTree := &Node{
		Data:  100,
		Left:  &Node{Data: 1000}, //left is bigger than right
		Right: &Node{Data: 99},   //right is smaller than left and parent
	}

	if normalTree.IsBST(normalTree) != true {
		t.Fail()
	}

	if badTree.IsBST(badTree) == true {
		t.Fail()
	}
}

func TestHeight(t *testing.T) {
	normalTree := &Node{Data: 60} //root level 0
	normalTree.Insert(121)        //right level 1
	normalTree.Insert(33)         //left level 1
	normalTree.Insert(6)          //left level 2
	//                   60
	//                  /  \
	//				   33   121
	//				  /
	//				 6
	// Max(1,1) + Max(2, 0) = 3
	height := normalTree.Height(normalTree)

	if height != 3 {
		t.Fail()
	}
}
