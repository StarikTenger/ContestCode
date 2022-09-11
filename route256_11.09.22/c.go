package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

func pattern_match(str *string, pos int, pattern string) int {
	l := len(*str)
	if pos + len(pattern) > l {
		return 0
	}
	for i := 0; i < len(pattern); i++ {
		if pattern[i] != (*str)[pos + i] {
			return 0
		}
	}
	return len(pattern)
}

func solution() {
	var str string
	fmt.Fscanf(reader, "%s\n", &str)

	for i := 0; i < len(str);  {
		if delta := pattern_match(&str, i, "00"); delta != 0 {
			fmt.Printf("a");
			i += delta
		}
		if delta := pattern_match(&str, i, "100"); delta != 0 {
			fmt.Printf("b");
			i += delta
		}
		if delta := pattern_match(&str, i, "101"); delta != 0 {
			fmt.Printf("c");
			i += delta
		}
		if delta := pattern_match(&str, i, "11"); delta  != 0{
			fmt.Printf("d");
			i += delta
		}
		
	}
	fmt.Printf("\n");
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solution()
	}
}