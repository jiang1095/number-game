package helper

import (
	"fmt"
)

type NumberNode struct {
	ParentNum *NumberNode
	NumSet    []string
	Number    string
	NextNum   map[string]*NumberNode
}

func NewTree(numberSet []string) *NumberNode {
	return buildTree(numberSet)
}

func buildTree(numSet []string) *NumberNode {
	var node = &NumberNode{
		ParentNum: nil,
		NumSet:    numSet,
		Number:    "",
		NextNum:   nil,
	}
	if len(numSet) > 1 {
		var statSet map[string][]string
		node.Number, statSet = MaxMin(numSet)
		node.NextNum = make(map[string]*NumberNode)
		for k, v := range statSet {
			childNode := buildTree(v)
			node.NextNum[k] = childNode
			childNode.ParentNum = node
		}
	} else {
		node.Number = numSet[0]
	}
	return node
}

func (tree *NumberNode) Print() {
	fmt.Print("{")
	fmt.Print("\"number\":" + "\"" + tree.Number + "\"" + ",")
	if tree.NextNum == nil {
		fmt.Print("\"next_number\":")
		fmt.Print("null")
	} else {
		fmt.Print("\"next_number\":")
		fmt.Print("{")
		for k, v := range tree.NextNum {
			fmt.Print("\"" + k + "\"" + ":")
			v.Print()
		}
		fmt.Print("},")
	}
	fmt.Print("},")
}
