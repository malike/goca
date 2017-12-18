package adjancencymatrix

import (
	"goca/graphs"
	"testing"
)

var testAdjMatrixDirected = AdjacencyMatrix{4, graphs.DIRECTED}
var testAdjMatrixUnDirected = AdjacencyMatrix{4, graphs.UNDIRECTED}

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
		t.Error("Error adding edge")
	}
	if AdjMatrix[2][1] != 1 {
		t.Error("Data not found at index")
	}
	if AdjMatrix[1][2] != 0 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeUndirected(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	err := testAdjMatrixUnDirected.AddEdge(2, 1)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjMatrix[2][1] != 1 {
		t.Error("Data not found at index")
	}
	if AdjMatrix[1][2] != 1 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeWeightDirectedGreaterThanVertex(t *testing.T) {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdgeWithWeight(10, 7, -3)
	if err == nil {
		t.Error("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyMatrix_AddEdgeWithWeightDirected(t *testing.T) {
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdgeWithWeight(2, 1, -3)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjMatrix[2][1] != -3 {
		t.Error("Data not found at index")
	}
	if AdjMatrix[1][2] != 0 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_AddEdgeWithWeightUnDirected(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	err := testAdjMatrixUnDirected.AddEdgeWithWeight(2, 1, -3)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjMatrix[2][1] != -3 {
		t.Error("Data not found at index")
	}
	if AdjMatrix[1][2] != -3 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_RemoveEdgeDirected(t *testing.T) {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	err := testAdjMatrixDirected.RemoveEdge(2, 1)
	if err != nil {
		t.Error("Error removing edge")
	}
	if AdjMatrix[2][1] != 0 {
		t.Error("Data not removed at index")
	}
}

func TestAdjacencyMatrix_RemoveEdgeUnDirected(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdge(2, 1)
	err := testAdjMatrixUnDirected.RemoveEdge(2, 1)
	if err != nil {
		t.Error("Error removing edge")
	}
	if AdjMatrix[2][1] != 0 {
		t.Error("Data not removed at index")
	}
	if AdjMatrix[1][2] != 0 {
		t.Error("Data not removed at index")
	}
}

func TestAdjacencyMatrix_HasEdgeDirected(t *testing.T) {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdge(2, 1)
	if !testAdjMatrixDirected.HasEdge(2, 1) {
		t.Error("No relationship, when there should be one")
	}
	if testAdjMatrixDirected.HasEdge(1, 2) {
		t.Error("Relationship, when there shouldn't be one")
	}
}

func TestAdjacencyMatrix_HasEdgeUndirected(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdge(2, 1)
	if !testAdjMatrixUnDirected.HasEdge(2, 1) {
		t.Error("No relationship, when there should be one")
	}
	if !testAdjMatrixUnDirected.HasEdge(1, 2) {
		t.Error("No relationship, when there should be one")
	}
}

func TestAdjacencyMatrix_GetGraphType(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	if testAdjMatrixUnDirected.GraphType != testAdjMatrixUnDirected.GetGraphType() {
		t.Error("GraphType not matching")
	}
}

func TestAdjacencyMatrix_GetAdjacentVerticesNodesForVertexUndirected(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdge(2, 1)
	testAdjMatrixUnDirected.AddEdge(2, 3)
	testAdjMatrixUnDirected.AddEdge(1, 1)
	testAdjMatrixUnDirected.AddEdge(2, 4)
	nodes := testAdjMatrixUnDirected.GetAdjacentVerticesNodesForVertex(2)
	if len(nodes) != 3 {
		t.Error("Nodes size not matching")
	}

}

func TestAdjacencyMatrix_GetWeightOfEdgeDirected(t *testing.T) {
	testAdjMatrixDirected.Init()
	testAdjMatrixDirected.AddEdgeWithWeight(2, 1, -3)
	weight := testAdjMatrixDirected.GetWeightOfEdge(2, 1)
	if weight != -3 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_GetWeightOfEdgeUndirected(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	testAdjMatrixUnDirected.AddEdgeWithWeight(2, 1, -3)
	weight := testAdjMatrixUnDirected.GetWeightOfEdge(2, 1)
	if weight != -3 {
		t.Error("Data not found at index")
	}
}

func TestAdjacencyMatrix_GetNumberOfVertices(t *testing.T) {
	testAdjMatrixUnDirected.Init()
	if testAdjMatrixUnDirected.GetNumberOfVertices() != testAdjMatrixUnDirected.Vertices {
		t.Error("Vertex count doesn't match")
	}
}

func TestAdjacencyMatrix_GetIndegreeForVertex(t *testing.T) {

}

func TestAdjacencyMatrix_GetOutdegreeForVertex(t *testing.T) {

}
