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
	Next   *Node
	Weight int
	Key    int
}

//Recursive method to add node to last available slot
func (node Node) AddNode(value int) Node {
	n := node.Next
	if n == nil {
		newnode := Node{Next: &Node{}, Key: value}
		return newnode
	}
	nd := n.AddNode(value)
	node.Next = &nd
	return node
}

//Recursive method to append with weight
func (node Node) AddNodeWithWeight(value int, weight int) Node {
	n := node.Next
	if n == nil {
		newnode := Node{Next: &Node{}, Key: value, Weight: weight}
		return newnode
	}
	nd := n.AddNodeWithWeight(value, weight)
	node.Next = &nd
	return node
}

func (node Node) FindNextNode(key int) (*Node, error) {
	n := node
	if n == (Node{}) {
		return &Node{}, errors.New("Node not found")
	}
	if n.Key == key {
		return &n, nil
	}
	nd := n.Next
	return nd.FindNextNode(key)
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
	node := AdjList[vertexOne].AddNode(vertexTwo)
	AdjList[vertexOne] = node
	if adjacencyList.GraphType == graphs.UNDIRECTED {
		node := AdjList[vertexTwo].AddNode(vertexOne)
		AdjList[vertexTwo] = node
	}
	return nil
}

func (adjacencyList AdjacencyList) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	node := AdjList[vertexOne].AddNodeWithWeight(vertexTwo, weight)
	AdjList[vertexOne] = node
	if adjacencyList.GraphType == graphs.UNDIRECTED {
		node := AdjList[vertexTwo].AddNodeWithWeight(vertexOne, weight)
		AdjList[vertexTwo] = node
	}
	return nil
}

func (adjacencyList AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	return nil
}

func (adjacencyList AdjacencyList) HasEdge(vertexOne int, vertexTwo int) bool {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	nodeAdj := AdjList[vertexOne]
	if nodeAdj == (Node{}) {
		return false
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return true
	}
	return false
}

func (adjacencyList AdjacencyList) GetGraphType() graphs.GraphType {
	return adjacencyList.GraphType
}

func (adjacencyList AdjacencyList) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	if vertex >= adjacencyList.Vertices || vertex < 0 {
		return map[int]bool{}
	}
	nodeAdj := AdjList[vertex]
	nextNode := &nodeAdj
	nodes := map[int]bool{}
	for nextNode != (&Node{}) && nextNode != nil {
		nodes[nextNode.Key] = true
		nextNode = nextNode.Next
	}
	return nodes

}

func (adjacencyList AdjacencyList) GetOutgoingNodesForVertex(vertex int) map[int]bool {
	return map[int]bool{}
}

func (adjacencyList AdjacencyList) GetWeightOfEdge(vertexOne int, vertexTwo int) (int,error) {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0,errors.New("Error getting weight for vertex")
	}
	nodeAdj := AdjList[vertexOne]
	if nodeAdj == (Node{}) {
		return 0,errors.New("Error getting weight for vertex")
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return nodeAdj.Weight,nil
	}
	return 0,errors.New("Error getting weight for vertex")
}

func (adjacencyList AdjacencyList) GetNumberOfVertices() int {
	return 0
}

func (adjacencyList AdjacencyList) GetIndegreeForVertex(vertex int) int {
	return 1
}
