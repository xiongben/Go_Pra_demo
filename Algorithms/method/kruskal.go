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
		for j := 0; j < vlen; j++ {
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
