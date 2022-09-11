package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

type Dev struct {
	level int
	blocked bool
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solution() {
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	devs := make([]Dev, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &devs[i].level)
		devs[i].blocked = false
	}
	fmt.Fscanf(reader, "\n")

	for i := 0; i < n; i++ {
		if devs[i].blocked {
			continue
		}

		devs[i].blocked = true
		j_best := i + 1
		diff_min := 100000
		for j := i + 1; j < n; j++ {
			if devs[j].blocked {
				continue
			}

			diff := abs(devs[i].level - devs[j].level)
			if diff < diff_min {
				j_best = j
				diff_min = diff				
			}
		}
		devs[j_best].blocked = true
		fmt.Printf("%d %d\n", i + 1, j_best + 1)
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