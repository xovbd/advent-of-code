package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var mapWordDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {

	args := os.Args

	if len(args) > 1 {
		inputFile := args[1]

		var sumLineValuesPart1 int = 0
		var sumLineValuesPart2 int = 0

		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			var digitsLinePart1 []int
			var digitsLinePart2 []int

			for i, char := range line {
				if char >= '0' && char <= '9' {
					digit := int(char - 48) // 48 = offset ascii table
					digitsLinePart1 = append(digitsLinePart1, digit)
					digitsLinePart2 = append(digitsLinePart2, digit)
				} else {
					for word, digit := range mapWordDigit {
						if strings.HasPrefix(line[i:], word) {
							digitsLinePart2 = append(digitsLinePart2, digit)
							break
						}
					}
				}
			}

			lineValuePart1 := digitsLinePart1[0]*10 + digitsLinePart1[len(digitsLinePart1)-1]
			sumLineValuesPart1 += lineValuePart1
			lineValuePart2 := digitsLinePart2[0]*10 + digitsLinePart2[len(digitsLinePart2)-1]
			sumLineValuesPart2 += lineValuePart2
		}

		fmt.Println("sum (part1):", sumLineValuesPart1)
		fmt.Println("sum (part2):", sumLineValuesPart2)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal("pass the filename of the input file as an argument to run the program")
	}
}
