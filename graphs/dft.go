package graphs

import (
	"log"
	"fmt"
)

func DFS(graph Graphs,root int) string {
	visited := map[int]bool{}
	current := root
	dfsString := ""
	dfs(graph,visited,current,dfsString)
	return dfsString
}


func dfs(graphs Graphs,visited map[int]bool, current int ,dfsString string)  {
	if visited[current]{
		return
	}
	visited[current] = true
	adjNodes := graphs.GetAdjacentNodesForVertex(current)
	for key,_ := range adjNodes {
		dfs(graphs,visited,key,dfsString)
	}
	log.Printf("%d ->",current)
	dfsString = dfsString+fmt.Sprintf("-> %d",current)
}