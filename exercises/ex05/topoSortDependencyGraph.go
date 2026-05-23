package exercises

import "fmt"

type OpType string

const (
	Read  OpType = "R"
	Write OpType = "W"
)

type Operation struct {
	TxID int
	Type OpType
	Item string
}

func conflicts(a, b Operation) bool {
	return a.TxID != b.TxID &&
		a.Item == b.Item &&
		(a.Type == Write || b.Type == Write)
}

func addEdge(graph map[int]map[int]bool, from, to int) {
	if from == to {
		return
	}
	if _, ok := graph[from]; !ok {
		graph[from] = make(map[int]bool)
	}
	graph[from][to] = true
}

func buildPrecedenceGraph(schedule []Operation) map[int]map[int]bool {
	graph := make(map[int]map[int]bool)

	for i := 0; i < len(schedule); i++ {
		for j := i + 1; j < len(schedule); j++ {
			if conflicts(schedule[i], schedule[j]) {
				addEdge(graph, schedule[i].TxID, schedule[j].TxID)
			}
		}
	}

	return graph
}

func collectNodes(graph map[int]map[int]bool) map[int]bool {
	nodes := make(map[int]bool)
	for from, nbrs := range graph {
		nodes[from] = true
		for to := range nbrs {
			nodes[to] = true
		}
	}
	return nodes
}

func topoSort(graph map[int]map[int]bool) ([]int, bool) {
	// TODO:
	// 1. collect all nodes
	nodes := collectNodes(graph)
	indegrees := make(map[int]int)
	q := []int{}
	res := []int{}

	// 2. compute indegree of each node
	for node := range nodes {
		indegrees[node] = 0
		for rcvr := range graph[node] {
			indegrees[rcvr] += 1
		}
	}

	for node, indegree := range indegrees {
		if indegree == 0 {
			q = append(q, node)
			delete(indegrees, node)
		}
	}

	// 3. initialize queue with indegree-0 nodes
	// 4. run Kahn's algorithm
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res = append(res, curr)

		for nei := range graph[curr] {
			indegrees[nei] -= 1
			if indegrees[nei] == 0 {
				q = append(q, nei)
			}
		}
	}
	// 5. if result size != number of nodes, return false
	if len(res) == len(nodes) {
		return res, true
	}

	return nil, false
}

func main() {
	schedule := []Operation{
		{TxID: 1, Type: Read, Item: "A"},
		{TxID: 1, Type: Write, Item: "A"},
		{TxID: 2, Type: Read, Item: "A"},
		{TxID: 2, Type: Write, Item: "A"},
	}

	graph := buildPrecedenceGraph(schedule)

	order, ok := topoSort(graph)
	if !ok {
		fmt.Println("Schedule is NOT conflict serializable")
		return
	}

	fmt.Println("Schedule IS conflict serializable")
	fmt.Println("One equivalent serial order:", order)
}
