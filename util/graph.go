package util

// Node represents a graph node
type Node struct {
	Name  string
	Value int
}

// Edge represents an edge in the graph with a weight
type Edge struct {
	Node   *Node
	Weight int
}

// Graph represents a graph
type Graph struct {
	Nodes   []*Node
	NodeMap map[string]*Node // Map for quick node lookup by name
	Edges   map[string][]*Edge
}

// AddNode adds a new node to the graph, ensuring no duplicates by name
func (g *Graph) AddNode(n *Node) {
	if _, exists := g.NodeMap[n.Name]; !exists {
		g.Nodes = append(g.Nodes, n)
		g.NodeMap[n.Name] = n
	}
}

// GetNode returns the node with the given name, using the map for O(1) lookup
func (g *Graph) GetNode(name string) *Node {
	return g.NodeMap[name]
}

// AddDirectedEdge adds a new edge to the graph
func (g *Graph) AddDirectedEdge(n1 *Node, n2 *Node, weight int) {
	g.Edges[n1.Name] = append(g.Edges[n1.Name], &Edge{n2, weight})
}

// AddEdge adds a new edge to the graph
func (g *Graph) AddUnDirectedEdge(n1 *Node, n2 *Node, weight int) {
	g.Edges[n1.Name] = append(g.Edges[n1.Name], &Edge{n2, weight})
	g.Edges[n2.Name] = append(g.Edges[n2.Name], &Edge{n1, weight})
}

// NewGraph creates and returns a new Graph instance
func NewGraph() *Graph {
	return &Graph{
		Nodes:   make([]*Node, 0),
		NodeMap: make(map[string]*Node),
		Edges:   make(map[string][]*Edge),
	}
}
