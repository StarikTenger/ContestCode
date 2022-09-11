package main
import (
	"os"
	"bufio"
	"fmt"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

var reader *bufio.Reader

type Node struct {
	children []*Node
	words_below int
	final bool
}

func make_node() *Node{
	var node Node
	node.children = make([]*Node, 'z'-'a' + 1)
	node.final = false
	node.words_below = 0
	return &node
}

func add_word(node *Node, word string, n int) {
	node.words_below ++

	if n >= len(word) {
		node.final = true
		return
	}

	index := word[n] - 'a'

	if node.children[index] == nil {
		node.children[index] = make_node()
	}

	add_word(node.children[index], word, n + 1);
}

func find_word(node *Node, prefix string, n int, acc string) string {
	
	//fmt.Println(n, acc)
	if n < len(prefix) && node.children[prefix[n] - 'a'] != nil {
		//fmt.Println(string(rune(prefix[n])))
		word := find_word(node.children[prefix[n] - 'a'], prefix, n + 1, acc + string(rune(prefix[n])))
		if word != "" {
			return word
		}
	} else if  node.final && prefix != acc{
		return acc
	}

	for i := 0; i <= 'z' - 'a'; i++ {
		if node.children[i] == nil {
			continue
		}
		word := find_word(node.children[i], prefix, n + 1, acc + string(rune(i + 'a')))
		//fmt.Println(string(rune(prefix[n])))
		if word != "" {
			return word
		}
	}
	return ""
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	tree := make_node()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscanf(reader, "%s\n", &word)
		add_word(tree, Reverse(word), 0)
		//fmt.Println(word)
	}


	var m int
	fmt.Fscanf(reader, "%d\n", &m)
	for i := 0; i < m; i++ {
		var word string
		fmt.Fscanf(reader, "%s\n", &word)
		fmt.Printf("%s\n", Reverse(find_word(tree, Reverse(word), 0, "")))
	}
}