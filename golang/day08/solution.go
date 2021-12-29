package day08

import (
	"fmt"
	"strings"
)

type Observation struct {
	signals []string
	outputs []string
}

type DS struct {
	d int
	s rune
}

func parseInput(input string) []Observation {
	var observations []Observation
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			panic(fmt.Errorf("invalid number of observation parts: %d", len(parts)))
		}
		obs := Observation{
			signals: strings.Split(parts[0], " "),
			outputs: strings.Split(parts[1], " "),
		}
		if len(obs.signals) != 10 {
			panic(fmt.Errorf("invalid number of observation signals: %d", len(obs.signals)))
		}
		if len(obs.outputs) != 4 {
			panic(fmt.Errorf("invalid number of observation outputs: %d", len(obs.outputs)))
		}
		observations = append(observations, obs)
	}
	return observations
}

func guessDigit(outp string) int {
	switch len(outp) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5:
		// Ambiguous: 2, 3, 5
		return -5
	case 6:
		// Ambiguous: 0, 6, 9
		return -6
	case 7:
		return 8
	default:
		panic(fmt.Errorf("invalid number of output digit signals: %d", len(outp)))
	}
}

func deduceOutput(obs Observation) int {
	// Count up signal values by digit group.
	counts := make(map[DS]int)
	for _, s := range obs.signals {
		digit := guessDigit(s)
		for _, r := range s {
			counts[DS{d: digit, s: r}]++
		}
	}

	// Deduce signal to line segment mapping.
	//   S0
	// S1  S2
	//   S3
	// S4  S5
	//   S6
	var segs [7]rune
	for r := 'a'; r <= 'g'; r++ {
		// S0 is present in D7 and not D1.
		if counts[DS{d: 7, s: r}] == 1 && counts[DS{d: 1, s: r}] == 0 {
			segs[0] = r
		}
		// S1 is present in D4 and only one of D[2,3,5].
		if counts[DS{d: 4, s: r}] == 1 && counts[DS{d: -5, s: r}] == 1 {
			segs[1] = r
		}
		// S2 is present in D1 and two of D[0,6,9].
		if counts[DS{d: 1, s: r}] == 1 && counts[DS{d: -6, s: r}] == 2 {
			segs[2] = r
		}
		// S3 is present in D4, two of D[0,6,9], and not D1.
		if counts[DS{d: 4, s: r}] == 1 && counts[DS{d: -6, s: r}] == 2 && counts[DS{d: 1, s: r}] == 0 {
			segs[3] = r
		}
		// S4 is present in two of D[0,6,9], and one of D[2,3,5].
		if counts[DS{d: -6, s: r}] == 2 && counts[DS{d: -5, s: r}] == 1 {
			segs[4] = r
		}
		// S5 is present in D1, and three of D[0,6,9].
		if counts[DS{d: 1, s: r}] == 1 && counts[DS{d: -6, s: r}] == 3 {
			segs[5] = r
		}
		// S6 is present in three of D[0,6,9], but not D4 or D7.
		if counts[DS{d: -6, s: r}] == 3 && counts[DS{d: 4, s: r}] == 0 && counts[DS{d: 7, s: r}] == 0 {
			segs[6] = r
		}
	}

	var answer int
	for _, outp := range obs.outputs {
		answer *= 10
		if strings.ContainsRune(outp, segs[0]) {
			// Must be 0, 2, 3, 5, 6, 7, 8, 9.
			if strings.ContainsRune(outp, segs[4]) {
				// Must be 0, 2, 6, 8.
				if !strings.ContainsRune(outp, segs[3]) {
					// Must be 0.
				} else {
					// Must be 2, 6, 8.
					if !strings.ContainsRune(outp, segs[1]) {
						answer += 2
					} else {
						// Must be 6, 8.
						if strings.ContainsRune(outp, segs[2]) {
							answer += 8
						} else {
							answer += 6
						}
					}
				}
			} else {
				// Must be 3, 5, 7, 9.
				if strings.ContainsRune(outp, segs[1]) {
					// Must be 5, 9.
					if strings.ContainsRune(outp, segs[2]) {
						answer += 9
					} else {
						answer += 5
					}
				} else {
					// Must be 3, 7.
					if strings.ContainsRune(outp, segs[3]) {
						answer += 3
					} else {
						answer += 7
					}
				}
			}
		} else {
			// Must be 1 or 4.
			if strings.ContainsRune(outp, segs[3]) {
				answer += 4
			} else {
				answer += 1
			}

		}
	}
	// fmt.Printf(" %c\n%c %c\n %c\n%c %c\n %c\n", segs[0], segs[1], segs[2], segs[3], segs[4], segs[5], segs[6])
	// fmt.Printf("%v %d\n", obs.outputs, answer)
	return answer
}

func Part1(input string) int {
	digits := make(map[int]int)
	observations := parseInput(input)
	for _, obs := range observations {
		for _, outp := range obs.outputs {
			digits[guessDigit(outp)]++
		}
	}
	return digits[1] + digits[4] + digits[7] + digits[8]
}

func Part2(input string) int {
	observations := parseInput(input)
	var sum int
	for _, obs := range observations {
		sum += deduceOutput(obs)
	}
	return sum
}
