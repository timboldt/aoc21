package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type LineSegment struct {
	start Point
	end   Point
}

func parseXYPair(input string) Point {
	xy := strings.Split(input, ",")
	if len(xy) != 2 {
		// Should never happen
		panic(fmt.Errorf("invalid number of points: %d", len(xy)))
	}
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		// Should never happen.
		panic(err)
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		// Should never happen.
		panic(err)
	}
	return Point{x, y}
}

func parseInput(input string) []LineSegment {
	var segments []LineSegment
	for _, line := range strings.Split(input, "\n") {
		points := strings.Split(line, " -> ")
		if len(points) != 2 {
			// Should never happen
			panic(fmt.Errorf("invalid number of points: %d", len(points)))
		}
		var seg LineSegment
		seg.start = parseXYPair(points[0])
		seg.end = parseXYPair(points[1])
		segments = append(segments, seg)
	}
	return segments
}

func linePoints(seg LineSegment) []Point {
	var (
		pts  []Point
		xInc int
		yInc int
	)

	if seg.start.x > seg.end.x {
		xInc = -1
	}
	if seg.start.x < seg.end.x {
		xInc = 1
	}
	if seg.start.y > seg.end.y {
		yInc = -1
	}
	if seg.start.y < seg.end.y {
		yInc = 1
	}

	pt := Point{seg.start.x, seg.start.y}
	for ; pt.x != seg.end.x+xInc || pt.y != seg.end.y+yInc; pt.x, pt.y = pt.x+xInc, pt.y+yInc {
		pts = append(pts, pt)
	}

	return pts
}

func debugPrint(m map[Point]int) {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			fmt.Printf("%d ", m[Point{c, r}])
		}
		fmt.Println()
	}
}

func Part1(input string) int {
	segments := parseInput(input)
	m := make(map[Point]int)
	for _, seg := range segments {
		if seg.start.x != seg.end.x && seg.start.y != seg.end.y {
			// Ignore non-rectilinear lines.
			continue
		}
		for _, pt := range linePoints(seg) {
			m[pt]++
		}
	}
	//debugPrint(m)
	var cnt int
	for _, v := range m {
		if v > 1 {
			cnt++
		}
	}
	return cnt
}

func Part2(input string) int {
	segments := parseInput(input)
	m := make(map[Point]int)
	for _, seg := range segments {
		for _, pt := range linePoints(seg) {
			m[pt]++
		}
	}
	//debugPrint(m)
	var cnt int
	for _, v := range m {
		if v > 1 {
			cnt++
		}
	}
	return cnt
}
