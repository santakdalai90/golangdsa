package binary_search_tree

import (
    "fmt"
)

type Node struct {
    Val int
    Left *Node
    Right *Node
}

type BST struct {
    Root *Node
}

func New() *BST {
    return &BST{nil}
}

func insert(root *Node, val int) *Node{
    if root == nil {
        return &Node{val, nil, nil}
    }
    if root.Val < val {
        root.Right = insert(root.Right, val)
    }
    if root.Val > val {
        root.Left = insert(root.Left, val)
    }
    return root
}

func (b* BST) Insert(val int) {
    b.Root = insert(b.Root, val)
}

func print(root *Node) {
    if root == nil {
        return
    }
    p := fmt.Sprintf("Node: %d", root.Val)
    if root.Left != nil {
        p += fmt.Sprintf("; Left: %d", root.Left.Val)
    } else {
         p += fmt.Sprintf("; Left: nil")
    }
    if root.Right != nil {
        p += fmt.Sprintf("; Right: %d", root.Right.Val)
    } else {
         p += fmt.Sprintf("; Right: nil")
    }
    fmt.Println(p)
    print(root.Left)
    print(root.Right)
}

func (b *BST) Search(val int) *Node{
    n := search(b.Root, val)
    if n != nil {
        fmt.Printf("Found at Node: %v, value: %d\n", n, n.Val)
    } else {
        fmt.Printf("%d not found in the tree\n", val)
    }
    return n
}

func search(root *Node, val int) *Node {
    if root == nil {
        return nil
    }

    if root.Val == val {
        return root
    }

    if root.Val < val {
        return search(root.Right, val)
    }

    if root.Val > val {
        return search(root.Left, val)
    }

    return nil
}

func (b *BST) Print(){
    fmt.Println("||--------*** Printing BST ***--------||")
    print(b.Root)

    fmt.Println("||--------********************--------||")
}