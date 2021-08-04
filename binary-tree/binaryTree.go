package binary_tree

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func (n *Node) Insert(number int) *Node {
	if n.Data < number {
		//add to the right
		if n.Right == nil {
			n.Right = &Node{Data: number}
		} else {
			n.Right.Insert(number)
		}
	} else if n.Data > number {
		//add to the left
		if n.Left == nil {
			n.Left = &Node{Data: number}
		} else {
			n.Left.Insert(number)
		}
	}
	return n
}

func (n *Node) Search(number int) *Node {
	if n.Data < number {
		if n.Right != nil {
			return n.Right.Search(number)
		}
	} else if n.Data > number {
		if n.Left != nil {
			return n.Left.Search(number)
		}
	} else if n.Data == number {
		return n
	}
	return nil
}

func (n *Node) IsBalanced(node *Node) bool {

	if node == nil {
		return true
	}

	rightIsOk := true
	leftIsOk := true

	if node.Right != nil {
		rightIsOk = node.Data < node.Right.Data && node.Right.IsBalanced(node.Right)
	}

	if node.Left != nil {
		leftIsOk = node.Data > node.Left.Data && node.Left.IsBalanced(node.Left)
	}

	return rightIsOk && leftIsOk
}

func (n *Node) IsBST(node *Node) bool {

	if node == nil {
		return true
	}

	rightIsOk := true
	leftIsOk := true

	if node.Right != nil {
		rightIsOk = node.Data < node.Right.Data && node.Right.IsBalanced(node.Right)
	}

	if node.Left != nil {
		leftIsOk = node.Data > node.Left.Data && node.Left.IsBalanced(node.Left)
	}

	return rightIsOk && leftIsOk
}
