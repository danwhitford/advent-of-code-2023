package day3

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
	"strings"
)

// import "aoc2023/helpers"

type Day3 struct {
	grid grid
}

type point [2]int

type grid struct {
	width  int
	height int
	cells  []byte
}

func newGrid(width, height int) grid {
	cells := make([]byte, width*height)
	return grid{width, height, cells}
}

func (g grid) iToXY(i int) point {
	x := i % g.width
	y := i / g.width
	return point{x, y}
}

func (g grid) xYToI(xy point) int {
	x := xy[0]
	y := xy[1]
	return (y * g.width) + x
}

func (g *grid) setXY(xy point, cell byte) {
	i := g.xYToI(xy)
	g.cells[i] = cell
}

func (g grid) getXY(xy point) byte {
	return g.cells[g.xYToI(xy)]
}

func (g grid) String() string {
	var sb strings.Builder
	for i, cell := range g.cells {
		if i > 0 && i%g.width == 0 {
			sb.WriteByte('\n')
		}
		sb.WriteByte(cell)
	}
	return sb.String()
}

func (g grid) rows() [][]byte {
	rows := make([][]byte, 0)
	row := make([]byte, 0)
	for i, cell := range g.cells {
		if i > 0 && i%g.width == 0 {
			rows = append(rows, row)
			row = make([]byte, 0)
		}
		row = append(row, cell)
	}
	rows = append(rows, row)
	return rows
}

func (g grid) getNumbers() []number {
	nums := make([]number, 0)
	for y, row := range g.rows() {
		for _, n := range parseRow(row) {
			n.row = y
			nums = append(nums, n)
		}
	}
	return nums
}

func (g grid) inBounds(xy point) bool {
	x := xy[0]
	y := xy[1]
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

func (g grid) getNeighbours(num number) []byte {
	nn := make([]byte, 0)
	for x := num.start - 1; x <= num.end+1; x++ {
		xy := point{x, num.row - 1}
		if g.inBounds(xy) {
			n := g.getXY(point{x, num.row - 1})
			nn = append(nn, n)
		}
	}
	if g.inBounds(point{num.start - 1, num.row}) {
		n := g.getXY(point{num.start - 1, num.row})
		nn = append(nn, n)
	}
	if g.inBounds(point{num.end + 1, num.row}) {
		n := g.getXY(point{num.end + 1, num.row})
		nn = append(nn, n)
	}
	for x := num.start - 1; x <= num.end+1; x++ {
		xy := point{x, num.row + 1}
		if g.inBounds(xy) {
			n := g.getXY(point{x, num.row + 1})
			nn = append(nn, n)
		}
	}
	return nn
}

func (g grid) getAdjacentGears(num number) []point {
	gears := make([]point, 0)
	for x := num.start - 1; x <= num.end+1; x++ {
		xy := point{x, num.row - 1}
		if g.inBounds(xy) {
			n := g.getXY(point{x, num.row - 1})
			if n == '*' {
				gears = append(gears, xy)
			}
		}
	}
	if g.inBounds(point{num.start - 1, num.row}) {
		n := g.getXY(point{num.start - 1, num.row})
		if n == '*' {
			gears = append(gears, point{num.start - 1, num.row})
		}
	}
	if g.inBounds(point{num.end + 1, num.row}) {
		n := g.getXY(point{num.end + 1, num.row})
		if n == '*' {
			gears = append(gears, point{num.end + 1, num.row})
		}
	}
	for x := num.start - 1; x <= num.end+1; x++ {
		xy := point{x, num.row + 1}
		if g.inBounds(xy) {
			n := g.getXY(point{x, num.row + 1})
			if n == '*' {
				gears = append(gears, xy)
			}
		}
	}
	return gears
}

type number struct {
	val   int
	start int
	end   int
	row   int
}

func parseRow(row []byte) []number {
	nums := make([]number, 0)
	var sb strings.Builder
	n := number{}
	for x, cell := range row {
		if cell >= '0' && cell <= '9' {
			if sb.Len() == 0 {
				n.start = x
			}
			sb.WriteByte(cell)
		} else {
			if sb.Len() > 0 {
				n.end = x - 1
				d, _ := strconv.Atoi(sb.String())
				n.val = d
				nums = append(nums, n)
				sb.Reset()
				n = number{}
			}
		}
	}
	if sb.Len() > 0 {
		n.end = len(row) - 1
		d, _ := strconv.Atoi(sb.String())
		n.val = d
		nums = append(nums, n)
	}
	return nums
}

func containsSymbol(syms []byte) bool {
	return !helpers.Every[byte](syms, func(t byte) bool {
		return (t >= '0' && t <= '9') || t == '.'
	})
}

func (day *Day3) SetInput(s string) {
	lines := helpers.GetLines(s)
	width := len(lines[0])
	height := len(lines)
	g := newGrid(width, height)
	for y, line := range lines {
		for x, cell := range line {
			g.setXY(point{x, y}, byte(cell))
		}
	}
	day.grid = g
}

func (day Day3) SolvePart1() string {
	total := 0

	for _, n := range day.grid.getNumbers() {
		if containsSymbol(day.grid.getNeighbours(n)) {
			total += n.val
		}
	}
	return fmt.Sprint(total)
}

func (day Day3) SolvePart2() string {
	gearAdjacencyMap := make(map[point][]number)
	for _, n := range day.grid.getNumbers() {
		adjacentGears := day.grid.getAdjacentGears(n)
		for _, gear := range adjacentGears {
			gearAdjacencyMap[gear] = append(gearAdjacencyMap[gear], n)
		}
	}

	total := 0
	for _, adjacentNumbers := range gearAdjacencyMap {
		if len(adjacentNumbers) == 2 {
			total += (adjacentNumbers[0].val * adjacentNumbers[1].val)
		}
	}
	return fmt.Sprint(total)
}
