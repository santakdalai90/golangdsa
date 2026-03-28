package binarytree

func (b *Node[T]) IsMirror(c *Node[T], compare func(T,T) int) bool {
    if b == nil && c == nil {
        return true
    }

    if b == nil || c == nil {
        return false
    }

    return compare(b.data, c.data) == 0 && 
        b.left.IsMirror(c.right, compare) && b.right.IsMirror(c.left, compare)
}