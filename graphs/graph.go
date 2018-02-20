package graphs

import (
	"fmt"

	"github.com/eyoung8/go-fun/queue"
)

type Graph struct {
	Nodes []*Node
}

func (g *Graph) DepthFirstSearchPrint() {
	for _, n := range g.Nodes {
		depthFirstSearchPrint(n)
	}
}

func depthFirstSearchPrint(n *Node) {
	if n == nil || n.visited {
		return
	}
	visitPrint(n)
	for _, adj := range n.Adjacent {
		depthFirstSearchPrint(adj)
	}
}

func visitPrint(n *Node) {
	fmt.Println(n.Name)
	n.visited = true
}

func (g *Graph) BreadthFirstSearchPrint() {
	if len(g.Nodes) == 0 {
		return
	}
	breadthFirstSearchPrint(g.Nodes[0])

}
func breadthFirstSearchPrint(n *Node) {
	var queue queue.Queue
	n.marked = true
	queue.Add(n)
	for queue.IsEmpty() != true {
		newNodeInterface, _ := queue.Remove() //check done on previous line
		if newNodeInterface != nil {
			newNode := newNodeInterface.(*Node)
			visitPrint(newNode)
			for _, adj := range newNode.Adjacent {
				if adj != nil && adj.marked == false {
					adj.marked = true
					queue.Add(adj)
				}
			}
		}
	}
}
