package main

import (
	"fmt"
)

type Pair struct {
	start int
	end   int
}

func main() {
	in := []Pair{
		{1, 3},
		{2, 3},
		{4, 4},
		{4, 5},
		{6, 6},
		{6, 9},
	}
	out := reduce(in)
	for _, p := range out {
		fmt.Printf("интервал %d - %d\n", p.start, p.end)
	}
}

func reduce(in []Pair) []Pair {
	if len(in) < 1 {
		return in
	}
	rez := make([]Pair, len(in))
	t := in[0]
	j := 0
	for _, el := range in[1:] {
		if el.start <= t.end {
			t.end = el.end
		} else {
			rez[j] = t
			j++
			t = el
		}
	}
	rez[j] = t
	j++

	return rez[:j]
}
