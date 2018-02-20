package trees

import (
	"fmt"
	"testing"
)

func TestTraversals(t *testing.T) {
	tree := Node{
		2,
		&Node{Val: 1},
		&Node{
			3,
			nil,
			&Node{4,
				nil,
				&Node{Val: 5},
			},
		},
	}
	fmt.Println("In order traversal")
	tree.InOrderTraversal()
	fmt.Println("Pre-order traversal")
	tree.PreOrderTraversal()
	fmt.Println("Post-order traversal")
	tree.PostOrderTraversal()
}
