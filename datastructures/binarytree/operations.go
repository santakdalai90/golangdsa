package binarytree

func (b *Node[T]) MaxDepth() int {
	if b == nil {
		return -1
	}
	return max(b.left.MaxDepth(), b.right.MaxDepth()) + 1
}

func (b *Node[T]) Size() int {
	if b == nil {
		return 0
	}
	return b.left.Size() + b.right.Size() + 1
}

func (b *Node[T]) Max(compare func(T, T) int) *Node[T] {
	if b == nil {
		return nil
	}
	maxNode := b
	leftMax := b.left.Max(compare)
	rightMax := b.right.Max(compare)

	if leftMax != nil {
		if compare(leftMax.data, maxNode.data) == 1 {
			maxNode = leftMax
		}
	}

	if rightMax != nil {
		if compare(rightMax.data, maxNode.data) == 1 {
			maxNode = rightMax
		}
	}

	return maxNode
}

func (b *Node[T]) IsIdentical(c *Node[T], compare func(T, T) int) bool {
	if b == nil && c == nil {
		return true
	}
	if b == nil {
		return false
	}
	if c == nil {
		return false
	}

	return compare(b.data, c.data) == 0 &&
		b.left.IsIdentical(c.left, compare) &&
		b.right.IsIdentical(c.right, compare)
}

func (b *Node[T]) Find(x T, cf func(T, T) int) *Node[T] {
	if b == nil {
		return nil
	}

	if cf(b.data, x) == 0 {
		return b
	}

	l := b.left.Find(x, cf)
	if l != nil {
		return l
	}

	r := b.right.Find(x, cf)
	if r != nil {
		return r
	}

	return nil
}
