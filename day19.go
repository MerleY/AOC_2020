package main

import (
	"aoc/input"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	sets           [][]int
	hardcoded      bool
	hardcodedValue string
}

func day19() {
	data := input.Load("19").ToStringArray()
	var data_rules []string
	var messages []string
	for i, l := range data {
		if l == "" {
			data_rules = data[:i]
			messages = data[i+1:]
			break
		}
	}

	// parse rules
	rules := make(map[int]*Rule)
	for _, rule := range data_rules {
		parts := strings.Split(rule, ":")
		numRule, _ := strconv.Atoi(parts[0])
		rules[numRule] = &Rule{hardcoded: false}

		if strings.Contains(parts[1], "\"") {
			rules[numRule].hardcoded = true
			rules[numRule].hardcodedValue = strings.Trim(parts[1], "\" ")
		} else {
			sets := strings.Split(parts[1], "|")
			rules[numRule].hardcoded = false
			rules[numRule].sets = make([][]int, len(sets))
			for i, set := range sets {
				trimedSet := strings.Trim(set, " ")
				trimedSets := strings.Split(trimedSet, " ")
				rules[numRule].sets[i] = make([]int, len(trimedSets))
				for j, v := range trimedSets {
					intV, _ := strconv.Atoi(v)
					rules[numRule].sets[i][j] = intV
				}
			}
		}
	}

	// part 1
	nbValid := 0
	for _, m := range messages {
		if matchRule(rules, 0, m) {
			nbValid++
		}
	}
	fmt.Printf("Star 1: %v\n", nbValid)
}

func matchRule(rules map[int]*Rule, ruleNumber int, message string) bool {
	rule0 := rules[ruleNumber]

}
