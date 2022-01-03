#!/usr/bin/env python

import heapq
import numpy
import os.path


class Node:
    def __init__(self, row, col, risk):
        self.row = row
        self.col = col
        self.risk = risk
        self.visited = False
        if self.row == 0 and self.col == 0:
            self.dist = 0
        else:
            self.dist = int(1e9)

    def __str__(self):
        return "({},{})->(r:{} v:{} d:{})".format(
            self.row, self.col, self.risk, self.visited, self.dist
        )

    def neighbors(self, nodes):
        n = []
        for delta in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            loc = (self.row + delta[0], self.col + delta[1])
            if loc in nodes and not nodes[loc].visited:
                n.append(nodes[loc])
        return n


class NodeQueue:
    def __init__(self):
        self._q = []

    def push(self, row, col, dist):
        heapq.heappush(self._q, (dist, (row, col)))

    def pop(self):
        _, loc = heapq.heappop(self._q)
        return loc

    def empty(self):
        return len(self._q) == 0


# TODO: Use a more efficient data structure!
def parse(input, multiple):
    nodes = dict()
    row = 0
    for line in input.splitlines():
        col = 0
        for ch in line:
            nodes[(row, col)] = Node(row, col, int(ch))
            col += 1
        row += 1
    base_rows = row
    base_cols = col
    rows = base_rows * multiple
    cols = base_cols * multiple
    for row in range(rows):
        for col in range(cols):
            if row == row % base_rows and col == col % base_cols:
                continue
            source_node = nodes[(row % base_rows, col % base_cols)]
            risk = (source_node.risk - 1 + row // base_rows + col // base_cols) % 9 + 1
            nodes[(row, col)] = Node(row, col, risk)
    return nodes, rows, cols


def dikstra(nodes, rows, cols):
    queue = NodeQueue()
    queue.push(0, 0, 0)
    while not queue.empty():
        loc = queue.pop()
        if loc == (rows - 1, cols - 1):
            return nodes[loc].dist
        curr = nodes[loc]
        if curr.visited:
            continue
        curr.visited = True
        for n in curr.neighbors(nodes):
            n.dist = min(n.dist, curr.dist + n.risk)
            queue.push(n.row, n.col, n.dist)
    return 0


def part1(s: str) -> int:
    nodes, rows, cols = parse(s, 1)
    return dikstra(nodes, rows, cols)


def part2(s: str) -> int:
    nodes, rows, cols = parse(s, 5)
    return dikstra(nodes, rows, cols)


EXAMPLE_1 = """\
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
"""

EXAMPLE_2 = """\
123
456
789"""


def test_parse_1() -> None:
    nodes, rows, cols = parse(EXAMPLE_1, 1)
    risks = []
    for r in range(rows):
        risks.append([])
        for c in range(cols):
            risks[r].append(0)
    for k in nodes:
        n = nodes[k]
        risks[n.row][n.col] = n.risk
    assert risks == [
        [1, 1, 6, 3, 7, 5, 1, 7, 4, 2],
        [1, 3, 8, 1, 3, 7, 3, 6, 7, 2],
        [2, 1, 3, 6, 5, 1, 1, 3, 2, 8],
        [3, 6, 9, 4, 9, 3, 1, 5, 6, 9],
        [7, 4, 6, 3, 4, 1, 7, 1, 1, 1],
        [1, 3, 1, 9, 1, 2, 8, 1, 3, 7],
        [1, 3, 5, 9, 9, 1, 2, 4, 2, 1],
        [3, 1, 2, 5, 4, 2, 1, 6, 3, 9],
        [1, 2, 9, 3, 1, 3, 8, 5, 2, 1],
        [2, 3, 1, 1, 9, 4, 4, 5, 8, 1],
    ]


def test_parse_2() -> None:
    nodes, rows, cols = parse(EXAMPLE_2, 2)
    risks = []
    for r in range(rows):
        risks.append([])
        for c in range(cols):
            risks[r].append(0)
    for k in nodes:
        n = nodes[k]
        risks[n.row][n.col] = n.risk
    print(risks)
    assert risks == [
        [1, 2, 3, 2, 3, 4],
        [4, 5, 6, 5, 6, 7],
        [7, 8, 9, 8, 9, 1],
        [2, 3, 4, 3, 4, 5],
        [5, 6, 7, 6, 7, 8],
        [8, 9, 1, 9, 1, 2],
    ]


def test_part1() -> None:
    assert part1(EXAMPLE_1) == 40


def test_part2() -> None:
    assert part2(EXAMPLE_1) == 315


def main() -> int:
    fname = os.path.join(
        os.path.dirname(__file__),
        "data",
        os.path.splitext(os.path.basename(__file__))[0] + ".txt",
    )

    with open(fname) as f:
        input = f.read()
        print(part1(input))
        print(part2(input))

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
