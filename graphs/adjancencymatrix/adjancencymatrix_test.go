package adjancencymatrix

import (
	"goca/graphs"
	"testing"
)

var testAdjMatrixDirected = AdjacencyMatrix{4, graphs.DIRECTED}
var testAdjaMatrixUnDirected = AdjacencyMatrix{4, graphs.UNDIRECTED}

func TestAdjacencyMatrix_AddEdgeDirectedGreaterThanVertex(t *testing.T) {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdge(10, 2)
	if err == nil {
		t.Error("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyMatrix_AddEdgeDirected(t *testing.T) {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdge(2, 1)
	if err != nil {
		t.Error("Index was out of bounds it should have failed")
	}
	if AdjMatrix[2][1] != 1 {
		t.Error("Data not found at index")
	}
	if AdjMatrix[1][2] != 0 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeUndirectedDirected(t *testing.T) {
	testAdjaMatrixUnDirected.Init()
	err := testAdjaMatrixUnDirected.AddEdge(2, 1)
	if err != nil {
		t.Error("Index was out of bounds it should have failed")
	}
	if AdjMatrix[2][1] != 1 {
		t.Error("Data not found at index")
	}
	if AdjMatrix[1][2] != 1 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeWithWeight(t *testing.T) {

}

func TestAdjacencyMatrix_HasEdge(t *testing.T) {

}

func TestAdjacencyMatrix_RemoveEdge(t *testing.T) {

}

func TestAdjacencyMatrix_GetGraphType(t *testing.T) {

}

func TestAdjacencyMatrix_GetAdjacentVerticesNodesForVertex(t *testing.T) {

}

func TestAdjacencyMatrix_GetWeightOfEdge(t *testing.T) {

}

func TestAdjacencyMatrix_GetNumberOfVertices(t *testing.T) {

}

func TestAdjacencyMatrix_GetIndegreeForVertex(t *testing.T) {

}

func TestAdjacencyMatrix_GetOutdegreeForVertex(t *testing.T) {

}
