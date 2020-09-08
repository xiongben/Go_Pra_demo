package method

func HorseBoardDemo() {

}

type HorseChessBoard struct {
	X int //棋盘的列数
	Y int //棋盘的行数
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
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 2
	p1.Y = curPoint.Y - 1
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 2
	p1.Y = curPoint.Y + 1
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X + 1
	p1.Y = curPoint.Y + 2
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X - 1
	p1.Y = curPoint.Y + 2
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	p1.X = curPoint.X - 2
	p1.Y = curPoint.Y + 1
	if p1.X >= 0 && p1.Y >= 0 {
		ps = append(ps, Point{X: p1.X, Y: p1.Y})
	}

	return ps

}

type Point struct {
	X int
	Y int
}
