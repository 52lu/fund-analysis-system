package test

import (
	"fmt"
	"math"
	"testing"
)

type TreeNode struct {
	Data      int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

// 判断是否是二叉树
func isBalance(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(maxHigh(root.RightNode))-float64(maxHigh(root.LeftNode))) <= 1 &&
		isBalance(root.LeftNode) && isBalance(root.RightNode)
}

// 计算节点的高度
func maxHigh(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		left := maxHigh(node.LeftNode)
		right := maxHigh(node.RightNode)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

// 判断是否是二叉树
func TestBalanceTree(t *testing.T) {
	leftNode3 := &TreeNode{Data: 1, LeftNode: nil, RightNode: nil}
	leftNode2 := &TreeNode{Data: 2, LeftNode: leftNode3, RightNode: nil}
	rootLeft := &TreeNode{Data: 3, LeftNode: leftNode2, RightNode: nil}
	rightNode1 := &TreeNode{Data: 8, LeftNode: nil, RightNode: nil}
	root := &TreeNode{Data: 5, LeftNode: rootLeft, RightNode: rightNode1}
	balance := isBalance(root)
	fmt.Println(balance)
}

//func makeTree(node *TreeNode, depth int) {
//	if depth < 3 {
//		left := &TreeNode{Data: 2 * depth}
//		right := &TreeNode{Data: 4 * depth}
//		node.LeftNode = left
//		node.RightNode = right
//		makeTree(node.LeftNode,depth+1)
//		makeTree(node.RightNode,depth+1)
//	}
//}

func getTree() *TreeNode {
	// 节点
	rootLeft := &TreeNode{
		Data:      13,
		LeftNode:  &TreeNode{Data: 12, LeftNode: nil, RightNode: nil},
		RightNode: &TreeNode{Data: 14, LeftNode: nil, RightNode: nil},
	}
	rootRight := &TreeNode{
		Data:      23,
		LeftNode:  &TreeNode{Data: 22, LeftNode: nil, RightNode: nil},
		RightNode: &TreeNode{Data: 26, LeftNode: nil, RightNode: nil},
	}
	return &TreeNode{Data: 15, LeftNode: rootLeft, RightNode: rootRight}
}

func TestForeachTree(t *testing.T) {
	root := getTree()
	//var res []int
	//fontForeach(root, &res)
	//middleForeach(root,&res)
	//backForeach(root, &res)
	res := levelForeach(root)
	fmt.Println(res)
}

//  前序遍历二叉树(根-左-右)
func fontForeach(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Data)
	fontForeach(root.LeftNode, result)
	fontForeach(root.RightNode, result)
}

// 中序遍历(左-根-右)
func middleForeach(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	middleForeach(root.LeftNode, result)
	*result = append(*result, root.Data)
	middleForeach(root.RightNode, result)
}

// 后序遍历(左-右-根)
func backForeach(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	backForeach(root.LeftNode, result)
	backForeach(root.RightNode, result)
	*result = append(*result, root.Data)
}

// 层序遍历（从上到下，从左到右）
func levelForeach(node *TreeNode) []int {
	// 定义变量保存结果
	var result []int
	if node == nil {
		return result
	}
	// 定义变量保存节点
	var nodeList []*TreeNode
	nodeList = append(nodeList, node)
	// 判断节点数量是否大于0
	for len(nodeList) > 0 {
		length := len(nodeList)
		for i := 0; i < length; i++ {
			// 取出第一个节点
			node := nodeList[0]
			if node.LeftNode != nil {
				nodeList = append(nodeList, node.LeftNode)
			}
			if node.RightNode != nil {
				nodeList = append(nodeList, node.RightNode)
			}
			// 取出当前节点的值
			result = append(result, node.Data)
			// 截取节点数量，继续遍历
			nodeList = nodeList[1:]
		}
	}
	return result
}
// 判断是否是对称二叉树
func isSymmetryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetryTreeCompare(root.LeftNode,root.RightNode)
}
// 比较对称节点的值
func isSymmetryTreeCompare(left,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.Data != right.Data {
		return false
	}
	return  isSymmetryTreeCompare(left.LeftNode,right.RightNode) && isSymmetryTreeCompare(left.RightNode,right.LeftNode)
}

func TestIsSymmetryTree(t *testing.T) {
	//tree := getTree()
	tree := &TreeNode{
		Data:      5,
		LeftNode:  &TreeNode{Data: 4},
		RightNode: &TreeNode{Data: 4},
	}
	fmt.Println(isSymmetryTree(tree))
}
