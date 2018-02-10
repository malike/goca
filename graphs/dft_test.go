package graphs

import "testing"

var  testAdjMatrixDirectedDFS Graphs


func Test_DFSAdjacencyMatrix(t *testing.T) {
	testAdjMatrixDirectedDFS = AdjacencyMatrix{4, DIRECTED}
	testAdjMatrixDirectedDFS.Init()
	testAdjMatrixDirectedDFS.AddEdge(2, 1)
	testAdjMatrixDirectedDFS.AddEdge(3, 2)
	testAdjMatrixDirectedDFS.AddEdge(2, 0)
	testAdjMatrixDirectedDFS.AddEdge(1, 0)
	testAdjMatrixDirectedDFS.AddEdge(1, 1)
	testAdjMatrixDirectedDFS.AddEdge(2, 3)
	t.Log("Output "+DFS(testAdjMatrixDirectedDFS,2)+"\n")
}
