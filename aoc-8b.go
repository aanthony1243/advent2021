package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)


func sortKey(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func main() {
	f, err := os.Open("8a-large.txt")
	if err != nil {
		return
	}


	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		markedVals := make(map[string]bool)
		knownVals := make(map[string]string)
		digits := make(map[string]int)
		sd := strings.Split(scanner.Text(), "|")
		signal, display := strings.Fields(sd[0]), strings.Fields(sd[1])
		sort.Slice(signal, func(i, j int) bool {
			return len(signal[i]) < len(signal[j])
		})

		fmt.Println(signal)
		fmt.Println(display)

		// template for conversation below
		//   A
		// B   C
		//   D
		// E   F
		//   G

		// 0: 8 is easy but doesn't tell us anything
		digits[sortKey(signal[9])] = 8

		// 1: find segments for 1, though order is unknown.
		sig := signal[0]
		markedVals[string(sig[0])] = true
		markedVals[string(sig[1])] = true
		digits[sortKey(sig)] = 1

		// 2: given segments for 1, find top of 7 by removing segments for 1.  (A) is known
		sig = signal[1]
		for _,x := range sig {
			_, ok := markedVals[string(x)]
			if !ok {
				knownVals[string(x)] = "A"
			}
			// don't mark known values
		}
		digits[sortKey(sig)] = 7

		// 3: mark all the parts of 4.  markedVals is now precisely the unordered segments of 4
		sig = signal[2]
		for _,x := range sig {
			markedVals[string(x)] = true
		}
		digits[sortKey(sig)] = 4

		// 4: 9 is the only 6 digit that covers 4, and Now that we know all the digits of 1, 4 and 7 find the bottom of 9 (G) is known.
		// comparing 9 to 8, (E) is known.
		for _, sig := range signal[6:9] {
			// the known + marked values should cover all of 9 but the bottom.  6 and 0 will miss 2 segments
			segcnt := 0
			for _, x := range sig {
				sx := string(x)
				_, ok := markedVals[sx]
				_, ok2 := knownVals[sx]
				if !(ok || ok2) {
					segcnt += 1
				}
			}
			if segcnt == 1 {
				for _, x := range sig {
					sx := string(x)
					_, ok := markedVals[sx]
					_, ok2 := knownVals[sx]
					if !(ok || ok2) {
						knownVals[string(x)] = "G"
						digits[sortKey(sig)] = 9
					}
				}
				eight := signal[9]
				for _, x := range eight {
					_, ok := markedVals[string(x)]
					_, ok2 := knownVals[string(x)]
					if !(ok || ok2) {
						knownVals[string(x)] = "E"
					}
				}
			}
		}

		// 5: 6 is the only digit that doesn't cover 1
		for _, sig := range signal[6:9] {
			one := signal[0]
			found0, found1 := false, false
			for _, x := range sig {
				if string(one[0]) == string(x) {
					found0 = true
				}
				if string(one[1]) == string(x) {
					found1 = true
				}
			}
			if !(found0 && found1) {
				digits[sortKey(sig)] = 6
				if found0{
					knownVals[string(one[0])] = "F"
					knownVals[string(one[1])] = "C"
				}
				if found1 {
					knownVals[string(one[1])] = "F"
					knownVals[string(one[0])] = "C"
				}
				break
			}
		}

		// 6: Now that we know 9 and 6, we know which is 0.
		for _, sig := range signal[6:9] {
			_, ok := digits[sortKey(sig)]
			if !ok {
				digits[sortKey(sig)] = 0
				// and the one not found segment will tell us D
				missing := ""
				for _, x := range []string{"a","b","c","d","e","f","g"} {

					for _,s := range sig {
						xs := string(s)
						if x != xs {
							missing = x
						} else {
							missing = ""
							break
						}
					}
					if missing != "" {
						knownVals[missing] = "D"
						break
					}
				}


				// and now we can find B since it's the last unknown segment
				missing = ""
				for _, x := range []string{"a","b","c","d","e","f","g"} {
					for k := range knownVals {
						if x != k {
							missing = x
						} else {
							missing = ""
							break
						}
					}
					if missing != "" {
						knownVals[missing] = "B"
						break
					}
				}

			}
		}

		// since we know all the segment mappings, we can find 2, 3, 5
		for _, sig := range signal[3:6]{
			// 2 is A,C,D,E,G
			found := true
			for _,x := range sig {
				xs := string(x)
				v := knownVals[xs]
				if v == "B" || v == "F" {
					found = false
					break
				}
			}
			if found {
				digits[sortKey(sig)] = 2
				continue
			}
			found = true
			for _,x := range sig {
				xs := string(x)
				v := knownVals[xs]
				if v == "B" || v == "E" {
					found = false
					break
				}
			}
			if found {
				digits[sortKey(sig)] = 3
				continue
			}

			found = true
			for _,x := range sig {
				xs := string(x)
				v := knownVals[xs]
				if v == "C" || v == "E" {
					found = false
					break
				}

			}
			if found {
				digits[sortKey(sig)] = 5
				continue
			}
		}

		// yay now we have all the digit mappings, let's print them and then interpret the display
		fmt.Println(digits)
		final := digits[sortKey(display[0])]*1000 + digits[sortKey(display[1])]*100 + digits[sortKey(display[2])]*10 + digits[sortKey(display[3])]
		total += final
		fmt.Println(final)


	}
	fmt.Printf("Total: %d\n", total)
}
