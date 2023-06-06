package graph

import "testing"

func TestGraph_AddBidirectionalEdge(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex(0)
	graph.AddVertex(0)
	graph.AddVertex(0)

	tests := []struct {
		name       string
		srcNode    int
		destNode   int
		weight     float64
		srcDegree  int
		destDegree int
		odds       int
	}{
		{
			name:       "add edge bidirectional 0 <---> 1",
			srcNode:    0,
			destNode:   1,
			weight:     1,
			srcDegree:  1,
			destDegree: 1,
			odds:       2,
		},
		{
			name:       "add edge bidirectional 1 <---> 2",
			srcNode:    1,
			destNode:   2,
			weight:     1,
			srcDegree:  2,
			destDegree: 1,
			odds:       2,
		},
		{
			name:       "add edge bidirectional 2 <---> 0",
			srcNode:    2,
			destNode:   0,
			weight:     1,
			srcDegree:  2,
			destDegree: 2,
			odds:       0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph.AddBidirectionalEdge(tt.srcNode, tt.destNode, tt.weight)
			if len(graph.OutGoingEdges[tt.srcNode]) != tt.srcDegree {
				t.Fatal("Degree of node", tt.srcNode, "should be ", tt.srcDegree)
			}
			if len(graph.OutGoingEdges[tt.destNode]) != tt.destDegree {
				t.Fatal("Degree of node", tt.destNode, "should be ", tt.destDegree)
			}
			if len(graph.OddVertices) != tt.odds {
				t.Fatal("Odds degree of grahp should be ", tt.odds, " but ", len(graph.OddVertices))
			}
		})
	}
}

func TestGraph_AddVertex(t *testing.T) {
	graph := NewGraph()

	tests := []struct {
		name string
		len  int
	}{
		{
			name: "add node 0",
			len:  1,
		},
		{
			name: "add node 1",
			len:  2,
		},
		{
			name: "add node 2",
			len:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph.AddVertex(0)
			if len(graph.Vertices) != tt.len {
				t.Fatal("Number of nodes should be ", tt.len)
			}
		})
	}
}

func TestGraph_HasEulerianCycle(t *testing.T) {

	// graph1
	//				0
	//			 /     \
	//          1 - - - 2
	graph1 := NewGraph()
	graph1.AddVertex(0)
	graph1.AddVertex(0)
	graph1.AddVertex(0)
	graph1.AddBidirectionalEdge(0, 1, 1)
	graph1.AddBidirectionalEdge(1, 2, 1)
	graph1.AddBidirectionalEdge(2, 0, 1)

	// graph2
	//				0
	//			 /     \
	//          1 - - - 2 --- 3
	graph2 := NewGraph()
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddBidirectionalEdge(0, 1, 1)
	graph2.AddBidirectionalEdge(1, 2, 1)
	graph2.AddBidirectionalEdge(2, 0, 1)
	graph2.AddBidirectionalEdge(2, 3, 1)
	tests := []struct {
		name  string
		graph Graph
		want  bool
	}{
		{
			name:  "graph1 has a Eulerian cycle",
			graph: graph1,
			want:  true,
		},
		{
			name:  "graph2 has not a Eulerian cycle",
			graph: graph2,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.graph.HasEulerianCycle(); got != tt.want {
				t.Errorf("HasEulerianCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_HasEulerianPath(t *testing.T) {

	// graph1
	//				0
	//			 /     \
	//          1 - - - 2
	graph1 := NewGraph()
	graph1.AddVertex(0)
	graph1.AddVertex(0)
	graph1.AddVertex(0)
	graph1.AddBidirectionalEdge(0, 1, 1)
	graph1.AddBidirectionalEdge(1, 2, 1)
	graph1.AddBidirectionalEdge(2, 0, 1)

	// graph2
	//				0
	//			 /     \
	//          1 - - - 2 --- 3
	graph2 := NewGraph()
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddBidirectionalEdge(0, 1, 1)
	graph2.AddBidirectionalEdge(1, 2, 1)
	graph2.AddBidirectionalEdge(2, 0, 1)
	graph2.AddBidirectionalEdge(2, 3, 1)

	// graph3
	//			0  - -  1  --  2
	//          |    /  |  \   |
	//			|  /    |    \ |
	//          3 - - - 4  - - 5
	graph3 := NewGraph()
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddBidirectionalEdge(0, 1, 1)
	graph3.AddBidirectionalEdge(0, 3, 1)
	graph3.AddBidirectionalEdge(1, 3, 1)
	graph3.AddBidirectionalEdge(1, 2, 1)
	graph3.AddBidirectionalEdge(1, 4, 1)
	graph3.AddBidirectionalEdge(1, 5, 1)
	graph3.AddBidirectionalEdge(2, 5, 1)
	graph3.AddBidirectionalEdge(3, 4, 1)
	graph3.AddBidirectionalEdge(4, 5, 1)

	tests := []struct {
		name  string
		graph Graph
		want  bool
	}{
		{
			name:  "graph1 has a Eulerian path",
			graph: graph1,
			want:  true,
		},
		{
			name:  "graph2 has a Eulerian path",
			graph: graph2,
			want:  true,
		},
		{
			name:  "graph3 has not a Eulerian path",
			graph: graph3,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.graph.HasEulerianPath(); got != tt.want {
				t.Errorf("HasEulerianPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_IsComplete(t *testing.T) {

	// graph1
	//				0
	//			 /     \
	//          1 - - - 2
	graph1 := NewGraph()
	graph1.AddVertex(0)
	graph1.AddVertex(0)
	graph1.AddVertex(0)
	graph1.AddBidirectionalEdge(0, 1, 1)
	graph1.AddBidirectionalEdge(1, 2, 1)
	graph1.AddBidirectionalEdge(2, 0, 1)

	// graph2
	//				0
	//			 /     \
	//          1 - - - 2 --- 3
	graph2 := NewGraph()
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddVertex(0)
	graph2.AddBidirectionalEdge(0, 1, 1)
	graph2.AddBidirectionalEdge(1, 2, 1)
	graph2.AddBidirectionalEdge(2, 0, 1)
	graph2.AddBidirectionalEdge(2, 3, 1)

	// graph3
	//			0 - - 3
	//			|  x  |
	//          1 - - 2
	graph3 := NewGraph()
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddVertex(0)
	graph3.AddBidirectionalEdge(0, 1, 1)
	graph3.AddBidirectionalEdge(0, 2, 1)
	graph3.AddBidirectionalEdge(0, 3, 1)
	graph3.AddBidirectionalEdge(1, 2, 1)
	graph3.AddBidirectionalEdge(1, 3, 1)
	graph3.AddBidirectionalEdge(2, 3, 1)
	tests := []struct {
		name  string
		graph Graph
		want  bool
	}{
		{
			name:  "graph1 is complete",
			graph: graph1,
			want:  true,
		},
		{
			name:  "graph2 not is complete",
			graph: graph2,
			want:  false,
		},
		{
			name:  "graph3 is complete",
			graph: graph3,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.graph.IsComplete(); got != tt.want {
				t.Errorf("HasEulerianCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_IsBipartite(t *testing.T) {

	// graph1
	//			   0(A)
	//			 /     \
	//          1(B)     2(B)
	graph1 := NewGraph()
	graph1.AddVertex(A)
	graph1.AddVertex(B)
	graph1.AddVertex(B)
	graph1.AddBidirectionalEdge(0, 1, 1)
	graph1.AddBidirectionalEdge(0, 2, 1)

	// graph2
	//			   0(A)
	//			 /     \
	//          1(B)    2(B) -- 3(A)
	graph2 := NewGraph()
	graph2.AddVertex(A)
	graph2.AddVertex(B)
	graph2.AddVertex(B)
	graph2.AddVertex(A)
	graph2.AddBidirectionalEdge(0, 1, 1)
	graph2.AddBidirectionalEdge(0, 2, 1)
	graph2.AddBidirectionalEdge(2, 3, 1)

	// graph3
	//			0(A) - 3(B)
	//			|   x  |
	//          1(B) - 2(B)
	graph3 := NewGraph()
	graph3.AddVertex(A)
	graph3.AddVertex(B)
	graph3.AddVertex(B)
	graph3.AddVertex(B)
	graph3.AddBidirectionalEdge(0, 1, 1)
	graph3.AddBidirectionalEdge(0, 2, 1)
	graph3.AddBidirectionalEdge(0, 3, 1)
	graph3.AddBidirectionalEdge(1, 2, 1)
	graph3.AddBidirectionalEdge(1, 3, 1)
	graph3.AddBidirectionalEdge(2, 3, 1)
	tests := []struct {
		name  string
		graph Graph
		want  bool
	}{
		{
			name:  "graph1 is bipartite",
			graph: graph1,
			want:  true,
		},
		{
			name:  "graph2 is bipartite",
			graph: graph2,
			want:  true,
		},
		{
			name:  "graph3 is not bipartite",
			graph: graph3,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.graph.IsBipartite(); got != tt.want {
				t.Errorf("IsBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
