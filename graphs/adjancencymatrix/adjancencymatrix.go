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
	i := 1
	for i < adjacencyMatrix.Vertices {
		AdjMatrix[i] = make([]int, adjacencyMatrix.Vertices)
		i++
	}
}

func (adjacencyMatrix AdjacencyMatrix) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne > adjacencyMatrix.Vertices || vertexTwo > adjacencyMatrix.Vertices {
		return errors.New("Index out of bounds")
	}

	AdjMatrix[vertexOne] = make([]int, adjacencyMatrix.Vertices)

	AdjMatrix[vertexOne][vertexTwo] = 1
	if adjacencyMatrix.GraphType == graphs.UNDIRECTED {

		AdjMatrix[vertexTwo][vertexOne] = 1
	}
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) RemoveEdge(vertexOne int, vertexTwo int) error {
	return nil
}

func (adjacencyMatrix AdjacencyMatrix) HasEdge(vertexOne int, vertexTwo int) bool {
	return false
}

func (adjacencyMatrix AdjacencyMatrix) GetGraphType() graphs.GraphType {
	return adjacencyMatrix.GraphType
}

func (adjacencyMatrix AdjacencyMatrix) GetAdjacentVerticesNodesForVertex(vertex int) []int {
	return []int{1, 2}
}

func (adjacencyMatrix AdjacencyMatrix) GetWeightOfEdge(vertexOne int, vertexTwo int) int {
	return AdjMatrix[vertexOne][vertexTwo]
}

func (adjacencyMatrix AdjacencyMatrix) GetNumberOfVertices() int {
	return 1
}

func (adjacencyMatrix AdjacencyMatrix) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencyMatrix AdjacencyMatrix) GetOutdegreeForVertex(vertex int) int {
	return 1
}
