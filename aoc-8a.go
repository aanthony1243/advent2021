package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("8a-small.txt")
	if err != nil {
		return
	}


	scanner := bufio.NewScanner(f)

	totalFound := 0
	for scanner.Scan() {
		knownVals := make(map[string]string)
		sd := strings.Split(scanner.Text(), "|")
		signal, display := strings.Fields(sd[0]), strings.Fields(sd[1])
		sort.Slice(signal, func(i, j int) bool {
			return len(signal[i]) < len(signal[j])
		})
		sort.Slice(display, func(i, j int) bool {
			return len(display[i]) < len(display[j])
		})

		fmt.Println(signal)
		fmt.Println(display)

		// we can map 1, 4, 7, 8 immediately
		// shortest signal is 1, at pos 0


		for _, dis := range display {
			// is a 1 displayed?
			sig := signal[0]
			if len(sig) == len(dis) {
				knownVals[string(sig[0])] = string(dis[0])
				knownVals[string(sig[1])] = string(dis[1])
				totalFound += 1
			}
			// next shortest signal is 7, at pos 1
			sig = signal[1]
			if len(sig) == len(dis) {
				knownVals[string(sig[0])] = string(dis[0])
				knownVals[string(sig[1])] = string(dis[1])
				knownVals[string(sig[2])] = string(dis[2])
				totalFound += 1
			}

			// now 4, at pos 2
			sig = signal[2]
			if len(sig) == len(dis) {
				knownVals[string(sig[0])] = string(dis[0])
				knownVals[string(sig[1])] = string(dis[1])
				knownVals[string(sig[2])] = string(dis[2])
				knownVals[string(sig[3])] = string(dis[3])
				totalFound += 1
			}

			// and 8, at pos 9
			sig = signal[9]
			if len(sig) == len(dis) {
				knownVals[string(sig[0])] = string(dis[0])
				knownVals[string(sig[1])] = string(dis[1])
				knownVals[string(sig[2])] = string(dis[2])
				knownVals[string(sig[3])] = string(dis[3])
				knownVals[string(sig[4])] = string(dis[4])
				knownVals[string(sig[5])] = string(dis[5])
				knownVals[string(sig[6])] = string(dis[6])
				totalFound += 1
			}
		}
	}
	fmt.Printf("final found digits: %d\n", totalFound)
}
