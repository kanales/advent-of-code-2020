package day16

import (
	"bytes"
	"regexp"
	"strconv"
)

var ticketRe *regexp.Regexp

func init() {
	ticketRe = regexp.MustCompile(`(?m)^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$|^(\d+(?:,\d+)*)$`)
}

type Range struct {
	Low  int
	High int
}

type Ticket []int
type Field struct {
	Name   string
	Ranges [2]Range
}

func NewField(name string, ranges [2]Range) Field {
	var field Field
	field.Name = name
	field.Ranges = ranges
	return field
}

func (field *Field) InRange(val int) bool {
	for _, r := range field.Ranges {
		if val >= r.Low && val <= r.High {
			return true
		}
	}
	return false
}

type Note struct {
	Fields        []Field
	YourTicket    Ticket
	NearbyTickets []Ticket
}

func ParseTicket(input []byte) Ticket {
	nums := bytes.Split(input, []byte{','})
	ticket := make([]int, len(nums))
	for i, num := range nums {
		ticket[i], _ = strconv.Atoi(string(num))
	}

	return ticket
}

func ParseNote(input []byte) Note {
	var note Note

	input = bytes.TrimRight(input, "\n")
	matches := ticketRe.FindAllSubmatch(input, -1)
	ptr := 0
	// fields
	for ; ptr < len(matches) && matches[ptr][1] != nil; ptr++ {
		match := matches[ptr]
		ranges := make([]int, 4)
		for i, v := range match[2:6] {
			ranges[i], _ = strconv.Atoi(string(v))
		}

		note.Fields = append(note.Fields, NewField(string(match[1]), [2]Range{
			{Low: ranges[0], High: ranges[1]},
			{Low: ranges[2], High: ranges[3]},
		}))
	}

	// your ticket
	note.YourTicket = ParseTicket(matches[ptr][6])

	// nearby tickets
	matches = matches[ptr+1:]
	note.NearbyTickets = make([]Ticket, len(matches))
	for i, m := range matches {
		note.NearbyTickets[i] = ParseTicket(m[6])
	}

	return note
}

func anyFieldRange(value int, fields []Field) bool {
	for _, field := range fields {
		if field.InRange(value) {
			return true
		}
	}
	return false
}

func (note *Note) Validate() int {
	rate := 0
	for _, ticket := range note.NearbyTickets {
		for _, el := range ticket {
			if !anyFieldRange(el, note.Fields) {
				rate += el
				break
			}
		}
	}
	return rate
}

func (note *Note) YieldValid() <-chan Ticket {
	ch := make(chan Ticket)

	go func() {
		defer close(ch)
	Outer:
		for _, ticket := range note.NearbyTickets {
			for _, el := range ticket {
				if !anyFieldRange(el, note.Fields) {
					continue Outer
				}
			}
			ch <- ticket
		}
	}()
	return ch
}

func flatten(in [][]Field) []Field {
	out := make([]Field, len(in))
	i := 0
	for _, fields := range in {
		for _, f := range fields {
			out[i] = f
			i++
		}
	}
	return out
}

func filter(fields []Field, cond func(*Field) bool) []Field {
	filter := make([]Field, 0, len(fields))
	for _, f := range fields {
		if cond(&f) {
			filter = append(filter, f)
		}
	}
	return filter
}

func (note *Note) OrderedFields() []Field {
	possible := make([][]Field, len(note.Fields))
	for i := range note.Fields {
		possible[i] = make([]Field, len(note.Fields))
		for j, f := range note.Fields {
			possible[i][j] = f
		}
	}

	for ticket := range note.YieldValid() {
		for i, e := range ticket {
			possible[i] = filter(possible[i], func(f *Field) bool {
				return f.InRange(e)
			})
		}
	}

	visited := make(map[string]bool)
	for len(visited) < len(note.Fields) {
		for i := range possible {
			switch len(possible[i]) {
			case 1:
				visited[possible[i][0].Name] = true
			case 0:
				panic("empty possibilities!")
			default:
				possible[i] = filter(possible[i], func(f *Field) bool {
					return !visited[f.Name]
				})
			}
		}
	}

	return flatten(possible)
}
