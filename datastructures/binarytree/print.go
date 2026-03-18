package binarytree

import (
	"fmt"
	"io"
	"os"
)

func (b *Node[T]) writeDot(w io.Writer) {
	fmt.Fprintln(w, "digraph G {")
	fmt.Fprintln(w, "  node [shape=circle];")
	b.writeDotNodes(w)
	fmt.Fprintln(w, "}")
}

func (b *Node[T]) writeDotNodes(w io.Writer) {
	if b == nil {
		return
	}

	// Create a unique ID for the node (using its pointer address)
	nodeID := fmt.Sprintf("node_%p", b)
	fmt.Fprintf(w, "  %s [label=\"%v\"];\n", nodeID, b.data)

	if b.left != nil {
		leftID := fmt.Sprintf("node_%p", b.left)
		fmt.Fprintf(w, "  %s -> %s [label=\"L\"];\n", nodeID, leftID)
		b.left.writeDotNodes(w)
	}

	if b.right != nil {
		rightID := fmt.Sprintf("node_%p", b.right)
		fmt.Fprintf(w, "  %s -> %s [label=\"R\"];\n", nodeID, rightID)
		b.right.writeDotNodes(w)
	}
}

func (b *Node[T]) Print() {
	f, _ := os.Create("tree.dot")
	b.writeDot(f)
}
