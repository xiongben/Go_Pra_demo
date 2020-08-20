package graph

import "fmt"

func GraphDemo() {
	graph1 := Graph{}
	graph1.newGraph(5)
	graph1.vertexList = []string{"A", "B", "C", "D", "E"}
	graph1.insertEdge(0, 1, 1)
	graph1.insertEdge(0, 2, 1)
	graph1.insertEdge(1, 2, 1)
	graph1.insertEdge(1, 3, 1)
	graph1.insertEdge(1, 4, 1)
	graph1.showGraph()
	//graph1.dfs2()
	//graph1.isVisited[0] = true
	//graph1.dfs3(0)
	graph1.bfs(0)
	fmt.Println(graph1.que)
}

type Graph struct {
	vertexList []string
	isVisited  []bool
	edges      [][]int
	numOfEdges int
	sum        int
	que        []int
}

func (this *Graph) showGraph() {
	for i := range this.edges {
		fmt.Println(this.edges[i])
	}
}

func (this *Graph) getFirstNeighbor(index int) int {
	for i := range this.vertexList {
		if this.edges[index][i] > 0 {
			return i
		}
	}
	return -1
}

func (this *Graph) getNextNeighbor(v1, v2 int) int {
	for j := v2 + 1; j < len(this.vertexList); j++ {
		if this.edges[v1][j] > 0 {
			return j
		}
	}
	return -1
}

func (this *Graph) dfs(isVisited []bool, i int) {
	fmt.Printf("%v -> ", this.vertexList[i])
	isVisited[i] = true
	w := this.getFirstNeighbor(i)
	for w != -1 {
		if !isVisited[w] {
			this.dfs(isVisited, w)
		}
		w = this.getNextNeighbor(i, w)
	}
}

func (this *Graph) dfs2() {
	//isVis := make([]bool,len(this.vertexList))
	for i := range this.vertexList {
		if !this.isVisited[i] {
			this.dfs(this.isVisited, i)
		}
	}
}

func (this *Graph) dfs3(i int) {
	fmt.Printf("%v -> ", this.vertexList[i])
	this.sum++
	if this.sum >= len(this.vertexList) {
		return
	}
	for j := range this.vertexList {
		if this.edges[i][j] > 0 && !this.isVisited[j] {
			this.isVisited[j] = true
			this.dfs3(j)
		}
	}
}

func (this *Graph) bfs(index int) {
	head := 0
	tail := 0
	this.que = make([]int, len(this.vertexList))
	this.que[tail] = index
	tail++
	this.isVisited[index] = true
	for head < tail && tail < len(this.vertexList) {
		cur := this.que[head]
		for i := range this.vertexList {
			if this.edges[cur][i] > 0 && !this.isVisited[i] {
				this.que[tail] = i
				tail++
				this.isVisited[i] = true
			}
			if tail > len(this.vertexList) {
				goto Loop
			}
		}
		head++
	}
Loop:
}

func (this *Graph) newGraph(n int) {
	this.edges = make([][]int, n)
	for i := range this.edges {
		this.edges[i] = make([]int, n)
	}
	this.vertexList = make([]string, n)
	this.numOfEdges = 0
	this.isVisited = make([]bool, n)
	this.sum = 0
}

func (this *Graph) insertVertex(vertex string) {
	this.vertexList = append(this.vertexList, vertex)
}

func (this *Graph) insertEdge(v1 int, v2 int, weight int) {
	this.edges[v1][v2] = weight
	this.edges[v2][v1] = weight
	this.numOfEdges++
}
