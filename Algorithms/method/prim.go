package method

func PrimAlgorithmDemo() {

}

//创建最小生成树
type MinTree struct {
}

func (this *MinTree) createGraph(graph MGraph, verxs int, data []string, weight [][]int) {
	var i, j int
	for i = 0; i < verxs; i++ {
		graph.data[i] = data[i]
		for j = 0; j < verxs; j++ {
			graph.weight[i][j] = weight[i][j]
		}
	}
}

func (this *MinTree) showGraph() {

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
