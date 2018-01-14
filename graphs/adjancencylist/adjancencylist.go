package adjancencylist

import (
	"errors"
	"goca/graphs"
)

type AdjacencyMap struct {
	Vertices  int
	GraphType graphs.GraphType
}

type Node struct {
	Next *Node
	Key  int
}

//Recursive method to add node to last availabel slot
func (node Node) AddNode(value int)  {

	if node.Next == nil {
		node.Next = &Node{Next: nil, Key: value}
		return
	}
	node.AddNode(value)
}

var AdjList []Node

func (adjacencyList AdjacencyMap) Init() {
	AdjList := make([]Node, adjacencyList.Vertices)
	i := 0
	for i < adjacencyList.Vertices {
		AdjList[i] = Node{}
		i++
	}
}

func (adjacencyList AdjacencyMap) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	adjacentVertex := AdjList[vertexOne]
	if adjacentVertex == (Node{}) {
		AdjList[vertexOne] = Node{Next: nil, Key: vertexTwo}
		return nil
	}
	adjacentVertex.AddNode(vertexTwo)
	return nil
}

func (adjacencyList AdjacencyMap) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	return nil
}

func (adjacencyList AdjacencyMap) RemoveEdge(vertexOne int, vertexTwo int) error {

	return nil
}

func (adjacencyList AdjacencyMap) HasEdge(vertexOne int, vertexTwo int) bool {
	return false
}

func (adjacencyList AdjacencyMap) GetGraphType() graphs.GraphType {
	return adjacencyList.GraphType
}

func (adjacencyList AdjacencyMap) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	return map[int]bool{}
}

func (adjacencyList AdjacencyMap) GetOutgoingNodesForVertex(vertex int) map[int]bool {
	return map[int]bool{}
}

func (adjacencyList AdjacencyMap) GetWeightOfEdge(vertexOne int, vertexTwo int) int {
	return 0
}

func (adjacencyList AdjacencyMap) GetNumberOfVertices() int {
	return 0
}

func (adjacencyList AdjacencyMap) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencyList AdjacencyMap) GetOutdegreeForVertex(vertex int) int {
	return 1
}
