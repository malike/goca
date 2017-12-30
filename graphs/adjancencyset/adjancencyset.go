package adjancencyset

import "goca/graphs"

type AdjacencySet struct {
	Vertices  int
	GraphType graphs.GraphType
}

func (adjacencySet AdjacencySet) Init() {

}

func (adjacencySet AdjacencySet) AddEdge(vertexOne int, vertexTwo int) error {
	return nil
}

func (adjacencySet AdjacencySet) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	return nil
}

func (adjacencySet AdjacencySet) RemoveEdge(vertexOne int, vertexTwo int) error {

	return nil
}

func (adjacencySet AdjacencySet) HasEdge(vertexOne int, vertexTwo int) bool {
	return nil
}

func (adjacencySet AdjacencySet) GetGraphType() graphs.GraphType {
	return nil
}

func (adjacencySet AdjacencySet) GetAdjacentVerticesNodesForVertex(vertex int) []int {
	return nil
}

func (adjacencySet AdjacencySet) GetWeightOfEdge(vertexOne int, vertexTwo int) int {
	return nil
}

func (adjacencySet AdjacencySet) GetNumberOfVertices() int {
	return nil
}

func (adjacencySet AdjacencySet) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencySet AdjacencySet) GetOutdegreeForVertex(vertex int) int {
	return 1
}
