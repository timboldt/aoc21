package day09

import (
	"sort"
	"strings"
)

const (
	MaxInputWidth = 100
	BorderWidth   = 1
	MaxMapWidth   = MaxInputWidth + 2*BorderWidth
)

type HeightMap [MaxMapWidth][MaxMapWidth]int

type Point struct {
	row, col int
}

func parseInput(input string) HeightMap {
	var hm HeightMap
	for row := range hm {
		for col := range hm[0] {
			hm[row][col] = 99
		}
	}
	for row, line := range strings.Split(input, "\n") {
		for col, val := range line {
			hm[row+BorderWidth][col+BorderWidth] = int(val - '0')
		}
	}
	return hm
}

func findLowPoints(hm HeightMap) []Point {
	var lps []Point
	for row := range hm {
		if row < BorderWidth || row >= MaxInputWidth+BorderWidth {
			continue
		}
		for col := range hm[0] {
			if col < BorderWidth || col >= MaxInputWidth+BorderWidth {
				continue
			}
			low := hm[row-1][col] > hm[row][col] &&
				hm[row+1][col] > hm[row][col] &&
				hm[row][col-1] > hm[row][col] &&
				hm[row][col+1] > hm[row][col]
			if low {
				lps = append(lps, Point{row: row, col: col})
			}
		}
	}
	return lps
}

func basinSize(hm HeightMap, lp Point) int {
	var total int
	visited := make(map[Point]bool)
	queue := []Point{lp}
	for {
		if len(queue) == 0 {
			break
		}
		pt := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if visited[pt] {
			continue
		}
		total++
		visited[pt] = true
		for i := 0; i < 4; i++ {
			nb := pt
			switch i {
			case 0:
				nb.row--
			case 1:
				nb.row++
			case 2:
				nb.col--
			case 3:
				nb.col++
			}
			if !visited[nb] && hm[nb.row][nb.col] < 9 {
				queue = append(queue, nb)
			}
		}
	}
	return total
}

func Part1(input string) int {
	hm := parseInput(input)
	lps := findLowPoints(hm)
	var risk int
	for _, lp := range lps {
		risk += hm[lp.row][lp.col] + 1
	}
	return risk
}

func Part2(input string) int {
	hm := parseInput(input)
	lps := findLowPoints(hm)
	var basins []int
	for _, lp := range lps {
		basins = append(basins, basinSize(hm, lp))
	}
	sort.Ints(basins)
	return basins[len(basins)-3] * basins[len(basins)-2] * basins[len(basins)-1]
}
