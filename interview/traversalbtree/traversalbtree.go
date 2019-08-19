package traversalbtree

type BTree struct {
	Value int
	Left  *BTree
	Right *BTree
}

func solution(t *BTree) int {
	result := recursion(t, t.Value, 0)
	return result
}

func recursion(t *BTree, root int, count int) int {
	result := count
	if t == nil {
		return result
	}

	if t.Value >= root {
		result += 1
	}

	return result + recursion(t.Left, root, 0) + recursion(t.Right, root, 0)
}
