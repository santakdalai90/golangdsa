package binary_search_tree_test

import(
    "testing"
    BST "golang-dsa/src/data_structures/binary_search_tree"
)

// go test -v ./test/data_structures_test/...

func Test_Insert(t *testing.T) {
    bst := BST.New()
    bst.Insert(9)
    bst.Insert(4)
    bst.Insert(6)
    bst.Insert(20)
    bst.Insert(170)
    bst.Insert(15)
    bst.Insert(1)
    t.Log(bst)
}

func Test_Search(t *testing.T) {
    bst := BST.New()
    bst.Insert(9)
    bst.Insert(4)
    bst.Insert(6)
    bst.Insert(20)
    bst.Insert(170)
    bst.Insert(15)
    bst.Insert(1)
    
    bst.Search(170)
    bst.Search(150)
    bst.Insert(150)
    bst.Search(150)
}

func Test_Delete(t *testing.T) {
    bst := BST.New()
    bst.Insert(9)
    bst.Insert(4)
    bst.Insert(6)
    bst.Insert(20)
    bst.Insert(170)
    bst.Insert(15)
    bst.Insert(1)
    t.Log(bst)

    bst.Remove(9)
    t.Log(bst)
}