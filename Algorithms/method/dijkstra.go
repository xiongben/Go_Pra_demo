package method

import "fmt"

func DijkstraDemo() {
	vertex := []string{"A", "B", "C", "D", "E", "F", "G"}
	//vertexLen := len(vertex)
	N := 65535
	matrix := [][]int{
		{N, 5, 7, N, N, N, 2},
		{5, N, N, 9, N, N, 3},
		{7, N, N, N, 8, N, N},
		{N, 9, N, N, N, 4, N},
		{N, N, 8, N, N, 5, 4},
		{N, N, N, 4, 5, N, 6},
		{2, 3, N, N, 4, 6, N},
	}
	graph1 := Graph{}
	graph1.create(vertex, matrix)
	//graph1.showGraph()
	graph1.dijkstra(6)
	graph1.showRes()

}

type Graph struct {
	vertex []string
	matrix [][]int
	vv     VisitedVertex
}

func (this *Graph) create(vertex []string, matrix [][]int) {
	this.vertex = vertex
	this.matrix = matrix
}

func (this *Graph) showGraph() {
	for _, v := range this.matrix {
		fmt.Println(v)
	}

}

func (this *Graph) showRes() {
	this.vv.show()
}

func (this *Graph) dijkstra(index int) {
	this.vv = VisitedVertex{}
	this.vv.create(len(this.vertex), index)
	this.update(index)

	for j := 1; j < len(this.vertex); j++ {
		index = this.vv.updateArr()
		this.update(index)
	}
}

//更新index下标顶点到周围顶点的距离和周围顶点的前驱节点
func (this *Graph) update(index int) {
	len1 := 0
	for j := 0; j < len(this.matrix[index]); j++ {
		len1 = this.vv.getDis(index) + this.matrix[index][j]
		if (!this.vv.in(j)) && (len1 < this.vv.getDis(j)) {
			this.vv.updatePre(j, index)
			this.vv.updateDis(j, len1)
		}
	}
}

type VisitedVertex struct {
	alreadArr  []int
	preVisited []int
	dis        []int
}

//index是出发顶点对应的下标
func (this *VisitedVertex) create(length int, index int) {
	this.alreadArr = make([]int, length)
	this.preVisited = make([]int, length)
	this.dis = make([]int, length)
	for i, _ := range this.dis {
		this.dis[i] = 65535
	}
	this.dis[index] = 0
	this.alreadArr[index] = 1
}

//判断index顶点是否被访问过
func (this *VisitedVertex) in(index int) bool {
	return this.alreadArr[index] == 1
}

//更新出发顶点到index顶点的距离
func (this *VisitedVertex) updateDis(index int, len int) {
	this.dis[index] = len
}

func (this *VisitedVertex) updatePre(pre int, index int) {
	this.preVisited[pre] = index
}

func (this *VisitedVertex) getDis(index int) int {
	return this.dis[index]
}

func (this *VisitedVertex) updateArr() int {
	minNum := 65535
	index := 0
	for i := 0; i < len(this.alreadArr); i++ {
		if this.alreadArr[i] == 0 && this.dis[i] < minNum {
			minNum = this.dis[i]
			index = i
		}
	}
	this.alreadArr[index] = 1
	return index
}

func (this *VisitedVertex) show() {
	fmt.Println("================")
	fmt.Println(this.alreadArr)
	fmt.Println("================")
	fmt.Println(this.preVisited)
	fmt.Println("================")
	fmt.Println(this.dis)
}
