package main
import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		fmt.Printf("%d\n", a + b)
	}
}