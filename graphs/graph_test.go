package graphs

import (
	"testing"
)

// a -> b <- c
// | \ / \
// f  e <-d
func makeGraphForTest() *Graph {
	f := &Node{
		"f",
		[]*Node{},
		false,
		false,
	}
	e := &Node{
		"e",
		[]*Node{},
		false,
		false,
	}
	d := &Node{
		"d",
		[]*Node{e},
		false,
		false,
	}
	b := &Node{
		"b",
		[]*Node{d, e},
		false,
		false,
	}
	c := &Node{
		"c",
		[]*Node{b},
		false,
		false,
	}
	a := &Node{
		"a",
		[]*Node{b, e, f},
		false,
		false,
	}
	g := &Graph{
		Nodes: []*Node{a, b, c, d, e, f},
	}
	return g
}

// Should print a, b, d, e, f, c
func TestDepthFirstSearchPrint(t *testing.T) {
	graph := makeGraphForTest()
	graph.DepthFirstSearchPrint()
}

// Should print a, b, e, f, d, (can't reach c)
func TestBreadthFirstSearchPrint(t *testing.T) {
	graph := makeGraphForTest()
	graph.BreadthFirstSearchPrint()
}
