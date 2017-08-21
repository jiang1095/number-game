package helper

import (
	"fmt"
)

type NumberNode struct {
	Parent_num *NumberNode
	Num_set    []string
	Number     string
	Next_num   map[string]*NumberNode
}

func NewTree(number_set []string) *NumberNode {
	return buildTree(number_set)
}

func buildTree(num_set []string) *NumberNode {
	var node *NumberNode = &NumberNode{
		Parent_num: nil,
		Num_set:    num_set,
		Number:     "",
		Next_num:   nil,
	}
	if len(num_set) > 1 {
		var stat_set map[string][]string
		node.Number, stat_set = MaxMin(num_set)
		node.Next_num = make(map[string]*NumberNode)
		for k, v := range stat_set {
			child_node := buildTree(v)
			node.Next_num[k] = child_node
			child_node.Parent_num = node
		}
	} else {
		node.Number = num_set[0]
	}
	return node
}

func (tree *NumberNode) Print() {
	fmt.Print("{")
	fmt.Print("\"number\":" + "\"" + tree.Number + "\"" + ",")
	if tree.Next_num == nil {
		fmt.Print("\"next_number\":")
		fmt.Print("null")
	} else {
		fmt.Print("\"next_number\":")
		fmt.Print("{")
		for k, v := range tree.Next_num {

			fmt.Print("\"" + k + "\"" + ":")
			v.Print()

		}
		fmt.Print("},")
	}
	fmt.Print("},")
}
