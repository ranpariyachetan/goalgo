package advent2021

import (
	"fmt"
	"strconv"
)

func TestGetPowerConsumption() {
	logs := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	r := getPowerConsumption(logs)

	fmt.Println(r)
}

func getPowerConsumption(logs []string) int64 {
	counter := make(map[rune]int)
	gamma := ""
	epsilon := ""

	for i := 0; i < len(logs[0]); i++ {
		counter['0'] = 0
		counter['1'] = 0

		for _, log := range logs {
			counter[rune(log[i])]++
		}
		if counter['1'] > counter['0'] {
			gamma += string('1')
			epsilon += string('0')
		} else {
			gamma += string('0')
			epsilon += string('1')
		}

	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return g * e
}
