package recursion

import "fmt"

func MigongDemo() {
	var map1 [8][7]int
	for i := 0; i < 7; i++ {
		map1[0][i] = 1
		map1[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		map1[i][0] = 1
		map1[i][6] = 1
	}
	//设置挡板
	map1[3][1] = 1
	map1[3][2] = 1
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(map1[i][j], " ")
		}
		fmt.Print("\n")
	}

}

//使用递归方法给小球找路

//i,j表示起始位置坐标
//小球到maparr[6][5],说明找到了通路
//约定：map[i][j]为0表示该点没有走过，为1表示墙，2表示通路可以走，3表示该点已经走过但是走不通
//走的策略，下右上左

func setWay(mapArr [][]int, i int, j int) bool {

}
