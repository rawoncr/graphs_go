package graph

type Graph struct {
	Vertices      []Node
	OutGoingEdges [][]Edge
	OddVertices   map[int]bool
	Bipartite     bool
}

type Edge struct {
	From, To int
	Weight   float64
}

type Node struct {
	ID   int
	Type Type
}

type Type int

const (
	A Type = iota
	B
)

func NewGraph() Graph {
	return Graph{
		Vertices:    []Node{},
		OddVertices: map[int]bool{},
		Bipartite:   true,
	}
}

func (graph *Graph) AddVertex(Type Type) {
	n := Node{ID: len(graph.Vertices), Type: Type}
	graph.Vertices = append(graph.Vertices, n)
	graph.OutGoingEdges = append(graph.OutGoingEdges, []Edge{})
}

func (graph *Graph) AddBidirectionalEdge(srcKey, destKey int, weight float64) {
	// add edge src --> dest
	graph.OutGoingEdges[srcKey] = append(graph.OutGoingEdges[srcKey], Edge{Weight: weight, From: srcKey, To: destKey})
	// add edge dest --> src
	graph.OutGoingEdges[destKey] = append(graph.OutGoingEdges[destKey], Edge{Weight: weight, From: destKey, To: srcKey})

	graph.updateOddVertices(srcKey, destKey)
	graph.updateBipartite(srcKey, destKey)
}

func (graph *Graph) updateOddVertices(ids ...int) {
	for _, id := range ids {
		// if number of edges is pair, then should be removed from OddVertices
		if len(graph.OutGoingEdges[id])%2 == 0 {
			delete(graph.OddVertices, id)
		} else { // if number of edges is odd, add to OddVertices
			graph.OddVertices[id] = true
		}
	}
}

func (graph *Graph) updateBipartite(srcKey, destKey int) {
	if graph.Vertices[srcKey].Type == graph.Vertices[destKey].Type {
		graph.Bipartite = false
	}
}

func (graph Graph) HasEulerianPath() bool {
	odds := len(graph.OddVertices)
	return odds == 0 || odds == 2
}

func (graph Graph) HasEulerianCycle() bool {
	odds := len(graph.OddVertices)
	return odds == 0
}

func (graph Graph) IsComplete() bool {
	vertices := len(graph.Vertices)
	for _, vertex := range graph.OutGoingEdges {
		if len(vertex) != vertices-1 {
			return false
		}
	}
	return true
}

func (graph Graph) IsBipartite() bool {
	return graph.Bipartite
}
