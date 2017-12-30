package adjancencylist

import "goca/graphs"

type AdjacencyList struct {
	Vertices  int
	GraphType graphs.GraphType
}

func (adjacencyList AdjacencyList) Init() {

}

func (adjacencyList AdjacencyList) AddEdge(vertexOne int, vertexTwo int) error {
	return nil
}

func (adjacencyList AdjacencyList) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error {
	return nil
}

func (adjacencyList AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) error {

	return nil
}

func (adjacencyList AdjacencyList) HasEdge(vertexOne int, vertexTwo int) bool {
	return nil
}

func (adjacencyList AdjacencyList) GetGraphType() graphs.GraphType {
	return nil
}

func (adjacencyList AdjacencyList) GetAdjacentVerticesNodesForVertex(vertex int) map[int]bool {
	return nil
}

func (adjacencyList AdjacencyList) GetWeightOfEdge(vertexOne int, vertexTwo int) int {
	return nil
}

func (adjacencyList AdjacencyList) GetNumberOfVertices() int {
	return nil
}

func (adjacencyList AdjacencyList) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencyList AdjacencyList) GetOutdegreeForVertex(vertex int) int {
	return 1
}
