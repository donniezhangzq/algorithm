package main

import "fmt"
import "github.com/golang-collections/collections/stack"
import "github.com/golang-collections/collections/queue"

type TreeNode struct {
	Val        int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

type TreeDepth struct {
	treeNode *TreeNode
	depth    int
}

var InputTree *TreeNode = &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{7, nil, nil}, nil}}, &TreeNode{3, nil, nil}}
var SearchTree *TreeNode = &TreeNode{10, &TreeNode{6, &TreeNode{4, nil, nil}, &TreeNode{8, nil, nil}}, &TreeNode{15, &TreeNode{13, nil, nil}, &TreeNode{16, nil, nil}}}

//pre order
func RecurPreOrder(inputTree *TreeNode) (result []int) {
	if inputTree != nil {
		result = append(result, (*inputTree).Val)
		fmt.Printf("start Recur result is %v\n", result)
		resultL := RecurPreOrder((*inputTree).LeftChild)
		result = append(result, resultL...)
		resultR := RecurPreOrder((*inputTree).RightChild)
		result = append(result, resultR...)
		fmt.Printf("end Recur result is %v\n", result)
	}
	return result
}

func PreOrder(inputTree *TreeNode) (result []int) {
	if inputTree == nil {
		return result
	}
	var treeStack = stack.New()
	var nowNode = inputTree
	result = append(result, (*nowNode).Val)
	treeStack.Push(nowNode)
	for treeStack.Len() != 0 {
		for (*nowNode).LeftChild != nil {
			nowNode = (*nowNode).LeftChild
			result = append(result, (*nowNode).Val)
			treeStack.Push(nowNode)
		}
		for (*nowNode).RightChild == nil {
			if treeStack.Len() == 0 {
				return result
			}
			nowNode = treeStack.Pop().(*TreeNode)
		}
		nowNode = (*nowNode).RightChild
		result = append(result, (*nowNode).Val)
		treeStack.Push(nowNode)
	}
	return result
}

func PreOrderNew(inputTree *TreeNode) {
	var nowNode = inputTree
	var treeStack = stack.New()
	for nowNode != nil || treeStack.Len() != 0 {
		for nowNode != nil {
			fmt.Println((*nowNode).Val)
			treeStack.Push(nowNode)
			nowNode = (*nowNode).LeftChild
		}
		if treeStack.Len() != 0 {
			nowNode = treeStack.Pop().(*TreeNode)
			nowNode = (*nowNode).RightChild
		}
	}
}

//middle order
func RecurMidOrder(inputTree *TreeNode) (result []int) {
	if inputTree == nil {
		return result
	}
	resultL := RecurMidOrder((*inputTree).LeftChild)
	result = append(result, (*inputTree).Val)
	resultR := RecurMidOrder((*inputTree).RightChild)
	result = append(resultL, result...)
	result = append(result, resultR...)
	return result
}

func MidOrder(inputTree *TreeNode) (result []int) {
	var treeStack = stack.New()
	treeStack.Push((inputTree))
	var nowNode = inputTree
	for treeStack.Len() != 0 {
		for (*nowNode).LeftChild != nil {
			treeStack.Push((*nowNode).LeftChild)
			nowNode = (*nowNode).LeftChild
		}
		for (*nowNode).RightChild == nil {
			nowNode = treeStack.Pop().(*TreeNode)
			result = append(result, (*nowNode).Val)
			if treeStack.Len() == 0 {
				break
			}
		}
		if (*nowNode).RightChild != nil {
			treeStack.Push((*nowNode).RightChild)
			nowNode = (*nowNode).RightChild
		}
	}
	return result
}

func MidOrderNew(inputTree *TreeNode) {
	var nowNode = inputTree
	var treeStack = stack.New()
	for nowNode != nil || treeStack.Len() != 0 {
		for nowNode != nil {
			treeStack.Push(nowNode)
			nowNode = (*nowNode).LeftChild
		}
		if treeStack.Len() != 0 {
			nowNode = treeStack.Pop().(*TreeNode)
			fmt.Println((*nowNode).Val)
			nowNode = (*nowNode).RightChild
		}
	}
}

//post order
func RecurPostOrder(inputTree *TreeNode) (result []int) {
	if inputTree == nil {
		return result
	}
	resultL := RecurPostOrder((*inputTree).LeftChild)
	resultR := RecurPostOrder((*inputTree).RightChild)
	resultD := (*inputTree).Val
	result = append(resultL, resultR...)
	result = append(result, resultD)
	return result
}

func PostOrder(inputTree *TreeNode) (result []int) {
	if inputTree == nil {
		return result
	}
	var treeStack = stack.New()
	var nowNode *TreeNode
	var isVisit = make(map[*TreeNode]bool)
	nowNode = inputTree
	treeStack.Push(nowNode)

	for treeStack.Len() != 0 {
		for (*nowNode).LeftChild != nil && !isVisit[(*nowNode).LeftChild] {
			treeStack.Push(nowNode)
			nowNode = (*nowNode).LeftChild
		}
		for (*nowNode).RightChild == nil {
			if treeStack.Len() == 0 {
				return result
			}
			result = append(result, (*nowNode).Val)
			nowNode = treeStack.Pop().(*TreeNode)
		}
		if _, ok := isVisit[nowNode]; ok {
			result = append(result, (*nowNode).Val)
			nowNode = treeStack.Pop().(*TreeNode)
		} else {
			isVisit[nowNode] = true
			treeStack.Push(nowNode)
			nowNode = (*nowNode).RightChild
		}
	}
	return result
}

func PostOrderNew(inputTree *TreeNode) {
	var nowNode = inputTree
	var treeStack = stack.New()
	var treeMap = make(map[*TreeNode]bool)

	for nowNode != nil || treeStack.Len() != 0 {
		for _, ok := treeMap[nowNode]; nowNode != nil && !ok; {
			treeStack.Push(nowNode)
			nowNode = (*nowNode).LeftChild
		}
		if treeStack.Len() != 0 {
			nowNode = treeStack.Pop().(*TreeNode)
			if _, ok := treeMap[nowNode]; !ok && (*nowNode).RightChild != nil {
				treeMap[nowNode] = true
				treeStack.Push(nowNode)
				nowNode = (*nowNode).RightChild
				continue
			}
			fmt.Println((*nowNode).Val)
			if _, ok := treeMap[nowNode]; !ok {
				nowNode = (*nowNode).RightChild
			}
		} else {
			break
		}
	}
}

// get node number
func GetNodeNumber(inputTree *TreeNode) (number int) {
	var result []int
	var treeStack = stack.New()
	var nowNode = inputTree
	treeStack.Push(nowNode)

	for treeStack.Len() != 0 {
		for (*nowNode).LeftChild != nil {
			nowNode = (*nowNode).LeftChild
			treeStack.Push(nowNode)
		}

		for (*nowNode).RightChild == nil {
			if treeStack.Len() == 0 {
				return len(result)
			}
			nowNode = treeStack.Pop().(*TreeNode)
			result = append(result, (*nowNode).Val)
		}

		nowNode = (*nowNode).RightChild
		treeStack.Push(nowNode)
	}
	return len(result)
}

// get max depth
func GetDepth(inputTree *TreeNode) (number int) {
	var treeDepthStack = stack.New()
	var maxDepth = 0
	var nowNode *TreeNode
	var nowDepth = 0
	var nowTreeDepth TreeDepth
	if inputTree == nil {
		return 0
	}
	nowDepth = 1
	nowNode = inputTree
	treeDepthStack.Push(TreeDepth{nowNode, nowDepth})
	for treeDepthStack.Len() != 0 {
		for (*nowNode).LeftChild != nil {
			nowNode = (*nowNode).LeftChild
			nowDepth++
			treeDepthStack.Push(TreeDepth{nowNode, nowDepth})
		}
		for (*nowNode).RightChild == nil {
			if treeDepthStack.Len() == 0 {
				return maxDepth
			}
			nowTreeDepth = treeDepthStack.Pop().(TreeDepth)
			nowNode = nowTreeDepth.treeNode
			nowDepth = nowTreeDepth.depth
			if nowDepth > maxDepth {
				maxDepth = nowDepth
			}
		}
		if nowDepth > maxDepth {
			maxDepth = nowDepth
		}
		nowNode = (*nowNode).RightChild
		nowDepth++
		treeDepthStack.Push(TreeDepth{nowNode, nowDepth})
	}
	return maxDepth
}

// layer scan tree node
func LayerScan(inputTree *TreeNode) {
	var treeQueue = queue.New()
	var nowNode = inputTree
	treeQueue.Enqueue(nowNode)
	for treeQueue.Len() != 0 {
		nowNode = treeQueue.Dequeue().(*TreeNode)
		fmt.Println((*nowNode).Val)
		if (*nowNode).LeftChild != nil {
			treeQueue.Enqueue((*nowNode).LeftChild)
		}
		if (*nowNode).RightChild != nil {
			treeQueue.Enqueue((*nowNode).RightChild)
		}
	}
}

// tree convert to bidirectional linked list
func TConvertBDLink(root *TreeNode) (LLinked *TreeNode,
	RLinked *TreeNode) {
	if root == nil {
		LLinked = nil
		RLinked = nil
		return LLinked, RLinked
	}

	if (*root).LeftChild == nil {
		LLinked = root
	} else {
		LFirst, LLast := TConvertBDLink((*root).LeftChild)
		LLinked = LFirst
		(*root).LeftChild = LLast
		(*LLast).RightChild = root
	}

	if (*root).RightChild == nil {
		RLinked = root
	} else {
		RFirst, RLast := TConvertBDLink((*root).RightChild)
		RLinked = RLast
		(*root).RightChild = RFirst
		(*RFirst).LeftChild = root
	}
	return LLinked, RLinked
}

// get k layer number
func GetKLayerNumber(root *TreeNode, k int) int {
	var LNumber, RNumber int
	if k < 1 || root == nil {
		return 0
	}
	if k == 1 {
		return 1
	}
	LNumber = GetKLayerNumber((*root).LeftChild, k-1)
	RNumber = GetKLayerNumber((*root).RightChild, k-1)
	return LNumber + RNumber
}

// get Leef node number
func GetLeefNodeNumber(root *TreeNode) int {
	var LNumber, RNumber int
	if root == nil {
		return 0
	}
	if (*root).LeftChild == nil && (*root).RightChild == nil {
		return 1
	}
	LNumber = GetLeefNodeNumber((*root).LeftChild)
	RNumber = GetLeefNodeNumber((*root).RightChild)
	return LNumber + RNumber

}

// judget tree is balance binary tree
func JudgeAVL(root *TreeNode) (result bool, depth int) {
	if root == nil {
		return true, 0
	}
	resultL, depthL := JudgeAVL((*root).LeftChild)
	resultR, depthR := JudgeAVL((*root).RightChild)
	if depthL-depthR > 1 || depthL-depthR < -1 {
		result = false
	} else {
		result = true
	}
	if depthL > depthR {
		depth = depthL + 1
	} else {
		depth = depthR + 1
	}
	return result && resultL && resultR, depth
}

// get image of the tree
func GetImage(root *TreeNode) *TreeNode {
	var LNode, RNode *TreeNode
	if root == nil {
		return nil
	}
	LNode = GetImage((*root).LeftChild)
	RNode = GetImage((*root).RightChild)
	(*root).LeftChild = RNode
	(*root).RightChild = LNode

	return root
}

//find lasted common root
//flag true:find node1 or node2
func FindCommonRoot(root *TreeNode, node1 int, node2 int) (flag bool, result *TreeNode) {
	var LFlag, RFlag bool
	var LResult, RResult *TreeNode
	if root == nil {
		return false, nil
	}
	if node1 == node2 {
		return false, nil
	}
	if (*root).Val == node1 || (*root).Val == node2 {
		flag = true
	} else {
		flag = false
	}
	LFlag, LResult = FindCommonRoot((*root).LeftChild, node1, node2)
	RFlag, RResult = FindCommonRoot((*root).RightChild, node1, node2)
	flag = flag || LFlag || RFlag
	if flag && (LFlag || RFlag) {
		return false, root
	}
	if LResult != nil {
		return false, LResult
	}
	if RResult != nil {
		return false, RResult
	}
	if LFlag && RFlag {
		return true, root
	}
	return flag, nil
}

func main() {
	_, result := FindCommonRoot(InputTree, 7, 4)
	fmt.Println(result)
}
