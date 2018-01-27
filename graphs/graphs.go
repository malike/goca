package graphs

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graphs interface {
	Init()

	AddEdge(vertexOne int, vertexTwo int) error

	AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error

	RemoveEdge(vertexOne int, vertexTwo int) error

	HasEdge(vertexOne int, vertexTwo int) bool

	GetGraphType() GraphType

	GetAdjacentNodesForVertex(vertex int) map[int]bool

	GetOutgoingNodesForVertex(vertex int) map[int]bool

	GetWeightOfEdge(vertexOne int, vertexTwo int) (int,error)

	GetNumberOfVertices() int

	GetIndegreeForVertex(vertex int) int
}
