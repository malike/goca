package adjancencymatrix

import (
	"errors"
	"goca/graphs"
)

type AdjacencyMatrix struct {
	Vertices  int
	GraphType graphs.GraphType
}

var AdjMatrix [][]int

func (adjacencyMatrix AdjacencyMatrix) Init() {
	AdjMatrix = make([][]int, adjacencyMatrix.Vertices)
	i := 0
	for i < adjacencyMatrix.Vertices {
		AdjMatrix[i] = make([]int, adjacencyMatrix.Vertices)
		i++
	}
}

func (adjacencyMatrix AdjacencyMatrix) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= adjacencyMatrix.Vertices || vertexTwo >= adjacencyMatrix.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	AdjMatrix[vertexOne][vertexTwo] = 1
	if adjacencyMatrix.GraphType == graphs.UNDIRECTED {

		AdjMatrix[vertexTwo][vertexOne] = 1
	}
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	if vertexOne >= adjacencyMatrix.Vertices || vertexTwo >= adjacencyMatrix.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	AdjMatrix[vertexOne][vertexTwo] = weight
	if adjacencyMatrix.GraphType == graphs.UNDIRECTED {

		AdjMatrix[vertexTwo][vertexOne] = weight
	}
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) RemoveEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= adjacencyMatrix.Vertices || vertexTwo >= adjacencyMatrix.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	AdjMatrix[vertexOne][vertexTwo] = 0
	if adjacencyMatrix.GraphType == graphs.UNDIRECTED {

		AdjMatrix[vertexTwo][vertexOne] = 0
	}
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) HasEdge(vertexOne int, vertexTwo int) bool {
	if vertexOne >= adjacencyMatrix.Vertices || vertexTwo >= adjacencyMatrix.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	return (AdjMatrix[vertexOne][vertexTwo] != 0)
}

func (adjacencyMatrix AdjacencyMatrix) GetGraphType() graphs.GraphType {
	return adjacencyMatrix.GraphType
}

func (adjacencyMatrix AdjacencyMatrix) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	adjacencyMatrixVertices := map[int]bool{}
	if vertex >= adjacencyMatrix.Vertices || vertex < 0 {
		return adjacencyMatrixVertices
	}
	for i := 0; i < adjacencyMatrix.Vertices; i++ {
		if AdjMatrix[vertex][i] != 0 {
			adjacencyMatrixVertices[i] = (AdjMatrix[vertex][i] != 0)
		}
	}
	return adjacencyMatrixVertices
}

func (adjacencyMatrix AdjacencyMatrix) GetOutgoingNodesForVertex(vertex int) map[int]bool {
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) GetWeightOfEdge(vertexOne int, vertexTwo int) (int,error) {
	if vertexOne >= adjacencyMatrix.Vertices || vertexTwo >= adjacencyMatrix.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0,errors.New("Error getting weight for vertex")
	}
	return AdjMatrix[vertexOne][vertexTwo],nil
}

func (adjacencyMatrix AdjacencyMatrix) GetNumberOfVertices() int {
	return adjacencyMatrix.Vertices
}

func (adjacencyMatrix AdjacencyMatrix) GetIndegreeForVertex(vertex int) int {
	indegree := 0
	adjacentNodes := adjacencyMatrix.GetAdjacentNodesForVertex(vertex)
	for key := range adjacentNodes {
		if adjacentNodes[key] {
			indegree++
		}
	}
	return indegree
}

func (adjacencyMatrix AdjacencyMatrix) GetOutdegreeForVertex(vertex int) int {
	return 1
}
