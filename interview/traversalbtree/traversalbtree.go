package traversalbtree

import (
	"fmt"
)

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

// Not finish
func iterate(t *BTree) {
	if t == nil {
		return
	}

	parents := []*BTree{}
	nums := []int{}

	curr := t
	hasLeft := true
	for len(parents) > 0 || hasLeft {
		if hasLeft {
			nums = append(nums, curr.Value)
			parents = append(parents, curr)
		}

		if curr.Left != nil && hasLeft {
			curr = curr.Left
			continue
		}

		if curr.Right != nil {
			curr = curr.Right
			hasLeft = true
			continue
		}

		hasLeft = false
		curr = parents[len(parents)-1]
		parents = parents[:len(parents)-1]
	}

	fmt.Println(nums)
}
