package main

import "fmt"

type TreeNode struct {
	Val        int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

var BST = &TreeNode{10, &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{8, nil, nil}}, &TreeNode{15, &TreeNode{13, &TreeNode{11, nil, nil}, nil}, nil}}

func JudgeBST(root *TreeNode) bool {
	var LResult, RResult, LResultRecur, RResultRecur = true, true, true, true
	var LChild = (*root).LeftChild
	var RChild = (*root).RightChild

	if root == nil {
		return false
	}
	if LChild != nil {
		if (*LChild).Val >= (*root).Val {
			LResult = false
		}
		LResultRecur = JudgeBST(LChild)
	}
	if RChild != nil {
		if (*RChild).Val <= (*root).Val {
			RResult = false
		}
		RResultRecur = JudgeBST(RChild)
	}
	return LResult && RResult && LResultRecur && RResultRecur
}

func MidScan(root *TreeNode) (result []int) {
	var resultL, resultR []int
	if root == nil {
		return result
	}
	resultL = MidScan((*root).LeftChild)
	resultR = MidScan((*root).RightChild)
	result = append(resultL, (*root).Val)
	result = append(result, resultR...)
	return result
}

func JudgeSort(slice []int) bool {
	if slice == nil {
		return false
	}
	for k := 0; k < len(slice)-1; k++ {
		if slice[k] >= slice[k+1] {
			return false
		}
	}
	return true
}

func JudgeBSTSort(root *TreeNode) bool {
	var slice []int
	var result bool
	slice = MidScan(root)
	fmt.Printf("slice is %v\n", slice)
	result = JudgeSort(slice)
	return result
}

func InsertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{val, nil, nil}
		return root
	}
	if val < (*root).Val {
		(*root).LeftChild = InsertBST((*root).LeftChild, val)
	} else if val > (*root).Val {
		(*root).RightChild = InsertBST((*root).RightChild, val)
	}
	return root
}

func DeleteBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if (*root).Val == val {
		// node is leaf
		if (*root).LeftChild == nil && (*root).RightChild == nil {
			return nil
		}
		if (*root).LeftChild == nil {
			return (*root).RightChild
		}
		if (*root).RightChild == nil {
			return (*root).LeftChild
		}
		// replace with min value of right child
		minRight := (*root).RightChild
		fatherMinRight := root
		for (*minRight).LeftChild != nil {
			fatherMinRight := minRight
			fmt.Println("fahterMinRight:%v\n", fatherMinRight)
			minRight = (*minRight).LeftChild
		}
		(*root).Val = (*minRight).Val
		if (*((*root).RightChild)).LeftChild == nil {
			(*fatherMinRight).RightChild = nil
		} else {
			(*fatherMinRight).LeftChild = nil
		}
		return root
	} else if val < (*root).Val {
		(*root).LeftChild = DeleteBST((*root).LeftChild, val)
	} else {
		(*root).RightChild = DeleteBST((*root).RightChild, val)
	}
	return root
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	var nowNode = root
	for nowNode != nil {
		if val == (*nowNode).Val {
			return nowNode
		}
		if val < (*nowNode).Val {
			nowNode = (*nowNode).LeftChild
		} else if val > (*nowNode).Val {
			nowNode = (*nowNode).RightChild
		}
	}
	return nowNode
}

func main() {
	resultBST := DeleteBST(BST, 15)
	result := MidScan(resultBST)
	fmt.Printf("result is %v\n", result)
}
