// treesort implements a simple binary search tree sort algorithm.
package treesort

// Note: A named struct type S can't declare a field of the same type S:
// an aggregate value type cannot contain itself.
// But S may declare a field of type *S, which is a pointer to S.
// which helps us create recursive data structures like linked lists and trees.
// the illustration below uses a binary tree to implement an insertion sort.

type tree struct {
	val         int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, val := range values {
		root = add(root, val)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, root *tree) []int {
	if root == nil {
		return values
	}
	values = appendValues(values, root.left)
	values = append(values, root.val)
	values = appendValues(values, root.right)
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		return &tree{val: v}
	}
	if v < t.val {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}
