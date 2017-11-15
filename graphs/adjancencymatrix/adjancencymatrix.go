package adjancencymatrix

import "gods/graphs"

type AdjacencyMatrix struct {
	Vertices  int
	GraphType graphs.GraphType
}

func (adjacencyMatrix AdjacencyMatrix) GetAdjacentVerticesNodesForVertex(vertex int) []int {
	return []int{1, 2}
}

func (adjacencyMatrix AdjacencyMatrix) GetWeightedOfEdge(vertexOne int, vertexTwo int) int {
	return 1
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

func (adjacencyMatrix AdjacencyMatrix) AddEdge(vertexOne int, vertexTwo int) {
	return 1
}

func (adjacencyMatrix AdjacencyMatrix) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) {
	return 1
}

func (adjacencyMatrix AdjacencyMatrix) RemoveEdge(vertexOne int, vertexTwo int) {

}

func (adjacencyMatrix AdjacencyMatrix) HasEdge(vertexOne int, vertexTwo int) bool {
	return 1
}

func (adjacencyMatrix AdjacencyMatrix) GetGraphType() graphs.GraphType {
	return graphs.DIRECTED
}
