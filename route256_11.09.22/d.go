package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

func check(seq *[]uint64, l int, r int) bool {
	var slot1 uint64
	var slot2 uint64
	slot1, slot2 = 0, 0
	for i := l; i < r; i++ {
		if slot1 == 0 {
			slot1 = (*seq)[i]
		} else if slot2 == 0 && (*seq)[i] != slot1 {
			slot2 = (*seq)[i]
		}
		if (*seq)[i] != slot1 && (*seq)[i] != slot2 {
			return false
		}
	}
	return true
}

func solution() {
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	
	seq := make([]uint64, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &seq[i])
	}
	fmt.Fscanf(reader, "\n")

	val_max := 0
	//l_max := 0
	//r_max := 0
	for l, r := 0, 1; l < r && r <= n; {
		if check(&seq, l, r) {
			if r - l > val_max {
				val_max=  r - l
				//l_max = l
				//r_max = r
			}
			r++
		} else {
			l++
		}
	}

	fmt.Printf("%d\n", val_max)
	/*for i := l_max; i < r_max; i++ {
		fmt.Printf("%d ", seq[i])
	}
	fmt.Printf("\n")*/
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solution()
	}
}