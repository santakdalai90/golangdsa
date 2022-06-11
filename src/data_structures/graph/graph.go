package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	NumVertices  int
	AdjacentList map[Vertex][]Vertex
}

type Vertex string

func (g *Graph) AddEdge(v1, v2 Vertex) {
	//undirected
	g.AdjacentList[v1] = append(g.AdjacentList[v1], v2)
	g.AdjacentList[v2] = append(g.AdjacentList[v2], v1)
}

func (g *Graph) AddVertex(data interface{}) error {
	newVertex := Vertex(data.(string))
	if _, ok := g.AdjacentList[newVertex]; !ok {
		g.AdjacentList[newVertex] = make([]Vertex, 0)
		g.NumVertices++
		return nil
	}
	return errors.New("Vertex already present")
}

func (g *Graph) ShowConnections() {
    fmt.Println("Number of Vertices:", g.NumVertices)
	for k,v := range(g.AdjacentList) {
        fmt.Printf("%v ---> %v\n", k, v)
    }
}

func New() *Graph {
	return &Graph{
		NumVertices:  0,
		AdjacentList: make(map[Vertex][]Vertex),
	}
}
