package graphs

import "fmt"

func DFT(graph Graphs) []int {
	visited := map[int]bool{}
	dftArrayOrdered := make([]int, graph.GetNumberOfVertices())
	for i := 0; i < graph.GetNumberOfVertices(); i++ {
		dft(graph, visited, i, dftArrayOrdered)
	}
	return dftArrayOrdered
}

func dft(graphs Graphs, visited map[int]bool, current int, dftArrayOrdered []int) {
	if visited[current] {
		return
	}
	visited[current] = true
	adjNodes := graphs.GetAdjacentNodesForVertex(current)
	for key, _ := range adjNodes {
		dft(graphs, visited, key, dftArrayOrdered)
	}
	dftArrayOrdered = append(dftArrayOrdered, current)
	fmt.Printf("%d->", current)

}
