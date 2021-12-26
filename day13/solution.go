package day13

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type FoldDirection byte

const (
	FoldUpward FoldDirection = 'y'
	FoldLeft   FoldDirection = 'x'
)

type Dot struct {
	x int
	y int
}

type Fold struct {
	dir FoldDirection
	loc int
}

type Paper struct {
	dots   []Dot
	folds  []Fold
	width  int
	height int
}

func parseInput(input string) Paper {
	var (
		dots       []Dot
		folds      []Fold
		width      int
		height     int
		parseFolds bool
	)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			parseFolds = true
			continue
		}
		if parseFolds {
			// Parse folds.
			fold := strings.Split(line, "=")
			if len(fold) != 2 {
				panic(fmt.Errorf("unparsable fold: %s", line))
			}
			dir := FoldDirection(fold[0][len(fold[0])-1])
			loc, err := strconv.Atoi(fold[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, Fold{dir, loc})
			if dir == FoldUpward && loc*2+1 > height {
				height = loc*2 + 1
			}
			if dir == FoldLeft && loc*2+1 > width {
				width = loc*2 + 1
			}
		} else {
			// Parse dots.
			xy := strings.Split(line, ",")
			if len(xy) != 2 {
				panic(fmt.Errorf("unparsable dot: %s", line))
			}
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				panic(err)
			}
			dots = append(dots, Dot{x, y})
		}
	}
	return Paper{dots: dots, folds: folds, width: width, height: height}
}

func applyFold(paper *Paper) {
	fold := paper.folds[0]
	paper.folds = paper.folds[1:]
	if fold.dir == FoldUpward {
		for i := range paper.dots {
			if paper.dots[i].y > fold.loc {
				paper.dots[i].y = 2*fold.loc - paper.dots[i].y
			}
		}
		paper.height = fold.loc
	} else {
		for i := range paper.dots {
			if paper.dots[i].x > fold.loc {
				paper.dots[i].x = 2*fold.loc - paper.dots[i].x
			}
		}
		paper.width = fold.loc
	}
}

func countUniqueDots(dots []Dot) int {
	m := make(map[Dot]bool)
	for _, d := range dots {
		m[d] = true
	}
	return len(m)
}

func (p *Paper) String() string {
	var b bytes.Buffer
	m := make(map[Dot]bool)
	for _, d := range p.dots {
		m[d] = true
	}
	for row := 0; row < p.height; row++ {
		for col := 0; col < p.width; col++ {
			if m[Dot{x: col, y: row}] {
				b.WriteRune('#')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func Part1(input string) int {
	paper := parseInput(input)
	applyFold(&paper)
	return countUniqueDots(paper.dots)
}

func Part2(input string) int {
	paper := parseInput(input)
	for len(paper.folds) > 0 {
		applyFold(&paper)
	}
	fmt.Println(paper.String())
	return 0
}
