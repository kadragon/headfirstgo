package count

import (
	"fmt"
	"log"
	"sort"

	"github.com/kadragon/headfirstgo/datafile"
)

func Count() {
	lines, err := datafile.GetStrings("./count/votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

func Count2() {
	lines, err := datafile.GetStrings("./count/votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for key, value := range counts {
		fmt.Printf("Votes for %s: %d", key, value)
	}

	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s: %d", name, counts[name])
	}
}
