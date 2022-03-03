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

func (b *BST) PrintJson() {
    // visualize the output using https://vanya.jp.net/vtree/
    raw,err := json.MarshalIndent(b.Root, "", "\t")
    if err != nil {
        log.Println("Error marshalling the tree")
        return
    }
    log.Println(string(raw))
}

func (b *BST) Remove(val int) {
    b.Root = remove(b.Root, val)
}

func remove(root *Node, val int) *Node{
    // find the node to remove
    var currentNode *Node = root
    var parentNode *Node
    var NTR, successor *Node    // NTR => Node To Remove
    for currentNode != nil{
        if currentNode.Val > val {
            parentNode = currentNode
            currentNode = currentNode.Left
        } else if currentNode.Val < val {
            parentNode = currentNode
            currentNode = currentNode.Right
        } else {
            NTR = currentNode
            break
        }
    }
    if NTR == nil {
        log.Println("Node not found in the tree")
        return root
    }
    // if no children
    if NTR.Left == nil && NTR.Right == nil {
        successor = nil
        if parentNode != nil {    // to handle root node
            if NTR == parentNode.Left {
                parentNode.Left = successor
            } else {
                parentNode.Right = successor
            }
        }
    }
    
    // if single child
    if NTR.Left == nil && NTR.Right != nil {
        successor = NTR.Right
        if parentNode != nil {    // to handle root node
            if NTR == parentNode.Left {
                parentNode.Left = successor
            } else {
                parentNode.Right = successor
            }
        }
    } else if NTR.Right == nil && NTR.Left != nil {
        successor = NTR.Left
        if parentNode != nil {    // to handle root node
            if NTR == parentNode.Left {
                parentNode.Left = successor
            } else {
                parentNode.Right = successor
            }
        }
    }

    
    
    // if both children
    if NTR.Left != nil && NTR.Right != nil {
        //get the min of right sub-tree of NTR as successor
        //Print(root)
        successor = getMinNode(NTR.Right)
        
        // exchange the value from NTR node
        NTR.Val = successor.Val
        // remove the successor
        NTR.Right = remove(NTR.Right, successor.Val)

         //Print(root)
        
        if parentNode == nil {
            // deleting root node
            return NTR
        }
    }

    if parentNode == nil {
        // deleting root node
        return successor
    }

    return root
}

func getMinNode(root *Node) *Node{
    if root == nil || root.Left == nil {
        return root
    }
    return getMinNode(root.Left)
}