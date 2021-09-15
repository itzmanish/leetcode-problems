package leetcodeplayground

// test case- [[2,4],[1,3],[2,4],[1,3]]

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	resultMap := make(map[*Node]*Node)
	visitedMap := make(map[*Node]bool)

	resultMap[node] = &Node{Val: node.Val}
	visitedMap[node] = true
	deepCopy(node, resultMap, visitedMap)
	return resultMap[node]
}

func deepCopy(src *Node, res map[*Node]*Node, visited map[*Node]bool) {
	for _, adjNode := range src.Neighbors {
		if _, ok := res[adjNode]; !ok {
			res[adjNode] = &Node{Val: adjNode.Val}
		}

		res[src].Neighbors = append(res[src].Neighbors, res[adjNode])

		if !visited[adjNode] {
			visited[adjNode] = true
			deepCopy(adjNode, res, visited)
		}
	}
}
