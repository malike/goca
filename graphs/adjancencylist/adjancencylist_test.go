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
	testAdjListUnDirected.AddEdge(2, 3)
	err := testAdjListUnDirected.AddEdge(2, 0)
	if err != nil {
		t.Error("Error adding edge " + err.Error())
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
	testAdjListUnDirected.Init()
	err := testAdjListUnDirected.AddEdgeWithWeight(10, 7, -3)
	if err == nil {
		t.Error("Index was out of bounds it should have failed")
	}
}

func TestAdjacencyList_AddEdgeWithWeightDirected(t *testing.T) {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdgeWithWeight(2, 1, -3)
	err := testAdjListDirected.AddEdgeWithWeight(2, 3, 6)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjList[2].Weight != -3 {
		t.Error("Data not found at index")
	}
	if AdjList[2].Next.Weight != 6 {
		t.Errorf("Data not found at index, %d", AdjList[2].Next.Weight)
	}
}

func TestAdjacencyList_AddEdgeWithWeightUnDirected(t *testing.T) {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdgeWithWeight(2, 1, -3)
	err := testAdjListUnDirected.AddEdgeWithWeight(2, 3, 6)
	if err != nil {
		t.Error("Error adding edge")
	}
	if AdjList[2].Weight != -3 {
		t.Error("Data not found at index")
	}
	if AdjList[1].Weight != -3 {
		t.Error("Data not found at index")
	}
	if AdjList[2].Next.Weight != 6 {
		t.Errorf("Data not found at index, %d", AdjList[2].Next.Weight)
	}
	if AdjList[3].Weight != 6 {
		t.Errorf("Data not found at index, %d", AdjList[2].Next.Weight)
	}
}

func TestAdjacencyList_RemoveEdgeDirected(t *testing.T) {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	testAdjListDirected.AddEdge(2, 3)
	err := testAdjListDirected.RemoveEdge(2, 1)
	if err != nil {
		t.Error("Error removing edge")
	}
	if AdjList[2].Key == 1 && AdjList[2].Key != 3 && AdjList[2].Next != (&Node{}) {
		t.Errorf("Data not removed at index , data is %d instead of 3", AdjList[2].Key)
	}
}

func TestAdjacencyList_RemoveEdgeUnDirected(t *testing.T) {

}

func TestAdjacencyList_HasEdgeDirected(t *testing.T) {
	testAdjListDirected.Init()
	testAdjListDirected.AddEdge(2, 1)
	testAdjListDirected.AddEdge(2, 3)
	testAdjListDirected.AddEdge(2, 0)
	if !testAdjListDirected.HasEdge(2, 0) {
		t.Error("No relationship, when there should be one")
	}
	if testAdjListDirected.HasEdge(1, 2) {
		t.Error("Relationship, when there shouldn't be one")
	}
}

func TestAdjacencyList_HasEdgeUndirected(t *testing.T) {
	testAdjListUnDirected.Init()
	testAdjListUnDirected.AddEdge(2, 1)
	testAdjListUnDirected.AddEdge(2, 3)
	testAdjListUnDirected.AddEdge(2, 0)
	if !testAdjListUnDirected.HasEdge(2, 0) {
		t.Error("No relationship, when there should be one")
	}
	if !testAdjListUnDirected.HasEdge(0, 2) {
		t.Error("No relationship, when there should be one")
	}
}

func TestAdjacencyList_GetGraphType(t *testing.T) {
	testAdjListDirected.Init()
	if testAdjListDirected.GraphType != testAdjListDirected.GetGraphType() {
		t.Error("GraphType not matching")
	}
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
