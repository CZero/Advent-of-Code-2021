package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/CZero/gofuncy/lfs"
)

type solution struct {
	path    []string       // The path followed -> Last entry is our location :D
	visited map[string]int // #Times a cave was visited
}

var found [][]string

func main() {
	_ = found
	// maplines, err := lfs.ReadLines("smallexample.txt")
	// maplines, err := lfs.ReadLines("example.txt")
	// maplines, err := lfs.ReadLines("bigexample.txt")
	maplines, err := lfs.ReadLines("input.txt")
	lfs.PanErr(err)
	cavesmap := ReadMap(maplines)
	_ = cavesmap
	start := solution{path: []string{"start"}, visited: map[string]int{"start": 1}}
	// fmt.Printf("%#v\n", start)
	explore(cavesmap, start)
	for i, path := range found {
		fmt.Printf("%2d: %v\n", i+1, path)
	}
}

func explore(cavesmap map[string][]string, Solution solution) {
	for _, exit := range cavesmap[Solution.path[len(Solution.path)-1]] {
		// First copy the solution to a new solution for detached travel per exit ;)
		var onwards solution
		onwards.path = append(onwards.path, Solution.path...)
		onwards.visited = make(map[string]int)
		for key, value := range Solution.visited {
			onwards.visited[key] = value
		}
		// Copy done

		// Check if the exit is valid (unvisited), or even the end, else bail or return the end
		switch {
		case exit == "end": // Einde!
			onwards.path = append(onwards.path, exit)
			found = append(found, onwards.path)
			// fmt.Printf("Einde gevonden\n")
		case !unicode.IsUpper(rune(exit[0])) && onwards.visited[exit] > 0: // Kleine grot reeds bezocht
			// fmt.Printf("Exit: %s\nVisited: %#v\nPath: %#v\n", exit, onwards.visited, onwards.path)
			// fmt.Printf("2e keer naar %s, ik stop..\n", exit)
		case !unicode.IsUpper(rune(exit[0])) && onwards.visited[exit] == 0: // Kleine grot eerste bezoek
			onwards.path = append(onwards.path, exit)
			onwards.visited[exit]++
			// fmt.Printf("Onwards: %#v\n", onwards.path)
			explore(cavesmap, onwards)
		case unicode.IsUpper(rune(exit[0])): // Grote grot, we kunnen altijd verder.
			onwards.path = append(onwards.path, exit)
			onwards.visited[exit]++
			// fmt.Printf("Onwards: %#v\n", onwards.path)
			explore(cavesmap, onwards)
		default: // Uhm, zou niet moeten lukken?
			probleem := fmt.Sprintf("Exit: %s\nVisited: %#v\nPath: %#v\n", exit, onwards.visited, onwards.path)
			panic(probleem)
		}
	}
	return
}

func ReadMap(maplines []string) map[string][]string {
	var (
		cavesmap   = make(map[string][]string)
		checkIfnew = make(map[string]int)
	)
	for _, line := range maplines {
		fields := strings.Split(line, "-")
		combo := fields[0] + fields[1]
		comboreverse := fields[1] + fields[0]
		if checkIfnew[combo] < 1 {
			cavesmap[fields[0]] = append(cavesmap[fields[0]], fields[1])
			checkIfnew[combo]++
		}
		if checkIfnew[comboreverse] < 1 {
			cavesmap[fields[1]] = append(cavesmap[fields[1]], fields[0])
			checkIfnew[comboreverse]++
		}
	}
	// fmt.Printf("%#v\n", cavesmap)
	return cavesmap
}
