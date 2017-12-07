package adjancencyset

import "goca/graphs"

type AdjacencySet struct {
	Vertices  int
	GraphType graphs.GraphType
}

func (adjacencySet AdjacencySet) GetAdjacentVerticesNodesForVertex(vertex int) []int {
	return []int{1, 2}
}

func (adjacencySet AdjacencySet) GetWeightOfEdge(vertexOne int, vertexTwo int) int {
	return 1
}

func (adjacencySet AdjacencySet) GetNumberOfVertices() int {
	return 1
}

func (adjacencySet AdjacencySet) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencySet AdjacencySet) GetOutdegreeForVertex(vertex int) int {
	return 1
}

func (adjacencySet AdjacencySet) AddEdge(vertexOne int, vertexTwo int) {
	return 1
}

func (adjacencySet AdjacencySet) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) {
	return 1
}

func (adjacencySet AdjacencySet) RemoveEdge(vertexOne int, vertexTwo int) {

}

func (adjacencySet AdjacencySet) HasEdge(vertexOne int, vertexTwo int) bool {
	return 1
}

func (adjacencySet AdjacencySet) GetGraphType() graphs.GraphType {
	return graphs.DIRECTED
}
