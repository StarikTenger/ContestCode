package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

func solution() {
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	finished_tasks := make(map[int]bool)

	var task_prev int
	fmt.Fscanf(reader, "%d", &task_prev)

	ans := true

	for i := 1; i < n; i++ {
		var task int
		fmt.Fscanf(reader, "%d", &task)

		if finished_tasks[task] {
			ans = false
		}

		if task != task_prev {
			finished_tasks[task_prev] = true
		}

		task_prev = task
	}
	fmt.Fscanf(reader, "\n")

	if ans {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
	
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solution()
	}
}