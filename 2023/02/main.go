package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	Blue      int
	Green     int
	Red       int
	InputFile string
}

var cfg Config

var countsGame = make(map[int]map[string]int)

func init() {
	flag.StringVar(&cfg.InputFile, "in", "", "input filename")
	flag.IntVar(&cfg.Red, "r", 0, "available red cubes")
	flag.IntVar(&cfg.Green, "g", 0, "available green cubes")
	flag.IntVar(&cfg.Blue, "b", 0, "available blue cubes")
}

func main() {

	flag.Parse()

	// check passed arguments
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	if !flagset["r"] || !flagset["g"] || !flagset["b"] || !flagset["in"] {
		fmt.Println("you have to pass all arguments:")
		flag.PrintDefaults()
	} else {

		file, err := os.Open(cfg.InputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var availableCubes = map[string]int{
			"red":   cfg.Red,
			"green": cfg.Green,
			"blue":  cfg.Blue,
		}

		// regex to find color and count
		var re = regexp.MustCompile(`(?:(\d+)\s([a-z]+))+`)

		var possibleSum int = 0
		var powerSum int = 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			game := strings.Split(line, ": ")

			gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])

			countsGame[gameId] = map[string]int{
				"blue":  0,
				"green": 0,
				"red":   0,
			}

			var possibleRound = true

			for _, round := range strings.Split(game[1], "; ") {
				for _, info := range re.FindAllStringSubmatch(round, -1) {

					count, _ := strconv.Atoi(info[1])
					color := info[2]

					if count > availableCubes[color] {
						possibleRound = false
					}
					if count > countsGame[gameId][color] {
						countsGame[gameId][color] = count
					}
				}
			}

			if possibleRound {
				possibleSum += gameId
			}

			powerSum += countsGame[gameId]["red"] * countsGame[gameId]["green"] * countsGame[gameId]["blue"]

		}

		fmt.Println("(part 1) sum of possible rounds:", possibleSum)
		fmt.Println("(part 2)           sum of power:", powerSum)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
