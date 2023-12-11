package util

func BFS(g *Graph, startName string) map[string]bool {
	queue := []*Node{g.GetNode(startName)}
	visited := make(map[string]bool)
	parent := make(map[string]string)

	visited[startName] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, edge := range g.Edges[node.Name] {
			if !visited[edge.Node.Name] {
				visited[edge.Node.Name] = true
				parent[edge.Node.Name] = node.Name
				queue = append(queue, edge.Node)
			}
		}
	}
	return visited
}

func constructPath(parent map[string]string, start, end string) []string {
	path := []string{}
	for at := end; at != start; at = parent[at] {
		path = append([]string{at}, path...)
	}
	path = append([]string{start}, path...)
	return path
}
