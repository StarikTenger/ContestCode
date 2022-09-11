package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

type Node struct {
	neighbor1 int
	neighbor2 int
	visited bool
}

func dfs(graph *map[int]Node, cur int, order *[]int) {
	*order = append(*order, cur)
	if entery, ok := (*graph)[cur]; ok {
		entery.visited = true
		(*graph)[cur] = entery
	}
	if !(*graph)[(*graph)[cur].neighbor1].visited {
		dfs(graph, (*graph)[cur].neighbor1, order)
	}
	if !(*graph)[(*graph)[cur].neighbor2].visited {
		dfs(graph, (*graph)[cur].neighbor2, order)
	}
}

func solution() {
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	
	graph := make(map[int]Node)

	start_v := 0
	for i := 0; i < n; i++ {
		var v, n1, n2 int
		fmt.Fscanf(reader, "%d %d %d\n", &v, &n1, &n2)
		start_v = v
		graph[v] = Node{n1, n2, false}
	}

	
	order := make([]int, 0)
	dfs(&graph, start_v, &order)

	for i := 0; i < len(order) / 2; i++ {
		fmt.Printf("%d %d\n", order[i], order[i + len(order) / 2])
	}
	fmt.Printf("\n")
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solution()
	}
}