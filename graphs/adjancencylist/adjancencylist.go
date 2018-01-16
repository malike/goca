package adjancencylist

import (
	"errors"
	"goca/graphs"
)

type AdjacencyList struct {
	Vertices  int
	GraphType graphs.GraphType
}

type Node struct {
	Next *Node
	Key  int
}

//Recursive method to add node to last available slot
func (node Node) AddNode(value int) Node {
	n := node.Next
	if (Node{}) == node {
		node := Node{Next: &Node{}, Key: value}
		return node
	}
	nd := n.AddNode(value)
	node.Next = &nd
	return node
}

var AdjList []Node

func (adjacencyList AdjacencyList) Init() {
	AdjList = make([]Node, adjacencyList.Vertices)
	i := 0
	for i < adjacencyList.Vertices {
		AdjList[i] = Node{}
		i++
	}
}

func (adjacencyList AdjacencyList) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	//if AdjList[vertexOne] == (Node{}) {
	//	AdjList[vertexOne] = Node{Next: &Node{}, Key: vertexTwo}
	//} else {
	node := AdjList[vertexOne].AddNode(vertexTwo)
	AdjList[vertexOne] = node
	//}
	return nil
}

func (adjacencyList AdjacencyList) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	return nil
}

func (adjacencyList AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) error {

	return nil
}

func (adjacencyList AdjacencyList) HasEdge(vertexOne int, vertexTwo int) bool {
	return false
}

func (adjacencyList AdjacencyList) GetGraphType() graphs.GraphType {
	return adjacencyList.GraphType
}

func (adjacencyList AdjacencyList) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	return map[int]bool{}
}

func (adjacencyList AdjacencyList) GetOutgoingNodesForVertex(vertex int) map[int]bool {
	return map[int]bool{}
}

func (adjacencyList AdjacencyList) GetWeightOfEdge(vertexOne int, vertexTwo int) int {
	return 0
}

func (adjacencyList AdjacencyList) GetNumberOfVertices() int {
	return 0
}

func (adjacencyList AdjacencyList) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencyList AdjacencyList) GetOutdegreeForVertex(vertex int) int {
	return 1
}
