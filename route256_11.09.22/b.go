package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

func solution() {
	var d, m, y int
	fmt.Fscanf(reader, "%d %d %d\n", &d, &m, &y)
	days_in_mounth := [12]int{
		31, // jan
		28, // feb
		31, // mar
		30, // apr
		31, // may
		30, // jun
		31, // jul
		31, // aug
		30, // sep
		31, // oct
		30, // nov
		31, // dec
	}
	if y % 4 == 0 && y % 100 != 0 || y % 400 == 0 {
		days_in_mounth[1] = 29
	}

	if d <= days_in_mounth[m - 1] {
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