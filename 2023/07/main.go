package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	InputFile string
}

const (
	HIGH_CARD int = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var CardStrength = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var CardStrengthJ = map[byte]int{
	'J': 0,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func IsDigit(char byte) bool {
	return 48 <= char && char <= 57
}

func GetStrength(label byte) int {
	//if IsDigit(label) {
	//	return int(label) - 48 // 48 = offset ascii table
	//}
	if val, ok := CardStrength[label]; ok {
		return val
	} else {
		return 0
	}
}

func GetStrengthJ(label byte) int {
	//if IsDigit(label) {
	//	return int(label) - 48 // 48 = offset ascii table
	//}
	if val, ok := CardStrengthJ[label]; ok {
		return val
	} else {
		return 0
	}
}

type Cards struct {
	cards     string // cards in given frequency
	cardsSort string // cards sorted for comparison
}

func (c *Cards) Value() int {
	// the order of function calls must be strictly adhered to in order to
	// correctly determine the value
	if c.is5k() {
		return FIVE_OF_A_KIND
	}
	if c.is4k() {
		return FOUR_OF_A_KIND
	}
	if c.isFH() {
		return FULL_HOUSE
	}
	if c.is3k() {
		return THREE_OF_A_KIND
	}
	if c.is2p() {
		return TWO_PAIR
	}
	if c.is1p() {
		return ONE_PAIR
	}
	return HIGH_CARD
}

// replace existing joker cards with the most highest cards
func ReplaceJoker(c string) string {
	cards := make(map[rune]int)

	for _, char := range c {
		if _, ok := cards[char]; ok {
			cards[char]++
		} else {
			cards[char] = 1
		}
	}

	if len(cards) > 1 {

		keys := make([]rune, 0, len(cards))
		for k := range cards {
			keys = append(keys, k)
		}

		sort.Slice(keys, func(i, j int) bool {
			if cards[keys[i]] == cards[keys[j]] {
				// check strength of cards with same counts
				return GetStrengthJ(byte(keys[i])) > GetStrengthJ(byte(keys[j]))
			}
			// check counts of card
			return cards[keys[i]] > cards[keys[j]]
		})

		if keys[0] == 'J' {
			c = strings.ReplaceAll(c, "J", string(keys[1]))
		} else {
			c = strings.ReplaceAll(c, "J", string(keys[0]))
		}
	}
	return c
}

func JokerExists(c *string) bool {
	cnt := 0
	for _, char := range *c {
		if char == 'J' {
			cnt++
		}
	}
	return cnt > 0
}

// five of a kind
// check if all five cards have the same label: [AAAAA]
func (c *Cards) is5k() bool {
	return c.cardsSort[0] == c.cardsSort[1] &&
		c.cardsSort[1] == c.cardsSort[2] &&
		c.cardsSort[2] == c.cardsSort[3] &&
		c.cardsSort[3] == c.cardsSort[4]
}

// four of a kind
// check if four cards have the same label and one card has a different label:
// [AA8AA]
func (c *Cards) is4k() bool {

	// checking [ a a a a x ]
	c1 := c.cardsSort[0] == c.cardsSort[1] &&
		c.cardsSort[1] == c.cardsSort[2] &&
		c.cardsSort[2] == c.cardsSort[3]

	// checking [ x a a a a ]
	c2 := c.cardsSort[1] == c.cardsSort[2] &&
		c.cardsSort[2] == c.cardsSort[3] &&
		c.cardsSort[3] == c.cardsSort[4]

	return c1 || c2
}

// full house
// check if three cards have the same label, and the remaining two cards share a
// different label: [23332]
func (c *Cards) isFH() bool {

	// checking [ a a a b b ]
	c1 := c.cardsSort[0] == c.cardsSort[1] &&
		c.cardsSort[1] == c.cardsSort[2] &&
		c.cardsSort[3] == c.cardsSort[4]

	// checking [ b b a a a ]
	c2 := c.cardsSort[0] == c.cardsSort[1] &&
		c.cardsSort[2] == c.cardsSort[3] &&
		c.cardsSort[3] == c.cardsSort[4]

	return c1 || c2
}

// three of a kind
// check if three cards have the same label, and the remaining two cards are
// each different from any other card in the hand: [TTT98]
func (c *Cards) is3k() bool {

	// checking [ a a a x y ]
	c1 := c.cardsSort[0] == c.cardsSort[1] && c.cardsSort[1] == c.cardsSort[2]

	// checking [ x a a a y ]
	c2 := c.cardsSort[1] == c.cardsSort[2] && c.cardsSort[2] == c.cardsSort[3]

	// checking [ x y a a a ]
	c3 := c.cardsSort[2] == c.cardsSort[3] && c.cardsSort[3] == c.cardsSort[4]

	return c1 || c2 || c3
}

// two pair
// check if two cards share one label, two other cards share a second label, and
// the remaining card has a third label: [23432]
func (c *Cards) is2p() bool {

	// checking [ a a b b x ]
	c1 := c.cardsSort[0] == c.cardsSort[1] && c.cardsSort[2] == c.cardsSort[3]

	// checking [ a a x b b ]
	c2 := c.cardsSort[0] == c.cardsSort[1] && c.cardsSort[3] == c.cardsSort[4]

	// checking [ x a a b b ]
	c3 := c.cardsSort[1] == c.cardsSort[2] && c.cardsSort[3] == c.cardsSort[4]

	return c1 || c2 || c3
}

// one pair
// check if wo cards share one label, and the other three cards have a different
// label from the pair and each other: [A23A4]
func (c *Cards) is1p() bool {

	// checking [ a a x y z ]
	c1 := c.cardsSort[0] == c.cardsSort[1]

	// checking [ x a a y z ]
	c2 := c.cardsSort[1] == c.cardsSort[2]

	// checking [ x y a a z ]
	c3 := c.cardsSort[2] == c.cardsSort[3]

	// checking [ x y z a a ]
	c4 := c.cardsSort[3] == c.cardsSort[4]

	return c1 || c2 || c3 || c4
}

type CardSet struct {
	cards Cards
	value int
	bid   int
}

type CardSets []CardSet

func NewCardSet(cards Cards, bid int) CardSet {
	return CardSet{
		cards,
		cards.Value(),
		bid,
	}
}

func NewCards(c string, joker bool) Cards {
	c0 := make([]string, 0, len(c))
	if joker && JokerExists(&c) {
		c0 = strings.Split(ReplaceJoker(c), "")
	} else {
		c0 = strings.Split(c, "")
	}
	sort.Strings(c0)
	return Cards{c, strings.Join(c0, "")}
}

func MultiplyRankBid(cs *CardSets) int {
	sum := 0
	for i := range *cs {
		sum += (i + 1) * (*cs)[i].bid
	}
	return sum
}

var cfg Config

func ParseInput(line *string, joker bool) CardSet {
	data := strings.Split(*line, " ")
	bid, _ := strconv.Atoi(data[1])
	return NewCardSet(NewCards(data[0], joker), bid)
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

		scanner := bufio.NewScanner(file)

		var cs, csJ CardSets // cardsets part 1

		for scanner.Scan() {
			line := scanner.Text()
			cs = append(cs, ParseInput(&line, false))
			csJ = append(csJ, ParseInput(&line, true))
		}

		// rank cardsets by value
		sort.Slice(cs, func(i, j int) bool {
			if cs[i].value == cs[j].value {
				for k := 0; k < 5; k++ {
					s1 := GetStrength(cs[i].cards.cards[k])
					s2 := GetStrength(cs[j].cards.cards[k])
					if s1 == s2 {
						continue
					}
					if s1 < s2 {
						return true
					} else {
						return false
					}
				}
			} else {
				return cs[i].value < cs[j].value
			}
			return false
		})
		// rank cardsets by value
		sort.Slice(csJ, func(i, j int) bool {
			if csJ[i].value == csJ[j].value {
				for k := 0; k < 5; k++ {
					s1 := GetStrengthJ(csJ[i].cards.cards[k])
					s2 := GetStrengthJ(csJ[j].cards.cards[k])
					if s1 == s2 {
						continue
					}
					if s1 < s2 {
						return true
					} else {
						return false
					}
				}
			} else {
				return csJ[i].value < csJ[j].value
			}
			return false
		})

		fmt.Println("(part 1) multiplied bids with ranks:", MultiplyRankBid(&cs))
		fmt.Println("(part 2) multiplied bids with ranks:", MultiplyRankBid(&csJ))

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
