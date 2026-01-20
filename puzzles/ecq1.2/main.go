package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/PhilAldridge/aoc-2024-go/pkg/files"
	"github.com/PhilAldridge/aoc-2024-go/pkg/ints"
)

func main() {
	start := time.Now()
	fmt.Println("Part 1 answer: ", part1("input.txt"))
	split := time.Now()
	fmt.Println("Part 2 answer: ", part2("input2.txt"))
	split2 := time.Now()
	fmt.Println("Part 3 answer: ", part3("input3.txt"))

	fmt.Println()
	fmt.Println("Part 1: ", split.Sub(start))
	fmt.Println("Part 2: ", split2.Sub(split))
	fmt.Println("Part 2: ", time.Since(split2))
}

func part1(name string) string {
	input:= files.ReadLines(name)

	var leftTree, rightTree *treeNode

	for i,row:= range input {
		nodes, swap:= parseRow(row)

		if i == 0 {
			leftTree = &nodes[0]
			rightTree = &nodes[1]
			continue
		}

		if swap !=0 {
			leftTree,rightTree = swapNodes(leftTree, rightTree, int(swap),true)
			continue
		}

		placeNewNode(nodes[0],leftTree)
		placeNewNode(nodes[1],rightTree)
	}

	return parseTreeString(leftTree) + parseTreeString(rightTree)
}

func part2(name string) string {
	return part1(name)
}

func part3(name string) string {
	input:= files.ReadLines(name)

	var leftTree, rightTree *treeNode

	for i,row:= range input {
		nodes, swap:= parseRow(row)

		if i == 0 {
			leftTree = &nodes[0]
			rightTree = &nodes[1]
			continue
		}

		if swap !=0 {
			leftTree,rightTree = swapNodes(leftTree, rightTree, int(swap),false)
			continue
		}

		placeNewNode(nodes[0],leftTree)
		placeNewNode(nodes[1],rightTree)
	}

	return parseTreeString(leftTree) + parseTreeString(rightTree)
}

type treeNode struct {
	id     int
	symbol string
	rank   int
	left   *treeNode
	right  *treeNode
}

type swapType int

func placeNewNode(node treeNode, tree *treeNode) {
	if node.rank < tree.rank {
		if tree.left == nil {
			tree.left = &node
		} else {
			placeNewNode(node, tree.left)
		}
		return
	}

	if node.rank > tree.rank {
		if tree.right == nil {
			tree.right = &node
		} else {
			placeNewNode(node, tree.right)
		}
	}
}

func parseRow(row string) ([2]treeNode, swapType) {
	var nodes [2]treeNode

	split := strings.Split(row, " ")
	if len(split) == 2 {
		return nodes, swapType(ints.FromString(split[1]))
	}

	id := ints.FromString(split[1][3:])
	leftStr := strings.Split(split[2][6:len(split[2])-1], ",")
	rightStr := strings.Split(split[3][7:len(split[3])-1], ",")

	nodes[0] =  treeNode{
		id:     id,
		symbol: leftStr[1],
		rank:   ints.FromString(leftStr[0]),
	}

	nodes[1] = treeNode{
		id:     id,
		symbol: rightStr[1],
		rank:   ints.FromString(rightStr[0]),
	}

	return nodes, 0
}

func parseTreeString(tree *treeNode) string {
	levels := countLevels(tree)
	maxLevel := slices.Index(levels, ints.Max(levels))

	return getString(tree, maxLevel)
}

func countLevels(tree *treeNode) []int {
	out := []int{0}
	var leftPath, rightPath []int

	if tree.left != nil {
		out[0]++
		leftPath = countLevels(tree.left)
	}

	if tree.right != nil {
		out[0]++
		rightPath = countLevels(tree.right)
	}

	combinedPaths := make([]int, ints.Max([]int{len(leftPath), len(rightPath)}))

	for i, val := range leftPath {
		combinedPaths[i] += val
	}

	for i, val := range rightPath {
		combinedPaths[i] += val
	}

	out = append(out, combinedPaths...)

	return out
}

func getString(tree *treeNode, level int) string {
	var out string
	if level == 0 {
		if tree.left != nil {
			out += tree.left.symbol
		}

		if tree.right != nil {
			out += tree.right.symbol
		}

		return out
	}

	if tree.left != nil {
		out += getString(tree.left, level-1)
	}

	if tree.right != nil {
		out += getString(tree.right, level-1)
	}

	return out
}

func swapNodes(a,b *treeNode, id int, swapBranches bool) (*treeNode, *treeNode) {
	exclusions:= []string{}
	firstNode,ok,exclusion:= findNode(a,id,exclusions)
	if !ok {
		firstNode,_,exclusion= findNode(b,id,exclusions)
	}
	exclusions = append(exclusions, exclusion)

	secondNode,ok,_:=findNode(b,id,exclusions)
	if !ok {
		secondNode,_,_ = findNode(a,id,exclusions)
	}

	if swapBranches {
		firstNode.left, firstNode.right, secondNode.left, secondNode.right = secondNode.left, secondNode.right, firstNode.left, firstNode.right
	}

	if a.id == id {
		return b,a
	}

	dummyNode:= &treeNode{}

	a.SwapRefs(firstNode,dummyNode)
	b.SwapRefs(firstNode,dummyNode)
	a.SwapRefs(secondNode,firstNode)
	b.SwapRefs(secondNode,firstNode)
	a.SwapRefs(dummyNode,secondNode)
	b.SwapRefs(dummyNode,secondNode)

	return a,b
}

func findNode(a *treeNode, id int, exclusions []string) (*treeNode, bool, string) {
	if a == nil {
		return a, false, ""
	}

	if a.id == id && !slices.Contains(exclusions,a.symbol) {
		return a,true, a.symbol
	}
		
	node,ok, str:= findNode(a.left,id,exclusions)
	if ok {
		return node,true, str
	}

	node,ok,str = findNode(a.right,id,exclusions)
	if ok {
		return node,true,str
	}
	
	return a,false,""
}

func (a *treeNode) SwapRefs (b,c *treeNode) {
	if a == nil {
		return
	}

	if a.left != nil && a.left == b {
		a.left = c
		return
	}

	if a.right != nil && a.right == b {
		a.right = c
		return
	}

	a.left.SwapRefs(b,c)
	a.right.SwapRefs(b,c)
}