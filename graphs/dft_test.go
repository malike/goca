package graphs

import "testing"

var testAdjMatrixUnDirectedDFT Graphs
var testAdjListUnDirectedDFT Graphs
var testAdjMatrixDirectedDFT Graphs
var testAdjListDirectedDFT Graphs

func Test_DFTAdjacencyMatrixUndirected(t *testing.T) {
	testAdjMatrixUnDirectedDFT = AdjacencyMatrix{4, UNDIRECTED}
	testAdjMatrixUnDirectedDFT.Init()
	testAdjMatrixUnDirectedDFT.AddEdge(2, 1)
	testAdjMatrixUnDirectedDFT.AddEdge(3, 2)
	testAdjMatrixUnDirectedDFT.AddEdge(2, 0)
	testAdjMatrixUnDirectedDFT.AddEdge(1, 0)
	testAdjMatrixUnDirectedDFT.AddEdge(1, 1)
	testAdjMatrixUnDirectedDFT.AddEdge(2, 3)
	t.Logf("Graph after DFT Undirected Matrix %v ", DFT(testAdjMatrixUnDirectedDFT))
}

func Test_DFTAdjacencyListUndirected(t *testing.T) {
	testAdjListUnDirectedDFT = AdjacencyList{4, UNDIRECTED}
	testAdjListUnDirectedDFT.Init()
	testAdjListUnDirectedDFT.AddEdge(2, 1)
	testAdjListUnDirectedDFT.AddEdge(3, 2)
	testAdjListUnDirectedDFT.AddEdge(2, 0)
	testAdjListUnDirectedDFT.AddEdge(1, 0)
	testAdjListUnDirectedDFT.AddEdge(1, 1)
	testAdjListUnDirectedDFT.AddEdge(2, 3)
	t.Logf("Graph after DFT Undirected List %v ", DFT(testAdjListUnDirectedDFT))
}

func Test_DFTAdjacencyMatrixDirected(t *testing.T) {
	testAdjMatrixDirectedDFT = AdjacencyMatrix{4, DIRECTED}
	testAdjMatrixDirectedDFT.Init()
	testAdjMatrixDirectedDFT.AddEdge(2, 1)
	testAdjMatrixDirectedDFT.AddEdge(3, 2)
	testAdjMatrixDirectedDFT.AddEdge(2, 0)
	testAdjMatrixDirectedDFT.AddEdge(1, 0)
	testAdjMatrixDirectedDFT.AddEdge(1, 1)
	testAdjMatrixDirectedDFT.AddEdge(2, 3)
	t.Logf("Graph after DFT Directed Matrix %v ", DFT(testAdjMatrixDirectedDFT))
}

func Test_DFTAdjacencyListDirected(t *testing.T) {
	testAdjListDirectedDFT = AdjacencyList{4, DIRECTED}
	testAdjListDirectedDFT.Init()
	testAdjListDirectedDFT.AddEdge(2, 1)
	testAdjListDirectedDFT.AddEdge(3, 2)
	testAdjListDirectedDFT.AddEdge(2, 0)
	testAdjListDirectedDFT.AddEdge(1, 0)
	testAdjListDirectedDFT.AddEdge(1, 1)
	testAdjListDirectedDFT.AddEdge(2, 3)
	t.Logf("Graph after DFT Directed List %v ", DFT(testAdjListDirectedDFT))
}
