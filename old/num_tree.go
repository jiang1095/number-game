package old

import (
	"fmt"
	"math"
	"strconv"
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
		node.Number, stat_set = max_min(num_set)
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

func compare(base, compare_to string) (int, int) {
	if len(base) != 4 || len(compare_to) != 4 {
		return -1, -1
	}
	a, b := 0, 0
	for i1, v1 := range []byte(base) {
		for i2, v2 := range []byte(compare_to) {
			if v1 == v2 {
				if i1 == i2 {
					a++
				} else {
					b++
				}
			}
		}
	}
	return a, b
}

func max_min(num_set []string) (string, map[string][]string) {
	var result_num = ""
	var score = math.MaxInt32
	var result_set map[string][]string = make(map[string][]string)
	for _, base := range num_set {
		var set map[string][]string = make(map[string][]string)
		var num_score = 0
		for _, v := range num_set {
			a, b := compare(base, v)
			key := strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
			set[key] = append(set[key], v)
		}
		for _, v := range set {
			if len(v) > num_score {
				num_score = len(v)
			}
		}
		if num_score < score {
			score = num_score
			result_num = base
			result_set = set
		}
	}
	return result_num, result_set
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
