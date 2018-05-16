/*
*sort algorithm including:
*    1.bable sort algorithm
*
*Author:nineno
 */

package main

import "fmt"
import "errors"

var inputSlice = []int{1, 7, 30, 2, 8, 9, 2, 4}

// babel sort
func BableSort(inputSlice []int) error {
	if len(inputSlice) == 0 {
		return errors.New("input slice is empty")
	}
	for key1, _ := range inputSlice {
		for key2, value2 := range inputSlice[key1 : len(inputSlice)-1] {
			if value2 > inputSlice[key1+key2+1] {
				inputSlice[key1+key2] = inputSlice[key1+key2+1]
				inputSlice[key1+key2+1] = value2
			}
		}
	}
	return nil
}

// insert sort
func InsertSort(inputSlice []int) error {
	if len(inputSlice) == 0 {
		return errors.New("input slice is empty")
	}
	for key1, _ := range inputSlice[:len(inputSlice)-1] {
		for key2, value2 := range inputSlice[key1+1:] {
			if inputSlice[key1] > value2 {
				inputSlice[key1+key2+1] = inputSlice[key1]
				inputSlice[key1] = value2
			}
		}
	}
	return nil
}

// quick sort
func QuickSort(inputSlice []int, start int, end int) {
	var startStore int = start
	var endStore int = end
	if end-start > 1 {
		value := inputSlice[start]
		for start != end {
			for inputSlice[end] >= value {
				if end == start {
					break
				} else {
					end--
				}
			}
			for inputSlice[start] <= value {
				if start == end {
					break
				} else {
					start++
				}
			}
			tmp := inputSlice[end]
			inputSlice[end] = inputSlice[start]
			inputSlice[start] = tmp
		}
		inputSlice[startStore] = inputSlice[end]
		inputSlice[end] = value
		QuickSort(inputSlice, startStore, end)
		QuickSort(inputSlice, end+1, endStore)
	}
}

//heap sort
func HeapSort(inputSlice []int) error {
	if len(inputSlice) == 0 {
		return errors.New("inputSlice is empty")
	}
	var size = len(inputSlice)
	for key, _ := range inputSlice {
		adjustHeap(inputSlice, key, size)
	}
	fmt.Printf("create heap is %v\n", inputSlice)
	for key, _ := range inputSlice {
		swap(inputSlice, 0, size-key-1)
		fmt.Printf("swap key:%d and key:%d,result is%v\n", 0, size-key-1, inputSlice)
		adjustHeap(inputSlice, 0, size-key-1)
		fmt.Printf("after adjust is %v\n", inputSlice)
	}
	return nil
}

func adjustHeap(inputSlice []int, key int, size int) {
	var leftChild = 2*key + 1
	var rightChild = 2 * (key + 1)
	var maxPoint = key
	if leftChild < size && inputSlice[leftChild] > inputSlice[maxPoint] {
		maxPoint = leftChild
		fmt.Printf("left max,leftChild:%d,inputSlice[leftChild]:%d\n", leftChild, inputSlice[leftChild])
	}
	if rightChild < size && inputSlice[rightChild] > inputSlice[maxPoint] {
		maxPoint = rightChild
		fmt.Printf("right max,rightChild:%d,inputSlice[rightChild]:%d\n", rightChild, inputSlice[rightChild])
	}
	fmt.Printf("ajdustHeap:key:%d, inputSlice[key]:%d,"+
		"leftChild:%d,rightChild:%d"+
		"maxPoint:%d, inputSlice[maxPoint]:%d\n",
		key, inputSlice[key], leftChild, rightChild,
		maxPoint, inputSlice[maxPoint])
	if maxPoint != key {
		swap(inputSlice, key, maxPoint)
		fmt.Printf("adjustHeap:after swap is:%v\n", inputSlice)
		adjustHeap(inputSlice, maxPoint, size)
	}
}

func swap(inputSlice []int, key1 int, key2 int) {
	tmp := inputSlice[key1]
	inputSlice[key1] = inputSlice[key2]
	inputSlice[key2] = tmp
}

// shell sort
func ShellSort(inputSlice []int) error {
	if len(inputSlice) == 0 {
		return errors.New("inputSlice is empty")
	}
	var d int
	d = len(inputSlice) / 2
	for d >= 1 {
		shellInsertSort(inputSlice, d)
		d = d / 2
	}
	return nil
}

func shellInsertSort(inputSlice []int, d int) {
	for i := 0; i < len(inputSlice); i += d {
		for j := i + d; j < len(inputSlice); j += d {
			if inputSlice[i] > inputSlice[j] {
				swap(inputSlice, i, j)
			}
		}
	}
}

// guibing sort
func GuiBingSort(inputSlice []int, start int, end int) {
	if end-start > 1 {
		GuiBingSort(inputSlice, start, start+(end-start)/2)
		GuiBingSort(inputSlice, start+(end-start)/2, end)
		for i := start; i <= end; i++ {
			for j := i; j <= end; j++ {
				if inputSlice[i] > inputSlice[j] {
					swap(inputSlice, i, j)
				}
			}
		}
	}
}

//number sort
func NumberSort(inputSlice []int) error {
	if len(inputSlice) == 0 {
		return errors.New("inputSlice is empty")
	}
	var max = inputSlice[0]
	for _, val := range inputSlice {
		if max < val {
			max = val
		}
	}
	var countArray = make([]int, max+1)
	var resultArray = make([]int, len(inputSlice))

	for key, _ := range inputSlice {
		countArray[inputSlice[key]]++
	}

	for i := 2; i < max+1; i++ {
		countArray[i] += countArray[i-1]
	}

	for j := 0; j < len(inputSlice); j++ {
		fmt.Println("countArray:", countArray)
		fmt.Println("inputSlice[j]:", inputSlice[j])
		resultArray[countArray[inputSlice[j]]-1] = inputSlice[j]
		fmt.Println("resultArray:", resultArray)
		countArray[inputSlice[j]]--
	}
	copy(inputSlice, resultArray)
	return nil
}

func main() {
	fmt.Println("before sort is:", inputSlice)
	NumberSort(inputSlice)
	fmt.Println("after sort is:", inputSlice)
}
