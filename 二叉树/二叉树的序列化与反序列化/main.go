package main

import (
	"strconv"
	"strings"
)

// 297

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	codec := Codec{}
	return codec
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := ""
	this.encode(root, &res)
	return res
}

func (this *Codec) encode(root *TreeNode, res *string) {
	if root == nil {
		*res += "#,"
		return
	}
	v := root.Val
	*res += strconv.Itoa(v) + ","
	this.encode(root.Left, res)
	this.encode(root.Right, res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	// 这句有没有都可以
	// nodes = nodes[:len(nodes)-1]
	return this.decode(&nodes)
}

// 因为函数体里有移除nodes数组第一个元素的操作，所以参数nodes的类型必须是 *[]string。不能是 []string
func (this *Codec) decode(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	node := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if node == "#" {
		return nil
	}
	rootVal, _ := strconv.Atoi(node)
	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	root.Left = this.decode(nodes)
	root.Right = this.decode(nodes)
	return root
}
