package method

import (
	"fmt"
	"sort"
)

func HorseBoardDemo() {
	horse1 := HorseChessBoard{}
	horse1.create()
}

type HorseChessBoard struct {
	X        int    //棋盘的列数
	Y        int    //棋盘的行数
	visited  []bool //记录棋盘各个位置是否被访问过
	finished bool   //为true，表示成功
}

func (this *HorseChessBoard) create() {
	this.X = 8
	this.Y = 8
	row := 1
	column := 1
	chessBoard := make([][]int, this.X)
	for i, _ := range chessBoard {
		chessBoard[i] = make([]int, this.Y)
	}
	this.visited = make([]bool, this.X*this.Y)
	this.traversalChessboard(chessBoard, row-1, column-1, 1)
	for _, v := range chessBoard {
		for _, v2 := range v {
			fmt.Printf("%6d", v2)
		}
		fmt.Println()
	}
}

//row 当前开始位置的行
//column 当前位置开始的列
//step表示第几步
func (this *HorseChessBoard) traversalChessboard(chessboard [][]int, row int, column int, step int) {
	chessboard[row][column] = step
	this.visited[row*this.X+column] = true
	pointNow := Point{
		X: column,
		Y: row,
	}
	ps := this.next(pointNow)
	this.sort(ps)
	for len(ps) > 0 {
		p := ps[0]
		ps = ps[1:]
		if !this.visited[p.Y*this.X+p.X] {
			this.traversalChessboard(chessboard, p.Y, p.X, step+1)
		}
	}
	if (step < this.X*this.Y) && !this.finished {
		chessboard[row][column] = 0
		this.visited[row*this.X+column] = false
	} else {
		this.finished = true
	}
}

func (this *HorseChessBoard) next(curPoint Point) []Point {
	ps := make([]Point, 0)
	p1 := Point{}
	p1.X = curPoint.X - 2
	p1.Y = curPoint.Y - 1
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X - 1
	p1.Y = curPoint.Y - 2
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 1
	p1.Y = curPoint.Y - 2
	if p1.X < this.X && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 2
	p1.Y = curPoint.Y - 1
	if p1.X < this.X && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 2
	p1.Y = curPoint.Y + 1
	if p1.X < this.X && p1.Y < this.Y {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 1
	p1.Y = curPoint.Y + 2
	if p1.X < this.X && p1.Y < this.Y {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X - 1
	p1.Y = curPoint.Y + 2
	if p1.X >= 0 && p1.Y < this.Y {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X - 2
	p1.Y = curPoint.Y + 1
	if p1.X >= 0 && p1.Y < this.Y {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}
	return ps

}

func (this *HorseChessBoard) sort(ps []Point) {
	sort.Slice(ps, func(i, j int) bool {
		count1 := len(this.next(ps[i]))
		count2 := len(this.next(ps[j]))
		if count1 > count2 {
			return false
		} else {
			return true
		}
	})
}

type Point struct {
	X int
	Y int
}
