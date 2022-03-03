package binary_search_tree

import (
    "fmt"
    "encoding/json"
    "log"
)

type Node struct {
    Val int `json:"Val"`
    Left *Node `json:"Left"`
    Right *Node `json:"Right"`
}

type BST struct {
    Root *Node  `json:"Root"`
}

func New() *BST {
    return &BST{nil}
}

func (b* BST) Insert(val int) {
    b.Root = insert(b.Root, val)
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

func (b *BST) Remove(val int) {
    b.Root = remove(b.Root, val)
}

func (b *BST) String() string{
    // visualize the output using https://vanya.jp.net/vtree/
    raw,err := json.MarshalIndent(b.Root, "", "\t")
    if err != nil {
        log.Println("Error marshalling the tree")
        return ""
    }
    return string(raw)
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

func remove(root *Node, val int) *Node{
    if root == nil {
        return root
    }
    if root.Val > val {
        root.Left = remove(root.Left, val)
    } else if root.Val < val {
        root.Right = remove(root.Right, val)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }

        min := getMinNode(root.Right)
        root.Val = min.Val
        root.Right = remove(root.Right, min.Val)
    }
    return root
}

func getMinNode(root *Node) *Node{
    for root.Left != nil {
        root = root.Left
    }
    return root
}