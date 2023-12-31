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
func (g *Graph) AddNode(n *Node) *Node {
	node, exists := g.NodeMap[n.Name]
	if !exists {
		g.Nodes = append(g.Nodes, n)
		g.NodeMap[n.Name] = n
		return n
	}
	return node
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

// TopologicalSort performs a topological sort on the graph
func (g *Graph) TopologicalSort() []*Node {
	var result []*Node
	visited := make(map[*Node]bool)
	var visit func(u *Node)

	visit = func(u *Node) {
		if !visited[u] {
			visited[u] = true
			for _, edge := range g.Edges[u.Name] {
				visit(edge.Node)
			}
			result = append([]*Node{u}, result...)
		}
	}

	for _, node := range g.Nodes {
		visit(node)
	}

	return result
}

// LongestPath finds the longest path between two nodes in the graph
func (g *Graph) LongestPath(src, dest *Node) int {
	topOrder := g.TopologicalSort()
	dist := make(map[*Node]int)

	for _, node := range g.Nodes {
		dist[node] = int(^uint(0) >> 1) // Initialize as negative infinity
	}
	dist[src] = 0

	for _, node := range topOrder {
		if dist[node] != int(^uint(0)>>1) {
			for _, edge := range g.Edges[node.Name] {
				if dist[edge.Node] < dist[node]+edge.Weight {
					dist[edge.Node] = dist[node] + edge.Weight
				}
			}
		}
	}

	return dist[dest]
}
