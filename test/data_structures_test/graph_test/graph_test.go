package graph_test

import (
	"golang-dsa/src/data_structures/graph"
	"testing"
)

// go test -v ./test/data_structures_test/graph_test/...

func Test_Insert(t *testing.T) {
	g := graph.New()

	g.AddVertex("0")
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")
	g.AddVertex("6")

	g.AddEdge("0", "1")
	g.AddEdge("0", "2")
	g.AddEdge("1", "2")
	g.AddEdge("1", "3")
	g.AddEdge("2", "4")
	g.AddEdge("3", "4")
	g.AddEdge("4", "5")
	g.AddEdge("5", "6")

	g.ShowConnections()
}
