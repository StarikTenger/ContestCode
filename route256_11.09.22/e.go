package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Max returns the larger of x or y.
func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

type Vec2 struct {
	x int
	y int
}

func calc_neighbors(matrix *[]string, pos Vec2) int {
	n, m := len(*matrix), len((*matrix)[0])
	neighbors := 0
	for x := max(0, pos.x - 1); x <= min(n - 1, pos.x + 1); x++ {
		for y := max(0, pos.y - 1); y <= min(m - 1, pos.y + 1); y++ {
			if abs(x - pos.x) == abs(y - pos.y) &&  abs(y - pos.y) != 0{
				continue
			}
			if (*matrix)[x][y] == '*' {
				neighbors ++
			}
		}	
	}
	return neighbors
}

func dfs(matrix *[]string, visited *[][]bool, pos Vec2) {
	n, m := len(*matrix), len((*matrix)[0])
	(*visited)[pos.x][pos.y] = true
	var pos_new Vec2

	pos_new = pos
	pos_new.x --
	if pos_new.x >= 0 && 
	(*matrix)[pos_new.x][pos_new.y] == '*' && 
	!(*visited)[pos_new.x][pos_new.y] {
		fmt.Printf("U");
		dfs(matrix, visited, pos_new)
	}

	pos_new = pos
	pos_new.x ++
	if pos_new.x < n && 
	(*matrix)[pos_new.x][pos_new.y] == '*' && 
	!(*visited)[pos_new.x][pos_new.y] {
		fmt.Printf("D");
		dfs(matrix, visited, pos_new)
	}

	pos_new = pos
	pos_new.y --
	if pos_new.y >= 0 && 
	(*matrix)[pos_new.x][pos_new.y] == '*' && 
	!(*visited)[pos_new.x][pos_new.y] {
		fmt.Printf("L");
		dfs(matrix, visited, pos_new)
	}

	pos_new = pos
	pos_new.y ++
	if pos_new.y < m && 
	(*matrix)[pos_new.x][pos_new.y] == '*' && 
	!(*visited)[pos_new.x][pos_new.y] {
		fmt.Printf("R");
		dfs(matrix, visited, pos_new)
	}
}

func solution() {
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	
	matrix := make([]string, n)
	visited := make([][]bool, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &matrix[i])
		visited[i] = make([]bool, m)
	}
	
	// Find start
	start := Vec2{0, 0}
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if matrix[x][y] == '*' && calc_neighbors(&matrix, Vec2{x, y}) == 2 {
				start = Vec2{x, y}
			}
		}	
	}

	dfs(&matrix, &visited, start)
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