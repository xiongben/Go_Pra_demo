package method

import "fmt"

func PrimAlgorithmDemo() {
	data := []string{"A", "B", "C", "D", "E", "F", "G"}
	verxs := len(data)
	weight := [][]int{
		{10000, 5, 7, 10000, 10000, 10000, 2},
		{5, 10000, 10000, 9, 10000, 10000, 3},
		{7, 10000, 10000, 10000, 8, 10000, 10000},
		{10000, 9, 10000, 10000, 10000, 4, 10000},
		{10000, 10000, 8, 10000, 10000, 5, 4},
		{10000, 10000, 10000, 4, 5, 10000, 6},
		{2, 3, 10000, 10000, 4, 6, 10000},
	}
	graph1 := MGraph{}
	graph1.createMGraph(verxs)
	minTree1 := MinTree{}
	minTree1.createGraph(graph1, verxs, data, weight)
	minTree1.showGraph(graph1)
	minTree1.prim(graph1, 1)
}

//创建最小生成树
type MinTree struct {
}

func (this *MinTree) createGraph(graph MGraph, verxs int, data []string, weight [][]int) {
	//var i, j int
	for i := 0; i < verxs; i++ {
		graph.data[i] = data[i]
		for j := 0; j < verxs; j++ {
			graph.weight[i][j] = weight[i][j]
		}
	}
}

func (this *MinTree) showGraph(graph MGraph) {
	for _, v := range graph.weight {
		fmt.Println(v)
	}
}

func (this *MinTree) prim(graph MGraph, v int) {
	visited := make([]int, graph.verxs)
	for i, _ := range visited {
		visited[i] = 0
	}
	visited[v] = 1 //标记为已访问
	//h1,h2记录两个顶点的下标
	h1 := -1
	h2 := -1
	minWeight := 10000
	for k := 1; k < graph.verxs; k++ {
		for i := 0; i < graph.verxs; i++ {
			for j := 0; j < graph.verxs; j++ {
				if visited[i] == 1 && visited[j] == 0 && graph.weight[i][j] < minWeight {
					minWeight = graph.weight[i][j]
					h1 = i
					h2 = j
				}
			}
		}
		//找到一条最小边
		fmt.Println("< ", graph.data[h1], ",", graph.data[h2], ">权值：", minWeight)
		visited[h1] = 1
		visited[h2] = 1
		minWeight = 10000
	}
}

type MGraph struct {
	verxs  int
	data   []string
	weight [][]int
}

func (this *MGraph) createMGraph(verxs int) {
	this.verxs = verxs
	this.data = make([]string, verxs)
	this.weight = make([][]int, verxs)
	for i := range this.weight {
		this.weight[i] = make([]int, verxs)
	}
}
