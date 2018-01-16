package adjancencylist

import (
	"goca/graphs"
	"testing"
)

var testAdjListDirected = AdjacencyList{4, graphs.DIRECTED}
var testAdjListUnDirected = AdjacencyList{4, graphs.UNDIRECTED}

func TestAdjacencyList_AddEdgeDirectedGreaterThanVertex(t *testing.T) {
	testAdjListDirected.Init()
	err := testAdjListDirected.AddEdge(10, 2)
	if err == nil {
		t.Error("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyList_AddEdgeDirected(t *testing.T) {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	err := testAdjListDirected.AddEdge(2, 3)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjList[2].Key != 1 {
		t.Error("Data not found at index")
	}
	if AdjList[2].Next.Key != 3 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyList_AddEdgeUndirected(t *testing.T) {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdge(2, 1)
	err := testAdjListUnDirected.AddEdge(2, 3)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjList[2].Key != 1 {
		t.Error("Data not found at index")
	}
	if AdjList[2].Next.Key != 3 {
		t.Error("Data not found at index")
	}
	if AdjList[1].Key != 2 {
		t.Error("Data not found at index")
	}
	if AdjList[3].Key != 2 {
		t.Error("Data not found at index")
	}
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
