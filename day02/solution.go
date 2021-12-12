package day02

import (
	"strconv"
	"strings"
)

type command int

const (
	invalid command = iota
	forward
	up
	down
)

type inputOp struct {
	cmd command
	val int
}

func parseInput(input string) []inputOp {
	lookup := map[string]command{
		"forward": forward,
		"up":      up,
		"down":    down,
	}
	var ops []inputOp
	for _, line := range strings.Split(input, "\n") {
		args := strings.Split(line, " ")
		val, err := strconv.Atoi(args[1])
		if err != nil {
			// Should never happen.
			panic(err)
		}
		ops = append(ops, inputOp{
			cmd: lookup[args[0]],
			val: val,
		})
	}
	return ops
}

func Part1(input string) int {
	var pos, depth int
	ops := parseInput(input)
	for _, op := range ops {
		switch op.cmd {
		case forward:
			pos += op.val
		case up:
			depth -= op.val
		case down:
			depth += op.val
		default:
			panic("invalid op command")
		}
	}
	return pos * depth
}

func Part2(input string) int {
	ops := parseInput(input)
	return len(ops)
}
