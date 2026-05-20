package exercises

import (
	"fmt"
)

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
	// TODO
	if a.TxID != b.TxID && (a.Type == Write || b.Type == Write) && a.Item == b.Item {
		return true
	}
	return false
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

	// TODO:
	// For every i < j:
	// if schedule[i] conflicts with schedule[j],
	// add edge schedule[i].TxID -> schedule[j].TxID

	for i := 0; i < len(schedule); i++ {
		for j := i + 1; j < len(schedule); j++ {
			if conflicts(schedule[i], schedule[j]) {
				addEdge(graph, schedule[i].TxID, schedule[j].TxID)
			}
		}
	}

	return graph
}

func hasCycle(graph map[int]map[int]bool) bool {
	visited := make(map[int]bool)
	inStack := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(node int) bool {
		// TODO:
		// standard DFS cycle detection in directed graph
		visited[node] = true
		inStack[node] = true

		neighbours := graph[node]
		for nei, _ := range neighbours {
			if !visited[nei] {
				if dfs(nei) {
					return true
				}
			} else if inStack[nei] {
				return true
			}
		}

		inStack[node] = false
		return false
	}

	// Need to visit all nodes appearing either as source or destination
	nodes := make(map[int]bool)
	for from, nbrs := range graph {
		nodes[from] = true
		for to := range nbrs {
			nodes[to] = true
		}
	}

	for node := range nodes {
		if !visited[node] {
			if dfs(node) {
				return true
			}
		}
	}

	return false
}

func printGraph(graph map[int]map[int]bool) {
	fmt.Println("Precedence Graph:")
	for from, nbrs := range graph {
		for to := range nbrs {
			fmt.Printf("T%d -> T%d\n", from, to)
		}
	}
}

func main() {
	schedule := []Operation{
		{TxID: 1, Type: Read, Item: "A"},
		{TxID: 2, Type: Read, Item: "A"},
		{TxID: 1, Type: Write, Item: "A"},
		{TxID: 2, Type: Write, Item: "A"},
	}

	graph := buildPrecedenceGraph(schedule)
	printGraph(graph)

	if hasCycle(graph) {
		fmt.Println("Schedule is NOT conflict serializable")
	} else {
		fmt.Println("Schedule IS conflict serializable")
	}
}
