package traversalbtree

type BTree struct {
	Value int
	Left  *BTree
	Right *BTree
}

func solution(t *BTree) int {
	// result := recursion(t, t.Value, 0)
	result := iterate(t, t.Value)
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

func iterate(t *BTree, root int) int {
	count := 0
	if t == nil {
		return count
	}

	parents := []*BTree{}
	direction := []bool{} // Right: true, Left: false
	curr := t
	for curr != nil {
		if curr.Left != nil {
			parents = append(parents, curr)
			direction = append(direction, false)
			curr = curr.Left
			continue
		}

		if curr.Right != nil {
			parents = append(parents, curr)
			direction = append(direction, true)
			curr = curr.Right
			continue
		}

		if curr.Value >= root {
			count += 1
		}

		if len(parents) == 0 {
			curr = nil
			continue
		}

		curr = parents[len(parents)-1]
		parents = parents[:len(parents)-1]
		goRight := direction[len(direction)-1]
		direction = direction[:len(direction)-1]
		if goRight {
			curr.Right = nil
		} else {
			curr.Left = nil
		}
	}

	return count
}
