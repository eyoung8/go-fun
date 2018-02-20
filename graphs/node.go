package graphs

// Directional graph node
type Node struct {
	Name     string
	Adjacent []*Node
	visited  bool
	marked   bool
}
