package days

import (
	"bytes"
	"strings"
	"testing"
)

// func getField(v *LuggageRule, field string) interface{} {
// 	r := reflect.ValueOf(v)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	return f.Int()
// }

func TestParseLuggageRule(t *testing.T) {
	input := []byte("dotted tan bags contain 2 faded fuchsia bags, 2 pale turquoise bags, 4 mirrored gray bags.")
	expect := LuggageRule{
		Container: "dotted tan",
		Contained: []struct {
			Num   int
			Color string
		}{
			{Num: 2, Color: "faded fuchsia"},
			{Num: 2, Color: "pale turquoise"},
			{Num: 4, Color: "mirrored gray"},
		},
	}
	got := parseLuggageRule(input)

	if expect.Container != got.Container {
		t.Errorf("got.Container = %q; want %q", got.Container, expect.Container)
	}

	for i, ex := range expect.Contained {
		if ex.Color != got.Contained[i].Color {
			t.Errorf("got.Contained[%v].Color = %q; want %q", i, got.Contained[i].Color, ex.Color)
		}

		if ex.Num != got.Contained[i].Num {
			t.Errorf("got.Contained[%v].Num = %q; want %q", i, got.Contained[i].Num, ex.Num)
		}
	}
}

func TestLuggageAppearences(t *testing.T) {
	input := bytes.Join([][]byte{
		[]byte("light red bags contain 1 bright white bag, 2 muted yellow bags."),
		[]byte("dark orange bags contain 3 bright white bags, 4 muted yellow bags."),
		[]byte("bright white bags contain 1 shiny gold bag."),
		[]byte("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags."),
		[]byte("shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags."),
		[]byte("dark olive bags contain 3 faded blue bags, 4 dotted black bags."),
		[]byte("vibrant plum bags contain 5 faded blue bags, 6 dotted black bags."),
		[]byte("faded blue bags contain no other bags."),
		[]byte("dotted black bags contain no other bags."),
	}, NL)

	t.Run("CountCanContain", func(t *testing.T) {
		rules := ParseLuggageRules(input)
		expects := []int{4, 2, 5}
		colors := []string{"shiny gold", "bright white", "dark olive"}

		for i, col := range colors {
			got := rules.CountCanContain(col)
			if got != expects[i] {
				t.Errorf("rules.MaxAppearences(%q) = %v; want %v", col, got, expects[i])
			}
		}

	})

}

func TestBagsContained(t *testing.T) {
	input := strings.Join([]string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}, "\n")
	rules := ParseLuggageRules([]byte(input))
	bag := "shiny gold"
	expect := 126
	got := rules.BagsContained(bag)
	if got != expect {
		t.Errorf("rules.BagsContained(%q) = %v; want %v", bag, got, expect)
	}
}
