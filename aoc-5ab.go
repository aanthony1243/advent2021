package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.

They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for you to review. For example:

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends. In other words:

An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines overlap?
*/

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
