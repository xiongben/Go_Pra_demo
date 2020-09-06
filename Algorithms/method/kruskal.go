package method

import "fmt"

func KruskalDemo() {
	INF := 10000
	vertexs := []string{"A", "B", "C", "D", "E", "F", "G"}
	matrix := [][]int{
		{0, 12, INF, INF, INF, 16, 14},
		{12, 0, 10, INF, INF, 7, INF},
		{INF, 10, 0, 3, 5, 6, INF},
		{INF, INF, 3, 0, 4, INF, INF},
		{INF, INF, 5, 4, 0, 2, 8},
		{16, 7, 6, INF, 2, 0, 9},
		{14, INF, INF, INF, 8, 9, 0},
	}
	kruskalCase1 := KruskalCase{}
	kruskalCase1.create(vertexs, matrix)
	kruskalCase1.print()

	aa := kruskalCase1.getEdges()
	for _, v := range aa {
		fmt.Print(v)
	}
	kruskalCase1.kruskal()

}

type KruskalCase struct {
	edgeNum int
	vertexs []string
	matrix  [][]int
	INF     int
}

func (this *KruskalCase) create(vertexs []string, matrix [][]int) {
	this.INF = 10000
	vlen := len(vertexs)
	this.vertexs = make([]string, vlen)
	this.matrix = make([][]int, vlen)
	for i, _ := range this.matrix {
		this.matrix[i] = make([]int, vlen)
	}
	copy(this.vertexs, vertexs)
	copy(this.matrix, matrix)

	//统计边
	for i := 0; i < vlen; i++ {
		for j := i + 1; j < vlen; j++ {
			if this.matrix[i][j] != this.INF {
				this.edgeNum++
			}
		}
	}

}

func (this *KruskalCase) print() {
	for i := 0; i < len(this.vertexs); i++ {
		for j := 0; j < len(this.vertexs); j++ {
			fmt.Printf("%7d", this.matrix[i][j])
		}
		fmt.Println()
	}
}

//对边进行冒泡排序
func (this *KruskalCase) sortEdges(edges []Edata) {
	for i := 0; i < len(edges); i++ {
		for j := 0; j < len(edges)-1-i; j++ {
			if edges[j].weight > edges[j+1].weight {
				tmp := edges[j]
				edges[j] = edges[j+1]
				edges[j+1] = tmp
			}
		}
	}
}

//得到顶点对应的下标
func (this *KruskalCase) getPosition(str string) int {
	for i := 0; i < len(this.vertexs); i++ {
		if this.vertexs[i] == str {
			return i
		}
	}
	return -1
}

//得到下标为i的顶点的终点，用于判断是否是回路,ends代表各个顶点的终点，在遍历过程中逐步形成
//返回对应顶点的终点的下标
func (this *KruskalCase) getEnd(ends []int, i int) int {
	for ends[i] != 0 {
		i = ends[i]
	}
	return i
}

//获取图中的边，放到Edata中，后面需要遍历该数组 formation is [['A','B',2],,,]
func (this *KruskalCase) getEdges() []Edata {
	index := 0
	edges := make([]Edata, this.edgeNum)
	for i := 0; i < len(this.vertexs); i++ {
		for j := i + 1; j < len(this.vertexs); j++ {
			if this.matrix[i][j] != this.INF {
				edgesTmp := Edata{}
				edgesTmp.create(this.vertexs[i], this.vertexs[j], this.matrix[i][j])
				edges[index] = edgesTmp
				index++
			}
		}
	}
	return edges
}

func (this *KruskalCase) kruskal() {
	index := 0 //表示最后结果数组的索引
	ends := make([]int, this.edgeNum)
	//创建结果数组，保存最后的最小生成树
	rets := make([]Edata, this.edgeNum)

	edges := this.getEdges()
	this.sortEdges(edges)

	for i := 0; i < this.edgeNum; i++ {
		p1 := this.getPosition(edges[i].start)
		p2 := this.getPosition(edges[i].end)

		m := this.getEnd(ends, p1)
		n := this.getEnd(ends, p2)
		if m != n { //不构成回路
			ends[m] = n
			rets[index] = edges[i]
			index++
		}
	}
	//打印结果
	fmt.Println("==============")
	for _, v := range rets {
		if v.weight != 0 {
			fmt.Print(v)
		}
	}
}

type Edata struct {
	start  string
	end    string
	weight int
}

func (this *Edata) create(start, end string, weight int) {
	this.start = start
	this.end = end
	this.weight = weight
}
