package graphs

import (
	"testing"
)

var  testAdjMatrixDirectedBFS Graphs


func Test_BFSAdjacencyMatrix(t *testing.T) {
	testAdjMatrixDirectedBFS = AdjacencyMatrix{4, UNDIRECTED}
	testAdjMatrixDirectedBFS.Init()
	testAdjMatrixDirectedBFS.AddEdge(2, 1)
	testAdjMatrixDirectedBFS.AddEdge(3, 2)
	testAdjMatrixDirectedBFS.AddEdge(2, 0)
	testAdjMatrixDirectedBFS.AddEdge(1, 0)
	testAdjMatrixDirectedBFS.AddEdge(1, 1)
	testAdjMatrixDirectedBFS.AddEdge(2, 3)
	t.Log(BFS(testAdjMatrixDirectedBFS)+" is the type of graph ")
}

