package main

import "fmt"

func buildFreqTable(data string) map[string]int {
	table := make(map[string]int)

	for _, char := range data {
		table[string(char)]++
	}

	return table
}

func buildArrayOfNodes(table map[string]int) []*heapNode {
	a := []*heapNode{}

	for k, v := range table {
		a = append(a, &heapNode{data: k, freq: v, left: nil, right: nil})
	}

	return a
}

func buildHuffmanTree(data string) *heapNode {
	table := buildFreqTable(data)
	nodes := buildArrayOfNodes(table)
	pqueue := buildMinHeap(nodes)

	for pqueue.size > 1 {
		first := popMinHeap(pqueue)
		second := popMinHeap(pqueue)

		newNode := &heapNode{data: "", freq: first.freq + second.freq, right: second, left: first}

		pushMinHeap(pqueue, newNode)
	}

	return popMinHeap(pqueue)
}

func buildCodesTable(treeRoot *heapNode) map[string]string {
	codesTable := make(map[string]string)
	buildCode(treeRoot, "", &codesTable)
	return codesTable
}

func buildCode(node *heapNode, code string, table *map[string]string) {
	if node.left != nil {
		buildCode(node.left, code+"0", table)
	}

	if node.right != nil {
		buildCode(node.right, code+"1", table)
	}

	if node.left == nil && node.right == nil {
		(*table)[node.data] = code
		return
	}
}

func printCodesTable(table map[string]string) {
	for k, v := range table {
		fmt.Println(k + ":" + v)
	}
}

func main() {
	root := buildHuffmanTree("eababacacd")
	table := buildCodesTable(root)
	printCodesTable(table)
	fmt.Println("=======================")
	root = buildHuffmanTree("ffcabracadabrarrarradff")
	table = buildCodesTable(root)
	printCodesTable(table)
	fmt.Println("=======================")
	root = buildHuffmanTree("aaaaabbbbbbbbbccccccccccccdddddddddddddeeeeeeeeeeeeeeeefffffffffffffffffffffffffffffffffffffffffffff")
	table = buildCodesTable(root)
	printCodesTable(table)
	fmt.Println("=======================")
}
