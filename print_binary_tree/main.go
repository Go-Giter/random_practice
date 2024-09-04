package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data  int
	right *Node
	left  *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(getHeight(root.left), getHeight(root.right))
}

func fill(tree *Node, ans [][]string, row, column, height int) {
	if tree == nil {
		return
	}

	ans[row][column] = strconv.Itoa(tree.data)

	d := delta(row+1, height)

	fill(tree.left, ans, row+1, column-d, height)
	fill(tree.right, ans, row+1, column+d, height)
}

func delta(row, height int) int {
	if row == height {
		return 0
	}

	return 1 << (height - row - 1)
}

func main() {
	tree := &Node{
		data: 1,
		right: &Node{
			data: 4,
			right: &Node{
				data: 5,
			},
			left: &Node{
				data: 6,
			},
		},
		left: &Node{
			data: 2,
			right: &Node{
				data: 3,
			},
		},
	}

	rows := getHeight(tree)
	cols := (1 << rows) - 1

	ans := make([][]string, rows)
	for i := range ans {
		ans[i] = make([]string, cols)
		for j := range ans[i] {
			ans[i][j] = ""
		}
	}

	fill(tree, ans, 0, (cols-1)/2, rows)

	fmt.Println(ans)
}
