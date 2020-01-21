package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func howMuchPaper(filename string) (int, int) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totalArea, totalRibbon := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		spl := strings.Split(line, "x")
		sizes := make([]int, len(spl))

		for i := 0; i < len(spl); i++ {
			sizes[i], _ = strconv.Atoi(spl[i])
		}

		face1 := sizes[0] * sizes[1]
		face2 := sizes[1] * sizes[2]
		face3 := sizes[0] * sizes[2]
		sort.Ints(sizes[:])
		minArea := sizes[0] * sizes[1]

		ribbonToWrap := 2*sizes[0] + 2*sizes[1]
		ribbonForBow := sizes[0] * sizes[1] * sizes[2]
		totalRibbon += ribbonToWrap + ribbonForBow

		packageArea := 2*(face1+face2+face3) + minArea
		totalArea += packageArea
	}
	return totalArea, totalRibbon
}

func main() {
	fmt.Println(howMuchPaper("input"))
}
