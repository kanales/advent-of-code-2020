package day3

import (
	"bytes"
)

type SlidingArea struct {
	Rows int
	Cols int
	data []byte
}

var nl = []byte{'\n'}

func ParseSlidingArea(input []byte) SlidingArea {
	cols := bytes.Index(input, nl)
	data := bytes.ReplaceAll(input, nl, []byte{})
	rows := len(data) / cols
	return SlidingArea{data: data, Cols: cols, Rows: rows}
}

func (this *SlidingArea) TreeAt(x, y int) bool {
	idx := this.Cols*y + (x % this.Cols)
	return this.data[idx] == '#'
}

func (this *SlidingArea) CastRay(dx, dy int) int {
	count := 0
	for x, y := 0, 0; y < this.Rows; x, y = x+dx, y+dy {
		if this.TreeAt(x, y) {
			count++
		}
	}
	return count
}
