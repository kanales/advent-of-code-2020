package days

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

type LuggageRule struct {
	Container string
	Contained []struct {
		Num   int
		Color string
	}
}

type LuggageRules map[string][]struct {
	Num   int
	Color string
}

var reLuggageRule *regexp.Regexp

func init() {
	// Matches one line
	reLuggageRule = regexp.MustCompile(`(\w+ \w+) bags?|(\d+)+ (\w+ \w+) bags?`)
}

func parseLuggageRule(input []byte) LuggageRule {
	tokens := reLuggageRule.FindAllSubmatch(input, -1)
	rule := LuggageRule{}
	rule.Container = string(tokens[0][1])
	for _, token := range tokens[1:] {
		num, _ := strconv.Atoi(string(token[2]))
		cont := struct {
			Num   int
			Color string
		}{Num: num, Color: string(token[3])}
		rule.Contained = append(rule.Contained, cont)
	}
	return rule
}

func ParseLuggageRules(input []byte) LuggageRules {
	lines := bytes.Split(bytes.TrimRight(input, "\n"), NL)
	rules := make(LuggageRules, len(lines))
	for _, line := range lines {
		rule := parseLuggageRule(line)
		rules[rule.Container] = rule.Contained
	}
	return rules
}

func (rule *LuggageRule) ColorAppearences(color string) int {
	count := 0
	for _, contained := range rule.Contained {
		if contained.Color == color {
			count += contained.Num
		}
	}
	return count
}

func (rule *LuggageRule) CanContain(color string) bool {
	for _, contained := range rule.Contained {
		if contained.Color == color {
			return true
		}
	}
	return false
}

func (rules *LuggageRules) canContainMemo(container string, contained string, memo map[string]bool) bool {
	if val, ok := memo[container]; ok {
		return val
	}
	for _, v := range (*rules)[container] {
		if v.Color == contained {
			memo[container] = true
			return true
		}
	}

	for _, v := range (*rules)[container] {
		if rules.CanContain(v.Color, contained) {
			memo[container] = true
			return true
		}
	}
	memo[container] = false
	return false
}

func (rules *LuggageRules) CanContain(container string, contained string) bool {
	for _, v := range (*rules)[container] {
		if v.Color == contained {
			return true
		}
	}

	for _, v := range (*rules)[container] {
		if rules.CanContain(v.Color, contained) {
			return true
		}
	}
	return false
}

func (rules *LuggageRules) CountCanContain(color string) int {
	count := 0
	memo := make(map[string]bool)

	fmt.Printf("%v\n", (*rules)["dotted tan"])
	for k := range *rules {
		if rules.canContainMemo(k, color, memo) {
			count += 1
		}
	}
	return count
}

func (rules *LuggageRules) BagsContained(color string) int {
	count := 0
	for _, col := range (*rules)[color] {
		count += col.Num + col.Num*rules.BagsContained(col.Color)
	}
	return count
}
