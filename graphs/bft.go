package graphs

import "fmt"

func BFT(graph Graphs) []int {

	visited := map[int]bool{}
	bftArrayOrdered := make([]int, graph.GetNumberOfVertices())
	for i := 0; i < graph.GetNumberOfVertices(); i++ {
		bftArrayOrdered = bft(graph, visited, i, bftArrayOrdered)
	}
	return bftArrayOrdered
}

func bft(graphs Graphs, visited map[int]bool, node int, bftArrayOrdered []int) []int {
	queue := make([]int, 0)
	queue = append(queue, node)
	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		fmt.Printf("%d->", current)
		bftArrayOrdered = append(bftArrayOrdered, current)
		visited[current] = true

		adjNodes := graphs.GetAdjacentNodesForVertex(current)
		for key, _ := range adjNodes {
			queue = append(queue, key)
		}

	}
	return bftArrayOrdered
}
