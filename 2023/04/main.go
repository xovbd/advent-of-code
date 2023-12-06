package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Config struct {
	InputFile string
}

var cfg Config

type Card struct {
	id        int
	numsWin   []string
	numsCheck []string
	matches   int
	points    int
	quantity  int
}

func New(id int, numsWin, numsCheck []string) Card {

	// determine matches
	matches := 0
	for i := range numsCheck {
		if slices.Contains(numsWin, numsCheck[i]) {
			matches += 1
		}
	}

	// calculate points
	points := 0
	if matches > 0 {
		points = int(math.Pow(2, float64(matches-1)))
	}

	return Card{id, numsWin, numsCheck, matches, points, 1}
}

func (c Card) GetPoints() int {
	return c.points
}

func (c Card) GetMatches() int {
	return c.matches
}

func (c Card) GetQuantity() int {
	return c.quantity
}

type Stack struct {
	cards []Card
}

func (s *Stack) AppendCard(c Card) {
	s.cards = append(s.cards, c)
}

func (s *Stack) GetPoints() int {
	points := 0
	for i := range s.cards {
		points += s.cards[i].GetPoints()
	}
	return points
}

func (s *Stack) GetQuantity() int {
	quantity := 0
	for i := range s.cards {
		quantity += s.cards[i].GetQuantity()
	}
	return quantity
}

func DetermineCardCopies(s *Stack) {
	for i := range s.cards {
		for j := i; j < i+s.cards[i].GetMatches() && j < len(s.cards)-1; j++ {
			s.cards[j+1].quantity += s.cards[i].GetQuantity()
		}
	}
}

func ParseInput(line string) Card {
	var re = regexp.MustCompile(`\w+ +(\d+):([ \d]+) \|([ \d]+)`)
	var reNums = regexp.MustCompile(`\d+`)

	matches := re.FindAllStringSubmatch(line, -1)[0]
	cardId, _ := strconv.Atoi(matches[1])
	numsWin := reNums.FindAllString(matches[2], -1)
	numsCheck := reNums.FindAllString(matches[3], -1)

	return New(cardId, numsWin, numsCheck)
}

func init() {
	flag.StringVar(&cfg.InputFile, "in", "", "input filename")
}

func main() {

	flag.Parse()

	// check passed arguments
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	if !flagset["in"] {
		fmt.Println("you have to pass all arguments:")
		flag.PrintDefaults()
	} else {

		file, err := os.Open(cfg.InputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		cards := Stack{}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			cards.AppendCard(ParseInput(line))
		}

		DetermineCardCopies(&cards)

		fmt.Println("(part 1) sum of points:", cards.GetPoints())
		fmt.Println("(part 2) sum of cards :", cards.GetQuantity())

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
