package day5

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
)

type Day5 struct {
	initialSeeds []int
	mapperGroups [][]mapper
	seedRanges   []simpleRange
}

type simpleRange struct {
	source int
	length int
}

type mapper struct {
	source      int
	destination int
	length      int
}

func (day *Day5) SetInput(s string) {
	lines := helpers.GetLines(s)

	tkr := helpers.NewTokeniser(lines[0])
	tokens := tkr.Tokenise()
	st := helpers.NewStream[helpers.Token](tokens)
	st.NextAssert(func(t helpers.Token) bool {
		return t.Val == "seeds"
	})
	st.NextAssert(func(t helpers.Token) bool {
		return t.T == "colon"
	})
	for st.HasNext() {
		t := st.NextAssert(func(t helpers.Token) bool {
			return t.T == "int"
		})
		d, err := strconv.Atoi(t.Val)
		if err != nil {
			panic(err)
		}
		day.initialSeeds = append(day.initialSeeds, d)
	}
	day.seedRanges = make([]simpleRange, 0)
	for i := 0; i < len(day.initialSeeds); i += 2 {
		rng := simpleRange{
			day.initialSeeds[i],
			day.initialSeeds[i+1],
		}
		day.seedRanges = append(day.seedRanges, rng)
	}

	mappers := make([][]mapper, 0)
	mapperGroup := make([]mapper, 0)
	for _, line := range lines[1:] {
		tkr = helpers.NewTokeniser(line)
		tokens = tkr.Tokenise()
		st := helpers.NewStream[helpers.Token](tokens)

		// Blank line
		if len(tokens) == 0 {
			continue
		} else if st.Peek().T == "keyword" {
			mappers = append(mappers, mapperGroup)
			mapperGroup = make([]mapper, 0)
		} else {
			mpr := mapper{}
			dt := st.NextAssert(func(t helpers.Token) bool {
				return t.T == "int"
			})
			d, err := strconv.Atoi(dt.Val)
			if err != nil {
				panic(err)
			}
			mpr.destination = d

			srct := st.NextAssert(func(t helpers.Token) bool {
				return t.T == "int"
			})
			d, err = strconv.Atoi(srct.Val)
			if err != nil {
				panic(err)
			}
			mpr.source = d

			lt := st.NextAssert(func(t helpers.Token) bool {
				return t.T == "int"
			})
			d, err = strconv.Atoi(lt.Val)
			if err != nil {
				panic(err)
			}
			mpr.length = d

			mapperGroup = append(mapperGroup, mpr)
		}
	}
	mappers = append(mappers, mapperGroup)
	day.mapperGroups = mappers
}

func mapGroup(i int, mapGroup []mapper) int {
	for _, mapper := range mapGroup {
		if i >= mapper.source && i < mapper.source+mapper.length {
			dest := mapper.destination + (i - mapper.source)
			return dest
		}
	}
	return i
}

func mapGroupReverse(i int, mapGroup []mapper) int {
	for _, mapper := range mapGroup {
		if i >= mapper.destination && i < mapper.destination + mapper.length {
			dest := mapper.source + (i - mapper.destination)
			return dest
		}
	}
	return i
}

func (day Day5) mapSeedToLocation(seed int) int {
	for _, mapperGroup := range day.mapperGroups {
		seed = mapGroup(seed, mapperGroup)
	}
	return seed
}

func (day Day5) mapLocationToSeed(seed int) int {
	for i := len(day.mapperGroups) - 1; i >= 0; i-- {
		seed = mapGroupReverse(seed, day.mapperGroups[i])
	}
	return seed
}

func (day Day5) SolvePart1() string {
	least := 0
	for _, seed := range day.initialSeeds {
		location := day.mapSeedToLocation(seed)
		if least == 0 || location < least {
			least = location
		}
	}
	return fmt.Sprint(least)
}

func (day Day5) SolvePart2() string {
	location := 1
	for {
		seed := day.mapLocationToSeed(location)
		for _, rng := range day.seedRanges {
			if seed >= rng.source && seed < rng.source + rng.length {
				return fmt.Sprint(location)
			}
		}
		location++
	}
}
