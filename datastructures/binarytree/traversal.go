package binarytree

import "container/list"

type TraverseFunc[T any] func(*Node[T])

func (b *Node[T]) PreOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	tf(b)
	b.left.PreOrderTraversal(tf)
	b.right.PreOrderTraversal(tf)
}

func (b *Node[T]) PostOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	b.left.PostOrderTraversal(tf)
	b.right.PostOrderTraversal(tf)
	tf(b)
}

func (b *Node[T]) InOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	b.left.InOrderTraversal(tf)
	tf(b)
	b.right.InOrderTraversal(tf)
}

func (b *Node[T]) MorrisInOrderTraversal(tf TraverseFunc[T]) {
	cur := b

	for cur != nil {
		if cur.left == nil {
			tf(cur)
			cur = cur.right
		} else {
			prev := cur.left
			for prev.right != nil && prev.right != cur {
				prev = prev.right
			}

			if prev.right == nil {
				prev.right = cur // thread from last node of in order traversal of left subtree to cur
				cur = cur.left
			} else {
				prev.right = nil
				tf(cur)
				cur = cur.right
			}
		}
	}
}

func (b *Node[T]) LevelOrderTraversal(tf TraverseFunc[T]) {
	if b == nil {
		return
	}

	queue := list.New()
	queue.PushBack(b)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*Node[T])
		tf(node)
		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
		queue.Remove(queue.Front())
	}
}
