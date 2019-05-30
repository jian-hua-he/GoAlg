package btree

type Node struct {
	key   string
	left  *Node
	right *Node
}
