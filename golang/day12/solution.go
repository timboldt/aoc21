package day12

import (
	"fmt"
	"strings"
)

type NodeList []string
type Network map[string]NodeList

func parseInput(input string) Network {
	net := make(Network)
	for _, line := range strings.Split(input, "\n") {
		nodes := strings.Split(line, "-")
		if len(nodes) != 2 {
			panic(fmt.Errorf("invalid input %q", line))
		}
		// Add both the forward and reverse direction.
		net[nodes[0]] = append(net[nodes[0]], nodes[1])
		net[nodes[1]] = append(net[nodes[1]], nodes[0])
	}
	return net
}

func isSmallCave(node string) bool {
	return len(node) > 0 && node[0] >= 'a' && node[0] <= 'z'
}

// Given a start of a path, find the unique paths from there to the end.
func countPaths(net Network, visited NodeList, extraVisitAllowed bool) int {
	var paths int
	currNode := visited[len(visited)-1]
	if currNode == "end" {
		return 1
	}
	// Find the neighbors of the last node.
NODELIST:
	for _, node := range net[currNode] {
		if node == visited[0] {
			continue
		}

		// Use a local shadow variable so it can be reset on each time round the loop.
		extraVisitAllowed := extraVisitAllowed

		// Small caves can only be visited once, unless extraVisitAllowed.
		if isSmallCave(node) {
			for _, v := range visited {
				if node == v {
					if extraVisitAllowed {
						extraVisitAllowed = false
					} else {
						continue NODELIST
					}
				}
			}
		}
		paths += countPaths(net, append(visited, node), extraVisitAllowed)
	}
	return paths
}

func Part1(input string) int {
	net := parseInput(input)
	return countPaths(net, NodeList{"start"}, false)
}

func Part2(input string) int {
	net := parseInput(input)
	return countPaths(net, NodeList{"start"}, true)
}
