package graphs

import (
	"errors"
)

type AdjacencyList struct {
	Vertices  int
	GraphType GraphType
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
	if adjacencyList.GraphType == UNDIRECTED {
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
	if adjacencyList.GraphType == UNDIRECTED {
		node := AdjList[vertexTwo].AddNodeWithWeight(vertexOne, weight)
		AdjList[vertexTwo] = node
	}
	return nil
}

func (adjacencyList AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	nodeAdj := AdjList[vertexOne]
	if nodeAdj == (Node{}) {
		return errors.New("Node not found")
	}
	nextNode := &nodeAdj
	newNodes := Node{}
	for nextNode != (&Node{}) && nextNode != nil {
		if nextNode.Key != vertexTwo {
			newNodes.Next = nextNode
			newNodes.Key = nextNode.Key
			newNodes.Weight = nextNode.Weight
			nextNode = nextNode.Next
		} else {
			newNodes.Next = nextNode.Next
			newNodes.Key = nextNode.Next.Key
			newNodes.Weight = nextNode.Next.Weight
			AdjList[vertexOne] = newNodes
			return nil
		}
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

func (adjacencyList AdjacencyList) GetGraphType() GraphType {
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

func (adjacencyList AdjacencyList) GetWeightOfEdge(vertexOne int, vertexTwo int) (int, error) {
	if vertexOne >= adjacencyList.Vertices || vertexTwo >= adjacencyList.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for vertex")
	}
	nodeAdj := AdjList[vertexOne]
	if nodeAdj == (Node{}) {
		return 0, errors.New("Error getting weight for vertex")
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return nodeAdj.Weight, nil
	}
	return 0, errors.New("Error getting weight for vertex")
}

func (adjacencyList AdjacencyList) GetNumberOfVertices() int {
	return adjacencyList.Vertices
}

func (adjacencyList AdjacencyList) GetIndegreeForVertex(vertex int) int {
	if vertex >= adjacencyList.Vertices || vertex < 0 {
		return 0
	}
	nodeAdj := AdjList[vertex]
	nextNode := nodeAdj.Next
	length := 0
	for nextNode != (&Node{}) && nextNode != nil {
		length += 1
		nextNode = nextNode.Next
	}
	return length
}
