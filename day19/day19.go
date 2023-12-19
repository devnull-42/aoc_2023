package day19

import (
	"aoc/util"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadInput("day19.txt")
	result := partA(lines)
	fmt.Printf("partA: %d\n", result)

	result = partB(lines)
	fmt.Printf("partB: %d\n", result)
}

func partA(lines []string) int {
	ruleMap, parts := parseLines(lines)

	var accepted []Part
	for _, part := range parts {
		if applyRules("in", ruleMap, part) == "A" {
			accepted = append(accepted, part)
		}
	}

	var result int
	for _, part := range accepted {
		result += part.Value()
	}

	return result
}

func partB(lines []string) int {
	workflows := parseLinesB(lines)
	result := getCombinations(categoryMap, "in", workflows)

	return result
}

type Rule struct {
	comparison func(p Part) bool
	dest       string
}

type Part struct {
	x int
	m int
	a int
	s int
}

func (p Part) Value() int {
	return p.x + p.m + p.a + p.s
}

func applyRules(ruleName string, ruleMap map[string][]Rule, part Part) string {
	if ruleName == "A" || ruleName == "R" {
		return ruleName
	}
	for _, rule := range ruleMap[ruleName] {
		if rule.comparison(part) {
			return applyRules(rule.dest, ruleMap, part)
		}
	}
	return ""
}

func parseLines(lines []string) (map[string][]Rule, []Part) {
	ruleMap := make(map[string][]Rule)
	var parts []Part
	idx := slices.Index(lines, "")

	reRuleLine := regexp.MustCompile(`^(\w+){(.*)}$`)
	reRule := regexp.MustCompile(`^(\w+)([<>])(\d+):(\w+)$`)
	for _, line := range lines[:idx] {
		matches := reRuleLine.FindStringSubmatch(line)
		ruleNamme := matches[1]
		rules := strings.Split(matches[2], ",")
		for _, r := range rules {
			matches = reRule.FindStringSubmatch(r)
			if len(matches) > 0 {
				ruleMap[ruleNamme] = append(ruleMap[ruleNamme], buildRule(matches[1], matches[2], util.MustAtoi(matches[3]), matches[4]))
			} else {
				ruleMap[ruleNamme] = append(ruleMap[ruleNamme], buildRule("", "", 0, r))
			}
		}

	}

	rePart := regexp.MustCompile(`^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
	for _, line := range lines[idx+1:] {
		matches := rePart.FindStringSubmatch(line)
		if len(matches) > 4 {
			parts = append(parts, Part{
				x: util.MustAtoi(matches[1]),
				m: util.MustAtoi(matches[2]),
				a: util.MustAtoi(matches[3]),
				s: util.MustAtoi(matches[4]),
			})
		}

	}

	return ruleMap, parts
}

func buildRule(category, op string, val int, dest string) Rule {
	rule := new(Rule)
	rule.dest = dest
	switch category {
	case "x":
		switch op {
		case "<":
			rule.comparison = func(p Part) bool {
				return p.x < val
			}
		case ">":
			rule.comparison = func(p Part) bool {
				return p.x > val
			}
		}
	case "m":
		switch op {
		case "<":
			rule.comparison = func(p Part) bool {
				return p.m < val
			}
		case ">":
			rule.comparison = func(p Part) bool {
				return p.m > val
			}
		}
	case "a":
		switch op {
		case "<":
			rule.comparison = func(p Part) bool {
				return p.a < val
			}
		case ">":
			rule.comparison = func(p Part) bool {
				return p.a > val
			}
		}
	case "s":
		switch op {
		case "<":
			rule.comparison = func(p Part) bool {
				return p.s < val
			}
		case ">":
			rule.comparison = func(p Part) bool {
				return p.s > val
			}
		}
	default:
		rule.comparison = func(p Part) bool {
			return true
		}
	}
	return *rule
}

// part b functions

type RuleB struct {
	cat  string
	op   string
	val  int
	dest string
}

var categoryMap = map[string][2]int{
	"x": {1, 4000},
	"m": {1, 4000},
	"a": {1, 4000},
	"s": {1, 4000},
}

type Workflow struct {
	rules []RuleB
	final string
}

func parseLinesB(lines []string) map[string]Workflow {
	workflows := make(map[string]Workflow)
	idx := slices.Index(lines, "")

	reRuleLine := regexp.MustCompile(`^(\w+){(.*)}$`)
	reRule := regexp.MustCompile(`^(\w+)([<>])(\d+):(\w+)$`)
	for _, line := range lines[:idx] {
		var name, final string
		matches := reRuleLine.FindStringSubmatch(line)
		name = matches[1]
		ruleStrings := strings.Split(matches[2], ",")
		rules := make([]RuleB, 0)
		for _, r := range ruleStrings {
			matches = reRule.FindStringSubmatch(r)
			if len(matches) > 0 {
				rules = append(rules, RuleB{cat: matches[1], op: matches[2], val: util.MustAtoi(matches[3]), dest: matches[4]})
			} else {
				final = r
			}
		}
		workflows[name] = Workflow{rules: rules, final: final}
	}
	return workflows
}

func getCombinations(ranges map[string][2]int, workflowName string, workflows map[string]Workflow) int {
	result := 0
	// base case
	if workflowName == "R" {
		return 0
	}
	if workflowName == "A" {
		baseResult := 1
		for _, v := range ranges {
			baseResult *= (v[1] - v[0] + 1)
		}
		return baseResult
	}

	for _, r := range workflows[workflowName].rules {
		min, max := ranges[r.cat][0], ranges[r.cat][1]
		var pass, fail [2]int
		switch r.op {
		case "<":
			pass = [2]int{min, slices.Min([]int{r.val - 1, max})}
			fail = [2]int{slices.Max([]int{r.val, min}), max}
		case ">":
			pass = [2]int{slices.Max([]int{r.val + 1, min}), max}
			fail = [2]int{min, slices.Min([]int{r.val, max})}
		}

		if pass[0] <= pass[1] {
			rangesCopy := make(map[string][2]int)
			maps.Copy(rangesCopy, ranges)
			rangesCopy[r.cat] = pass
			result += getCombinations(rangesCopy, r.dest, workflows)
		}
		if fail[0] <= fail[1] {
			maps.Copy(ranges, ranges)
			ranges[r.cat] = fail
		}
	}
	result += getCombinations(ranges, workflows[workflowName].final, workflows)

	return result
}
