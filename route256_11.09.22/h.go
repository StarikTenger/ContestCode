package main
import (
	"os"
	"bufio"
	"fmt"
	"sort"
)

var reader *bufio.Reader

type Vec2 struct {
	x int
	y int
}

func find_rcorner(matrix *[]string, p Vec2) Vec2 {
	for ;p.x < len(*matrix) && (*matrix)[p.x][p.y] == '*'; p.x++ {}
	p.x--
	for ;p.y < len((*matrix)[0]) && (*matrix)[p.x][p.y] == '*'; p.y++ {}
	p.y--
	return p
}

func find(matrix *[]string, l Vec2, r Vec2, nest int, visited *[][]bool, order *[]int) {

	for x := l.x + 1; x < r.x - 1; x++ {
		for y := l.y + 1; y < r.y - 1; y++ {
			if !(*visited)[x][y] &&
				(*matrix)[x][y] == '*' &&
				(*matrix)[x + 1][y] == '*' &&
				(*matrix)[x][y + 1] == '*' {
					
					p := find_rcorner(matrix, Vec2{x, y})
					//fmt.Println(Vec2{x, y}, p)
					//fmt.Printf("%d ", nest)
					(*order) = append(*order, nest)
					(*visited)[x][y] = true
					find(matrix, Vec2{x, y}, p, nest + 1, visited, order)					
				}
		}	
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

	order := make([]int, 0)

	find(&matrix, Vec2{-1, -1}, Vec2{n, m}, 0, &visited, &order)
	sort.Ints(order)

	//fmt.Println(order)
	for _, v := range order {
		fmt.Printf("%d ", v)
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