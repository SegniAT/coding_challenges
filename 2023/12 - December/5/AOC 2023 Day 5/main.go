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

	input, err := readInput("sample_input.txt")
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

	seedsRanges := [][]int{}
	for i := 0; i < len(seeds); i += 2 {
		seedsRanges = append(seedsRanges, []int{seeds[i], seeds[i] + seeds[i+1] - 1})
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

	fmt.Println("seeds: ", seedsRanges)
	fmt.Println("seed to soil: ", seedToSoil)
	fmt.Println("soil to fertilizer: ", soilToFertilizer)
	fmt.Println("fertilizer to water: ", fertilizerToWater)
	fmt.Println("water to light: ", waterToLight)
	fmt.Println("light to temperature: ", lightToTemprature)
	fmt.Println("temperature to humidity: ", temperatureToHumidity)
	fmt.Println("humidity to location: ", humidityToLocation)
	fmt.Println()
	fmt.Println()

	// soil
	soil := modifyRange(seedsRanges, seedToSoil)
	fmt.Println("soil: ", soil)

	// fertilizer
	fertilizer := modifyRange(soil, soilToFertilizer)
	fmt.Println("fertilizer: ", fertilizer)

	// water
	water := modifyRange(fertilizer, fertilizerToWater)
	fmt.Println("water: ", water)

	// light
	light := modifyRange(water, waterToLight)
	fmt.Println("light: ", light)

	// temperature
	temperature := modifyRange(light, lightToTemprature)
	fmt.Println("temperature: ", temperature)

	// humidity
	humidity := modifyRange(temperature, temperatureToHumidity)
	fmt.Println("humidity: ", humidity)

	// location
	location := modifyRange(humidity, humidityToLocation)
	fmt.Println("location: ", location)

	minLocation := math.MaxInt

	for _, loc := range location {
		if loc[0] < minLocation {
			minLocation = loc[0]
		}
	}

	fmt.Println("Minimum Location: ", minLocation)
	// 57020831 - too low
}

func modifyRange(srcRanges [][]int, mappings [][]int) (destRanges [][]int) {
	for i := 0; i < len(srcRanges); i++ {
		srcRange := srcRanges[i]
		srcXok, srcYok := -1, -1
		mostRangeContained, mostRangeContainedInd := -1, -1
		for ind, mapping := range mappings {
			start, end := mapping[1], mapping[1]+mapping[2]-1

			if start <= srcRange[0] && srcRange[0] <= end {
				srcXok = ind
			}

			if start <= srcRange[1] && srcRange[1] <= end {
				srcYok = ind
			}

			if srcRange[0] < start && end < srcRange[1] {
				if end-start > mostRangeContained {
					mostRangeContainedInd = ind
					mostRangeContained = end - start
				}
			}

		}

		if srcXok != -1 {
			candidateMapping := mappings[srcXok]

			_, end := candidateMapping[1], candidateMapping[1]+candidateMapping[2]-1

			if srcRange[1] <= end {
				destRanges = append(destRanges, convRange(srcRange, candidateMapping))
			} else {
				srcRanges = append(srcRanges, []int{end + 1, srcRange[1]})
				destRanges = append(destRanges, convRange([]int{srcRange[0], end}, candidateMapping))
			}
		} else if srcYok != -1 {
			candidateMapping := mappings[srcYok]
			start, _ := candidateMapping[1], candidateMapping[1]+candidateMapping[2]-1

			srcRanges = append(srcRanges, convRange([]int{start, srcRange[1]}, candidateMapping))
			srcRanges = append(srcRanges, []int{srcRange[0], start - 1})
		}

		if srcXok == -1 && srcYok == -1 {
			if mostRangeContainedInd != -1 {
				candidateMapping := mappings[mostRangeContainedInd]
				start, end := candidateMapping[1], candidateMapping[1]+candidateMapping[2]-1

				destRanges = append(destRanges, []int{srcRange[0], start - 1})
				destRanges = append(destRanges, candidateMapping)
				destRanges = append(destRanges, []int{end + 1, srcRange[1]})

			} else {
				destRanges = append(destRanges, srcRange)
			}
		}

	}
	return
}

func convRange(src []int, mapping []int) (dest []int) {
	dest = make([]int, 2)
	dest[0], _ = getDestination(src[0], mapping)
	dest[1], _ = getDestination(src[1], mapping)
	dest[1]--
	return
}

func getDestination(source int, mapping []int) (int, bool) {
	destRangeStart, srcRangeStart, rangeLength := mapping[0], mapping[1], mapping[2]
	if srcRangeStart <= source && source <= srcRangeStart+rangeLength-1 {
		distance := source - srcRangeStart
		return destRangeStart + distance, true
	}

	return source, false
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
