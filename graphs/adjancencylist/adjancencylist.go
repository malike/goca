package adjancencylist

/*
*
* Adjancency List is one of the representations of Graphs. Although this is not an algorithm.
* We'll need this to test all the Graph Algorithms.
*
 */

import "goca/graphs"

type AdjacencyList struct {
	Vertices  int
	GraphType graphs.GraphType
}

func (adjacencyList AdjacencyList) GetAdjacentVerticesNodesForVertex(vertex int) []int {
	return []int{1, 2}
}

func (adjacencyList AdjacencyList) GetWeightedOfEdge(vertexOne int, vertexTwo int) int {
	return 1
}

func (adjacencyList AdjacencyList) GetNumberOfVertices() int {
	return 1
}

func (adjacencyList AdjacencyList) GetIndegreeForVertex(vertex int) int {
	return 1
}

func (adjacencyList AdjacencyList) GetOutdegreeForVertex(vertex int) int {
	return 1
}

func (adjacencyList AdjacencyList) AddEdge(vertexOne int, vertexTwo int) {
	return 1
}

func (adjacencyList AdjacencyList) AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) {
	return 1
}

func (adjacencyList AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) {

}

func (adjacencyList AdjacencyList) HasEdge(vertexOne int, vertexTwo int) bool {
	return 1
}

func (adjacencyList AdjacencyList) GetGraphType() graphs.GraphType {
	return graphs.DIRECTED
}
