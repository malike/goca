package adjancencylist

import (
	"goca/graphs"
	"testing"
)

var testAdjMapDirected = AdjacencyMap{4, graphs.DIRECTED}
var testAdjMapUnDirected = AdjacencyMap{4, graphs.UNDIRECTED}

func TestAdjacencyList_AddEdgeDirectedGreaterThanVertex(t *testing.T) {
	testAdjMapDirected.Init()
	err := testAdjMapDirected.AddEdge(10, 2)
	if err == nil {
		t.Error("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyList_AddEdgeDirected(t *testing.T) {
	testAdjMapDirected.Init()
	err := testAdjMapDirected.AddEdge(2, 1)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjList[2].Next.Key != 1 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyList_AddEdgeUndirected(t *testing.T) {

}

func TestAdjacencyList_AddEdgeWeightDirectedGreaterThanVertex(t *testing.T) {

}

func TestAdjacencyList_AddEdgeWithWeightDirected(t *testing.T) {

}

func TestAdjacencyList_AddEdgeWithWeightUnDirected(t *testing.T) {

}

func TestAdjacencyList_RemoveEdgeDirected(t *testing.T) {

}

func TestAdjacencyList_RemoveEdgeUnDirected(t *testing.T) {

}

func TestAdjacencyList_HasEdgeDirected(t *testing.T) {

}

func TestAdjacencyList_HasEdgeUndirected(t *testing.T) {

}

func TestAdjacencyList_GetGraphType(t *testing.T) {

}

func TestAdjacencyList_GetAdjacentVerticesNodesForVertexUndirected(t *testing.T) {

}

func TestAdjacencyList_GetWeightOfEdgeDirected(t *testing.T) {

}

func TestAdjacencyList_GetWeightOfEdgeUndirected(t *testing.T) {

}

func TestAdjacencyList_GetNumberOfVertices(t *testing.T) {

}

func TestAdjacencyList_GetIndegreeForVertex(t *testing.T) {

}

func TestAdjacencyList_GetOutdegreeForVertex(t *testing.T) {

}
