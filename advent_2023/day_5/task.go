package day_5

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	seedsText             = "seeds"
	seedToSoil            = "seed-to-soil"
	soilToFertilizer      = "soil-to-fertilizer"
	fertilizerToWater     = "fertilizer-to-water"
	waterToLight          = "water-to-light"
	lightToTemperature    = "light-to-temperature"
	temperatureToHumidity = "temperature-to-humidity"
	humidityToLocation    = "humidity-to-location"
)

var sections = []string{
	seedToSoil,
	soilToFertilizer,
	fertilizerToWater,
	waterToLight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation,
}

type mapper struct {
	a, b, r int
}

func Solver(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	m := make(map[string][]mapper)
	var seeds []int
	section := ""
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case strings.Contains(text, seedsText):
			seeds, err = readSeeds(text)
			if err != nil {
				return 0, err
			}
			continue
		case strings.Contains(text, seedToSoil):
			section = seedToSoil
			continue
		case strings.Contains(text, soilToFertilizer):
			section = soilToFertilizer
			continue
		case strings.Contains(text, fertilizerToWater):
			section = fertilizerToWater
			continue
		case strings.Contains(text, waterToLight):
			section = waterToLight
			continue
		case strings.Contains(text, lightToTemperature):
			section = lightToTemperature
			continue
		case strings.Contains(text, temperatureToHumidity):
			section = temperatureToHumidity
			continue
		case strings.Contains(text, humidityToLocation):
			section = humidityToLocation
			continue
		}
		if text == "" {
			continue
		}
		v, ok := m[section]
		if !ok || v == nil {
			v = make([]mapper, 0, 50)
			m[section] = v
		}
		v, err = readMapping(text, v)
		m[section] = v
		if err != nil {
			return 0, err
		}
	}

	for _, mappings := range m {
		sortMappings(mappings)
	}

	loc := math.MaxInt
	for _, seed := range seeds {
		candidate := calculateLocation(seed, m)
		if candidate < loc {
			loc = candidate
		}
	}

	return loc, nil
}

func Solver2(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	m := make(map[string][]mapper)
	var seeds []int
	section := ""
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case strings.Contains(text, seedsText):
			seeds, err = readSeeds(text)
			if err != nil {
				return 0, err
			}
			continue
		case strings.Contains(text, seedToSoil):
			section = seedToSoil
			continue
		case strings.Contains(text, soilToFertilizer):
			section = soilToFertilizer
			continue
		case strings.Contains(text, fertilizerToWater):
			section = fertilizerToWater
			continue
		case strings.Contains(text, waterToLight):
			section = waterToLight
			continue
		case strings.Contains(text, lightToTemperature):
			section = lightToTemperature
			continue
		case strings.Contains(text, temperatureToHumidity):
			section = temperatureToHumidity
			continue
		case strings.Contains(text, humidityToLocation):
			section = humidityToLocation
			continue
		}
		if text == "" {
			continue
		}
		v, ok := m[section]
		if !ok || v == nil {
			v = make([]mapper, 0, 50)
			m[section] = v
		}
		v, err = readMapping(text, v)
		m[section] = v
		if err != nil {
			return 0, err
		}
	}

	for _, mappings := range m {
		sortMappings(mappings)
	}

	loc := math.MaxInt
	for i := range seeds {
		if i%2 != 0 {
			continue
		}
		if i == len(seeds)-1 {
			break
		}
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			candidate := calculateLocation(j, m)
			if candidate < loc {
				loc = candidate
			}
		}
	}

	return loc, nil
}

func calculateLocation(seed int, m map[string][]mapper) int {
	val := seed
	// fmt.Printf("NEW SEED: %v\n", seed)
	for _, section := range sections {
		val = findValueBinarySearch(m[section], val)
	}
	return val
}

func readMapping(s string, mapping []mapper) ([]mapper, error) {
	var a, b, r int
	_, err := fmt.Sscanf(s, "%d %d %d", &a, &b, &r)
	if err != nil {
		return mapping, err
	}
	return append(mapping, mapper{a: a, b: b, r: r}), nil
}

func readSeeds(s string) ([]int, error) {
	sl := strings.Split(s, ":")
	if len(sl) != 2 {
		return nil, errors.New("cannot read seeds, no ':' present")
	}
	var seeds []int
	for _, candidate := range strings.Split(sl[1], " ") {
		if candidate == "" {
			continue
		}
		v, err := strconv.Atoi(candidate)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, v)
	}
	return seeds, nil
}

func sortMappings(m []mapper) {
	sort.SliceStable(m, func(i, j int) bool { return m[i].b < m[j].b })
}

func findValueBinarySearch(array []mapper, v int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		m := array[mid]
		if v >= m.b && v < m.b+m.r {
			return v - m.b + m.a
		}
		if v > m.b {
			low = mid + 1
			continue
		}
		high = mid - 1
	}
	return v
}
