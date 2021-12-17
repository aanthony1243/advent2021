package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct{
	x int
	y int
}

func getPairs(line string) (pair, pair) {
	ab := strings.Split(line, "->")
	a := strings.Split(strings.TrimSpace(ab[0]), ",")
	b := strings.Split(strings.TrimSpace(ab[1]), ",")

	x1, err := strconv.Atoi(a[0])
	if err != nil {
		return pair{},pair{}
	}
	y1, err := strconv.Atoi(a[1])
	if err != nil {
		return pair{},pair{}
	}
	x2, err := strconv.Atoi(b[0])
	if err != nil {
		return pair{},pair{}
	}
	y2, err := strconv.Atoi(b[1])
	if err != nil {
		return pair{},pair{}
	}

	return pair{x:x1, y: y1}, pair{x:x2, y:y2}

}


func main() {
	f, err := os.Open("5a-large.txt")
	if err != nil {
		return
	}

	hits := make(map[pair]int)

	scanner := bufio.NewScanner(f)
	totalHits := 0
	for scanner.Scan() {
		p1, p2 := getPairs(scanner.Text())
		if p1.x == p2.x {
			if p1.y > p2.y {
				p1.y, p2.y = p2.y, p1.y
			}
			for p1.y <= p2.y {
				prev, ok := hits[p1]
				if ok {
					hits[p1] = prev + 1
					if prev + 1 == 2 {
						totalHits += 1
					}
				} else {
					hits[p1] = 1
				}
				p1.y += 1
			}
		}else if p1.y == p2.y {
			if p1.x > p2.x {
				p1.x, p2.x = p2.x, p1.x
			}
			for p1.x <= p2.x {
				prev, ok := hits[p1]
				if ok {
					hits[p1] = prev + 1
					if prev + 1 == 2 {
						totalHits += 1
					}
				} else {
					hits[p1] = 1
				}
				p1.x += 1
			}
		} else {
			// diagonal case.  We will only work left -- right
			if p1.x > p2.x {
				p1,p2 = p2,p1
			}
			// next: do left-up
			if p1.y <= p2.y {
				for p1.x <= p2.x && p1.y <= p2.y {
					prev, ok := hits[p1]
					if ok {
						hits[p1] = prev + 1
						if prev + 1 == 2 {
							totalHits += 1
						}
					} else {
						hits[p1] = 1
					}
					p1.x += 1
					p1.y += 1
				}

			} else {
				// do left-down
				for p1.x <= p2.x && p1.y >= p2.y {
					prev, ok := hits[p1]
					if ok {
						hits[p1] = prev + 1
						if prev + 1 == 2 {
							totalHits += 1
						}
					} else {
						hits[p1] = 1
					}
					p1.x += 1
					p1.y -= 1
				}

			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Print("Final Count: ")
	fmt.Println(totalHits)

}
