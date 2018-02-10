package graphs

import (
	"testing"
)

var testAdjMatrixUnDirectedBFT Graphs
var testAdjListUnDirectedBFT Graphs
var testAdjMatrixDirectedBFT Graphs
var testAdjListDirectedBFT Graphs

func Test_BFTAdjacencyMatrixUndirected(t *testing.T) {
	testAdjMatrixUnDirectedBFT = AdjacencyMatrix{4, UNDIRECTED}
	testAdjMatrixUnDirectedBFT.Init()
	testAdjMatrixUnDirectedBFT.AddEdge(2, 1)
	testAdjMatrixUnDirectedBFT.AddEdge(3, 2)
	testAdjMatrixUnDirectedBFT.AddEdge(2, 0)
	testAdjMatrixUnDirectedBFT.AddEdge(1, 0)
	testAdjMatrixUnDirectedBFT.AddEdge(1, 1)
	testAdjMatrixUnDirectedBFT.AddEdge(2, 3)
	t.Logf("Graph after BFT Undirected Matrix %v ", BFT(testAdjMatrixUnDirectedBFT))
}

func Test_BFTAdjacencyListUndirected(t *testing.T) {
	testAdjListUnDirectedBFT = AdjacencyList{4, UNDIRECTED}
	testAdjListUnDirectedBFT.Init()
	testAdjListUnDirectedBFT.AddEdge(2, 1)
	testAdjListUnDirectedBFT.AddEdge(3, 2)
	testAdjListUnDirectedBFT.AddEdge(2, 0)
	testAdjListUnDirectedBFT.AddEdge(1, 0)
	testAdjListUnDirectedBFT.AddEdge(1, 1)
	testAdjListUnDirectedBFT.AddEdge(2, 3)
	t.Logf("Graph after BFT Undirected List %v ", BFT(testAdjListUnDirectedBFT))
}

func Test_BFTAdjacencyMatrixDirected(t *testing.T) {
	testAdjMatrixDirectedBFT = AdjacencyMatrix{4, DIRECTED}
	testAdjMatrixDirectedBFT.Init()
	testAdjMatrixDirectedBFT.AddEdge(2, 1)
	testAdjMatrixDirectedBFT.AddEdge(3, 2)
	testAdjMatrixDirectedBFT.AddEdge(2, 0)
	testAdjMatrixDirectedBFT.AddEdge(1, 0)
	testAdjMatrixDirectedBFT.AddEdge(1, 1)
	testAdjMatrixDirectedBFT.AddEdge(2, 3)
	t.Logf("Graph after BFT Directed Matrix %v ", BFT(testAdjMatrixDirectedBFT))
}

func Test_BFTAdjacencyListDirected(t *testing.T) {
	testAdjListDirectedBFT = AdjacencyList{4, DIRECTED}
	testAdjListDirectedBFT.Init()
	testAdjListDirectedBFT.AddEdge(2, 1)
	testAdjListDirectedBFT.AddEdge(3, 2)
	testAdjListDirectedBFT.AddEdge(2, 0)
	testAdjListDirectedBFT.AddEdge(1, 0)
	testAdjListDirectedBFT.AddEdge(1, 1)
	testAdjListDirectedBFT.AddEdge(2, 3)
	t.Logf("Graph after BFT Directed List %v ", BFT(testAdjListDirectedBFT))
}
