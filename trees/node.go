package trees

import (
	"fmt"
)

type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

func NewNode(val interface{}) *Node {
	return &Node{Val: val}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Val)
}
