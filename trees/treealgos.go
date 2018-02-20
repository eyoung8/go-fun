package trees

import (
	"fmt"
)

func (n *Node) InOrderTraversal() {
	inOrderTraversal(n)
}

func inOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	inOrderTraversal(n.Left)
	fmt.Println(n)
	inOrderTraversal(n.Right)
}

func (n *Node) PreOrderTraversal() {
	preOrderTraversal(n)
}

func preOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n)
	preOrderTraversal(n.Left)
	preOrderTraversal(n.Right)
}

func (n *Node) PostOrderTraversal() {
	postOrderTraversal(n)
}

func postOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	postOrderTraversal(n.Left)
	postOrderTraversal(n.Right)
	fmt.Println(n)
}
