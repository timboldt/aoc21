#!/usr/bin/env python

import argparse
import os.path


def parse(s: str) -> list[(str, int)]:
    result = list()
    for line in s.splitlines():
        dir, raw_n = line.split()
        result.append((dir, int(raw_n)))
    return result


def part1(s: str) -> int:
    commands = parse(s)
    pos = 0
    depth = 0
    for cmd in commands:
        if cmd[0] == "up":
            depth -= cmd[1]
        elif cmd[0] == "down":
            depth += cmd[1]
        elif cmd[0] == "forward":
            pos += cmd[1]

    return pos * depth


def part2(s: str) -> int:
    commands = parse(s)
    pos = 0
    depth = 0
    aim = 0
    for cmd in commands:
        if cmd[0] == "up":
            aim -= cmd[1]
        elif cmd[0] == "down":
            aim += cmd[1]
        elif cmd[0] == "forward":
            pos += cmd[1]
            depth += aim * cmd[1]

    return pos * depth


EXAMPLE = """\
forward 5
down 5
forward 8
up 3
down 8
forward 2"""


def test_parse() -> None:
    assert parse(EXAMPLE) == [
        ("forward", 5),
        ("down", 5),
        ("forward", 8),
        ("up", 3),
        ("down", 8),
        ("forward", 2),
    ]


def test_part1() -> None:
    assert part1(EXAMPLE) == 150


def test_part2() -> None:
    assert part2(EXAMPLE) == 900


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
