package graphs

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graphs interface {
	GetAdjacentVerticesNodesForVertex(vertex int) []int

	GetWeightOfEdge(vertexOne int, vertexTwo int) int

	GetNumberOfVertices() int

	GetIndegreeForVertex(vertex int) int

	GetOutdegreeForVertex(vertex int) int

	AddEdge(vertexOne int, vertexTwo int)

	AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int)

	RemoveEdge(vertexOne int, vertexTwo int)

	HasEdge(vertexOne int, vertexTwo int) bool

	GetGraphType() GraphType
}
