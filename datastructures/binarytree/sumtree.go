package binarytree

type sumFunc[T any] func(...T) T

func (b *Node[T]) IsSumTree(sf sumFunc[T], zero T, compare func(T, T) int) bool {
	if b == nil || (b.left == nil && b.right == nil) {
		return true
	}

	childrenSum := sf(b.left.sum(sf, zero), b.right.sum(sf, zero))
	isLeftSumTree := b.left.IsSumTree(sf, zero, compare)
	isRightSumTree := b.right.IsSumTree(sf, zero, compare)
	return compare(b.data, childrenSum) == 0 && isLeftSumTree && isRightSumTree
}

func (b *Node[T]) sum(sf sumFunc[T], zero T) T {
	if b == nil {
		return zero
	}

	return sf(b.left.sum(sf, zero), b.right.sum(sf, zero), b.data)
}
