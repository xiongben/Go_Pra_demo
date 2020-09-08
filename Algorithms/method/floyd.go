package method

import "fmt"

func FloydDemo() {
	vertex := []string{"A", "B", "C", "D", "E", "F", "G"}
	//vertexLen := len(vertex)
	N := 65535
	matrix := [][]int{
		{0, 5, 7, N, N, N, 2},
		{5, 0, N, 9, N, N, 3},
		{7, N, 0, N, 8, N, N},
		{N, 9, N, 0, N, 4, N},
		{N, N, 8, N, 0, 5, 4},
		{N, N, N, 4, 5, 0, 6},
		{2, 3, N, N, 4, 6, 0},
	}
	graph1 := Graph2{}
	graph1.create(len(vertex), vertex, matrix)
	graph1.floyd()
	graph1.showGraph()

}

type Graph2 struct {
	vertex []string
	dis    [][]int //从各个顶点出发到其他顶点的距离
	pre    [][]int
}

func (this *Graph2) create(length int, vertex []string, matrix [][]int) {
	this.vertex = vertex
	this.dis = matrix
	this.pre = make([][]int, length)
	for i, _ := range this.pre {
		this.pre[i] = make([]int, length)
		for j, _ := range this.pre[i] {
			this.pre[i][j] = i
		}
	}

}

func (this *Graph2) showGraph() {
	fmt.Println("======dis=========")
	for _, v := range this.dis {
		fmt.Println(v)
	}
	fmt.Println("======pre=========")
	for _, v := range this.pre {
		fmt.Println(v)
	}
}

func (this *Graph2) floyd() {
	lenNum := 0
	for k := 0; k < len(this.dis); k++ { //k是中间节点
		for i := 0; i < len(this.dis); i++ {
			for j := 0; j < len(this.dis); j++ {
				lenNum = this.dis[i][k] + this.dis[k][j]
				if lenNum < this.dis[i][j] {
					this.dis[i][j] = lenNum
					this.pre[i][j] = this.pre[k][j]
				}
			}
		}
	}
}
