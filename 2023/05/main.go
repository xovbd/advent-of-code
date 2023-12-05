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
	InputFile string
}

var cfg Config

type Rule struct {
	dst uint64
	src uint64
	len uint64
}

func ExtractSeeds(line *string) []uint64 {
	var seeds []uint64
	var re = regexp.MustCompile(`\d+`)
	matches := re.FindAllString(*line, -1)
	for i := range matches {
		seed, _ := strconv.Atoi(matches[i])
		seeds = append(seeds, uint64(seed))
	}
	return seeds
}

func ExtractTopic(line *string) string {
	return strings.Split(*line, " ")[0]
}

func ExtractRule(line *string) Rule {
	rules := strings.Split(*line, " ")
	dst, _ := strconv.Atoi(rules[0])
	src, _ := strconv.Atoi(rules[1])
	len, _ := strconv.Atoi(rules[2])
	return Rule{uint64(dst), uint64(src), uint64(len)}
}

func GetValue(v uint64, rs []Rule) uint64 {
	for _, r := range rs {
		if r.src <= v && v <= r.src+r.len {
			return r.dst + (v - r.src)
		}
	}
	return v
}

func DetermineLocation(seed uint64, rules map[string][]Rule) uint64 {
	soil := GetValue(seed, rules["seed-to-soil"])
	fertilizer := GetValue(soil, rules["soil-to-fertilizer"])
	water := GetValue(fertilizer, rules["fertilizer-to-water"])
	light := GetValue(water, rules["water-to-light"])
	temperature := GetValue(light, rules["light-to-temperature"])
	humidity := GetValue(temperature, rules["temperature-to-humidity"])
	location := GetValue(humidity, rules["humidity-to-location"])
	//fmt.Printf("seed<%d> soil<%d> fertilizer<%d> water<%d> light<%d> temperature<%d> humidity<%d> location<%d>\n",
	//	seed, soil, fertilizer, water, light, temperature, humidity, location)
	return location
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

		rules := make(map[string][]Rule)
		var seeds []uint64
		var topic string

		scanner := bufio.NewScanner(file)

		firstLine := true
		for scanner.Scan() {
			line := scanner.Text()

			// skip blank lines
			if len(line) == 0 {
				continue
			}

			// determine seeds
			if firstLine {
				firstLine = false
				seeds = ExtractSeeds(&line)
				continue
			}

			// determine mapping-topic
			if 97 <= line[0] && line[0] <= 122 { // 97=a 122=z
				topic = ExtractTopic(&line)
			} else { // determine rule
				rules[topic] = append(rules[topic], ExtractRule(&line))
			}
		}

		//fmt.Println("----- PART 1 -----")
		shortestLocationP1 := ^uint64(0)
		for _, seed := range seeds {
			location := DetermineLocation(seed, rules)
			if location < shortestLocationP1 {
				shortestLocationP1 = location
			}
		}

		//fmt.Println("----- PART 2 -----")
		shortestLocationP2 := ^uint64(0)
		for i := 0; i < len(seeds); i += 2 {
			for seed := seeds[i]; seed < (seeds[i] + seeds[i+1]); seed++ {
				location := DetermineLocation(seed, rules)
				if location < shortestLocationP2 {
					shortestLocationP2 = location
				}
			}
		}

		fmt.Println("(part 1) shortest location:", shortestLocationP1)
		fmt.Println("(part 2) shortest location:", shortestLocationP2)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
