/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := make(map[int]*Node)
    visited[1] = &Node{Val: node.Val}
    var dfs func(node *Node)
    dfs = func(node *Node) {
        cpyNode := visited[node.Val]
        for _, n := range node.Neighbors {
            cpy := visited[n.Val]
            if cpy == nil {
                cpy = &Node{Val: n.Val}
                visited[n.Val] = cpy
                dfs(n)
            }
            cpyNode.Neighbors = append(cpyNode.Neighbors, cpy)
        }
    }
    dfs(node)
    return visited[1]
}
