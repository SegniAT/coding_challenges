package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	seeds := []int{}
	seedsStr := strings.Split(strings.Split(input[0], ": ")[1], " ")
	for _, seedStr := range seedsStr {
		seedInt, err := strconv.Atoi(seedStr)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seedInt)
	}

	seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight := [][]int{}, [][]int{}, [][]int{}, [][]int{}
	lightToTemprature, temperatureToHumidity, humidityToLocation := [][]int{}, [][]int{}, [][]int{}

	context := ""
	for ind, line := range input {
		if ind == 0 || ind == 1 {
			continue
		}

		if line == "" {
			context = ""
			continue
		}

		if line[len(line)-1] == ':' {
			context = line
			continue
		}

		switch context {
		case "seed-to-soil map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			seedToSoil = append(seedToSoil, mapping)
		case "soil-to-fertilizer map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			soilToFertilizer = append(soilToFertilizer, mapping)
		case "fertilizer-to-water map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			fertilizerToWater = append(fertilizerToWater, mapping)
		case "water-to-light map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			waterToLight = append(waterToLight, mapping)
		case "light-to-temperature map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			lightToTemprature = append(lightToTemprature, mapping)
		case "temperature-to-humidity map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			temperatureToHumidity = append(temperatureToHumidity, mapping)
		case "humidity-to-location map:":
			mapping, err := mapNums(line)
			if err != nil {
				panic(err)
			}
			humidityToLocation = append(humidityToLocation, mapping)
		}
	}

	fmt.Println("seeds: ", seeds)
	fmt.Println("seed to soil: ", seedToSoil)
	fmt.Println("soil to fertilizer: ", soilToFertilizer)
	fmt.Println("fertilizer to water: ", fertilizerToWater)
	fmt.Println("water to light: ", waterToLight)
	fmt.Println("light to temperature: ", lightToTemprature)
	fmt.Println("temperature to humidity: ", temperatureToHumidity)
	fmt.Println("humidity to location: ", humidityToLocation)

	seedLocations := make([]int, len(seeds))
	fmt.Println()

	for seedInd, seed := range seeds {

		// fmt.Printf("Seed (%d): %d\n", seedInd, seed)

		// soil
		soil := -1
		found := false
		for _, sts := range seedToSoil {
			soil, found = getDestination(seed, sts)
			if found {
				break
			}
		}
		// fmt.Println("soil: ", soil)

		// fertilizer
		fertilizer := -1
		found = false
		for _, stf := range soilToFertilizer {
			fertilizer, found = getDestination(soil, stf)
			if found {
				break
			}
		}
		// fmt.Println("fertilizer: ", fertilizer)

		// water
		water := -1
		found = false
		for _, ftw := range fertilizerToWater {
			water, found = getDestination(fertilizer, ftw)
			if found {
				break
			}
		}
		// fmt.Println("water: ", water)

		// light
		light := -1
		found = false
		for _, wtl := range waterToLight {
			light, found = getDestination(water, wtl)
			if found {
				break
			}
		}
		// fmt.Println("light: ", light)

		// temperature
		temperature := -1
		found = false
		for _, ltt := range lightToTemprature {
			temperature, found = getDestination(light, ltt)
			if found {
				break
			}
		}
		// fmt.Println("temperature: ", temperature)

		// humidity
		humidity := -1
		found = false
		for _, tth := range temperatureToHumidity {
			humidity, found = getDestination(temperature, tth)
			if found {
				break
			}
		}
		// fmt.Println("humidity: ", humidity)

		// location
		location := -1
		found = false
		for _, htl := range humidityToLocation {
			location, found = getDestination(humidity, htl)
			if found {
				break
			}
		}
		// fmt.Println("location: ", location)
		// fmt.Println()

		seedLocations[seedInd] = location
	}

	minSeedLocation := math.MaxInt

	for _, seedLocation := range seedLocations {
		if seedLocation < minSeedLocation {
			minSeedLocation = seedLocation
		}
	}

	fmt.Println("Seed locations: ", seedLocations)
	fmt.Println("Min seed location: ", minSeedLocation)
	// Attempt 1 : 910845529
}

func getDestination(source int, mapping []int) (int, bool) {
	destRangeStart, srcRangeStart, rangeLength := mapping[0], mapping[1], mapping[2]
	if srcRangeStart <= source && source <= srcRangeStart+rangeLength-1 {
		distance := source - srcRangeStart
		return destRangeStart + distance, true
	}

	return source, false
}

func mapNums(mapping string) ([]int, error) {
	mappingNums := strings.Split(mapping, " ")
	out := make([]int, 3)

	for ind, mappingNum := range mappingNums {
		num, err := strconv.Atoi(mappingNum)
		if err != nil {
			return nil, err
		}

		out[ind] = num
	}

	return out, nil
}

func readInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	out := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, nil
}
