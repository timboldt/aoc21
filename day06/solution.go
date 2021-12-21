package day06

import (
	"strconv"
	"strings"
)

const MultiplyRate = 7
const SpawningTime = 0
const AfterSpawnTime = MultiplyRate - 1
const NewFishTime = AfterSpawnTime + 2

type Fish [NewFishTime + 1]int

func parseInput(input string) Fish {
	var fish Fish
	for _, raw := range strings.Split(input, ",") {
		f, err := strconv.Atoi(raw)
		if err != nil {
			// Should never happen.
			panic(err)
		}
		fish[f]++
	}
	return fish
}

func growFish(fish Fish) Fish {
	spawn := fish[SpawningTime]
	for idx := 0; idx < NewFishTime; idx++ {
		fish[idx] = fish[idx+1]
	}
	fish[AfterSpawnTime] += spawn
	fish[NewFishTime] = spawn
	return fish
}

func totalFish(fish Fish) int {
	var cnt int
	for _, f := range fish {
		cnt += f
	}
	return cnt
}

func Part1(input string) int {
	fish := parseInput(input)
	for i := 0; i < 80; i++ {
		fish = growFish(fish)
	}
	return totalFish(fish)
}

func Part2(input string) int {
	fish := parseInput(input)
	for i := 0; i < 256; i++ {
		fish = growFish(fish)
	}
	return totalFish(fish)
}
