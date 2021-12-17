package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readBoard(scanner *bufio.Scanner) (bool, [][]string) {
	i:=0
	numbers := make([][]string, 5)
	for ; i < 5 && scanner.Scan(); i++ {
		numbers[i] = strings.Fields(scanner.Text())
	}
    if i == 5 {
    	scanner.Scan()
    	return true, numbers
    }
    return false, numbers
}

func markBoard(calls []string, numbers [][]string) (int,string) {
	for ind, call := range calls {
		for i, row := range numbers {
			for j, square := range row {
				if square == call {
					//rowWin += 0
					//n, err :=  strconv.Atoi(square)
					//if err != nil {
					//	return -1
					//}
					//rowScore += n
					numbers[i][j] = "-" + numbers[i][j]
				}
				//if rowWin == 4 {
				//	return rowScore
				//}
			}
			// 1. check row for win
			rowWin := 0
			for _, square := range row{
				if strings.HasPrefix(square, "-") {
					rowWin += 1
				}
			}
			if rowWin == 5 {
				return ind, call
			}
			// 2.  on last row, check cols for win
			if i == 4 {
				for j, square := range row {
					if strings.HasPrefix(square, "-") {
						// then this column might be a winner
						colWin := 0
						for r := 0; r < 5; r++ {
							if strings.HasPrefix(numbers[r][j], "-") {
								colWin += 1
							}
						}
						if colWin == 5 {
							return ind, call
						}
					}
				}
			}
		}
	}
	return -1, ""
}

func scoreBoard(numbers [][]string) (int) {
	score := 0
	for _, row := range numbers {
		for _, square := range row {
			if strings.HasPrefix(square, "-") {
				continue
			}
			n, err := strconv.Atoi(square)
			if err != nil {
				return -1
			}
			score += n
		}
	}
	return score
}


func main() {
	f, err := os.Open("4a-large.txt")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	calls := strings.Split(scanner.Text(), ",")
	fmt.Println(calls)
	scanner.Scan()

	var bestBoard [][]string
	// a
    // bestInd := 1000000000
	// b
    bestInd := 0
	bestCall := ""
	for {
		ok, board := readBoard(scanner)
		if !ok {
			break
		}
		fmt.Println(board)
		ind, call := markBoard(calls, board)
		// a
        // if ind < bestInd {
		// b
		if ind >= bestInd {
			bestInd = ind
			bestBoard = board
			bestCall = call
		}

	}

	fmt.Println("Best Board: ")
	fmt.Println(bestBoard)
	fmt.Print("Final Call = ")
	fmt.Println(bestCall)
	fmt.Print("Index: ")
	fmt.Println(bestInd)
	c, err := strconv.Atoi(bestCall)
	if err != nil {
		return
	}
	score := scoreBoard(bestBoard)
	fmt.Print("BoardSum: ")
	fmt.Println(score)
	fmt.Print("Puzzle Answer: ")
	fmt.Println(score*c)

}
