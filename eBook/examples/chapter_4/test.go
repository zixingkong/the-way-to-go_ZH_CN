package main

import (
	"container/heap"
	"fmt"
)

// 定义一个堆，用来存储霍夫曼树
type HuffmanHeap []*HuffmanTree

// 重写堆的三个方法
func (h HuffmanHeap) Len() int {
	return len(h)
}

func (h HuffmanHeap) Less(i, j int) bool {
	return h[i].frequency < h[j].frequency
}

func (h HuffmanHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 定义Push方法
func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanTree))
}

// 定义Pop方法
func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 定义霍夫曼树
type HuffmanTree struct {
	data        byte  // 字符
	frequency   int   // 频率
	left, right *HuffmanTree
}

// 构建霍夫曼树
func buildHuffmanTree(data []byte, frequency []int, n int) *HuffmanTree {
	// 初始化一个霍夫曼堆
	h := &HuffmanHeap{}
	for i := 0; i < n; i++ {
		node := &HuffmanTree{
			data:      data[i],
			frequency: frequency[i],
		}
		heap.Push(h, node)
	}

	// 将堆中的元素合并
	for h.Len() > 1 {
		// 弹出两个元素
		node1 := heap.Pop(h).(*HuffmanTree)
		node2 := heap.Pop(h).(*HuffmanTree)

		// 新建一棵树
		node := &HuffmanTree{
			data:      0,
			frequency: node1.frequency + node2.frequency,
			left:      node1,
			right:     node2,
		}
		heap.Push(h, node)
	}

	// 返回堆顶元素
	return heap.Pop(h).(*HuffmanTree)
}

// 定义一个map用来记录编码
var codeMap map[byte]string

// 遍历霍夫曼树，生成霍夫曼编码
func generateHuffmanCode(node *HuffmanTree, code string) {
	if node == nil {
		return
	}
	// 节点为叶子节点
	if node.data != 0 {
		codeMap[node.data] = code
		return
	}
	// 节点不是叶子节点
	generateHuffmanCode(node.left, code+"0")
	generateHuffmanCode(node.right, code+"1")
}

// 霍夫曼编码
func huffmanCode(data []byte, frequency []int, n int) {
	// 构建霍夫曼树
	root := buildHuffmanTree(data, frequency, n)

	// 生成霍夫曼编码
	codeMap = make(map[byte]string)
	generateHuffmanCode(root, "")

	// 输出编码结果
	for i := 0; i < n; i++ {
		fmt.Printf("%c: %s\n", data[i], codeMap[data[i]])
	}
}

func main() {
	data := []byte{'a', 'b', 'c', 'd', 'e'}
	frequency := []int{5, 9, 12, 13, 16}
	n := len(data)
	huffmanCode(data, frequency, n)
}


/*
a: 1100
b: 1101
c: 100
d: 101
e: 0
*/
