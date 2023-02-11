package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/kokizzu/gotro/S"
)

// merge result from multiple runs

//go:embed last.txt
var last string

//go:embed README.md
var readme string

type sorter struct {
	TestName   string
	RunCount   int64
	Duration   float64
	AllocBytes int
	AllocCount int
}

func printHeader(prefix string) {
	fmt.Printf("\n%48s %8s %5s %6s %3s\n", prefix, "total", "ns/op", "B/op", "allocs/op")
}
func (s sorter) PrintUnique(m map[string]bool) {
	if _, ok := m[s.TestName]; !ok {
		fmt.Printf("%48s %8d %5.0f %6d %3d\n", s.TestName, s.RunCount, s.Duration, s.AllocBytes, s.AllocCount)
		m[s.TestName] = true
	}
}

func main() {
	map2struct := []sorter{}
	struct2map := []sorter{}
	struct2struct := []sorter{}
	lines := S.Split(last+"\n"+readme, "\n")
	for _, line := range lines {
		line = S.Trim(line)
		if len(line) == 0 {
			continue
		}
		if line == `## Last Result? Go 1.18` {
			break
		}
		cells := strings.Fields(line)
		var row sorter
		if len(cells) == 8 {
			if cells[3] != `ns/op` {
				log.Println(cells, `not in ns/op`)
				continue
			}
			if cells[5] != `B/op` {
				log.Println(cells, `not in B/op`)
				continue
			}
			if cells[7] != `allocs/op` {
				log.Println(cells, `not in allocs/op`)
				continue
			}
			row = sorter{
				TestName:   S.RightOf(cells[0], `_`),
				RunCount:   S.ToI(cells[1]),
				Duration:   S.ToF(cells[2]),
				AllocBytes: S.ToInt(cells[4]),
				AllocCount: S.ToInt(cells[6]),
			}
		} else if len(cells) == 5 {
			row = sorter{
				TestName:   cells[0],
				RunCount:   S.ToI(cells[1]),
				Duration:   S.ToF(cells[2]),
				AllocBytes: S.ToInt(cells[3]),
				AllocCount: S.ToInt(cells[4]),
			}
		} else {
			log.Println(cells, `wrong cell count`)
			continue
		}

		if row.RunCount <= 0 || row.Duration <= 0 {
			continue
		}

		if S.StartsWith(row.TestName, "M2S_") {
			map2struct = append(map2struct, row)
		} else if S.StartsWith(row.TestName, `S2M_`) {
			struct2map = append(struct2map, row)
		} else if S.StartsWith(row.TestName, `S2S_`) {
			struct2struct = append(struct2struct, row)
		}
	}
	sort.Slice(map2struct, func(i, j int) bool {
		return map2struct[i].Duration < map2struct[j].Duration
	})
	sort.Slice(struct2map, func(i, j int) bool {
		return struct2map[i].Duration < struct2map[j].Duration
	})
	sort.Slice(struct2struct, func(i, j int) bool {
		return struct2struct[i].Duration < struct2struct[j].Duration
	})

	printHeader("map to struct")
	m := map[string]bool{}
	for _, row := range map2struct {
		row.PrintUnique(m)
	}
	printHeader("struct to map")
	m = map[string]bool{}
	for _, row := range struct2map {
		row.PrintUnique(m)
	}
	printHeader("struct to struct")
	m = map[string]bool{}
	for _, row := range struct2struct {
		row.PrintUnique(m)
	}
}
