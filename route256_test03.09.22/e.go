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


	for i := 1; i < n; i++ {
		
	}
	fmt.Fscanf(reader, "\n")	
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solution()
	}
}