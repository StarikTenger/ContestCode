package main
import (
	"os"
	"bufio"
	"fmt"
)

var reader *bufio.Reader

func solution() {
	price_amount := make(map[int]int)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	//fmt.Printf("n = %d\n", n)

	for i := 0; i < n; i++ {
		var price int
		fmt.Fscanf(reader, "%d", &price)
		price_amount[price] ++
		//fmt.Printf("%d ", price)
	}
	fmt.Fscanf(reader, "\n")

	sum := 0
	for price, amount := range price_amount {
		sum += price * ((amount / 3) * 2 + amount % 3)
	}

	//fmt.Println(price_amount)
	fmt.Printf("%d\n", sum)
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		solution()
	}
}