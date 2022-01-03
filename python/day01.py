#!/usr/bin/env python

import argparse
import os.path


def parse(s: str) -> list[int]:
    return [int(x) for x in s.split()]


def part1(s: str) -> int:
    numbers = parse(s)
    return sum(numbers[i] > numbers[i - 1] for i in range(1, len(numbers)))


def part2(s: str) -> int:
    numbers = parse(s)
    return sum(numbers[i] > numbers[i - 3] for i in range(3, len(numbers)))


EXAMPLE = """\
199
200
208
210
200
207
240
269
260
263
"""


def test_parse() -> None:
    assert parse(EXAMPLE) == [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]


def test_part1() -> None:
    assert part1(EXAMPLE) == 7


def test_part2() -> None:
    assert part2(EXAMPLE) == 5


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
