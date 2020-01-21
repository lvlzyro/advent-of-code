package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func howManyLit(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lights := [1000][1000]bool{}
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		p1 := regexp.MustCompile("^(toggle|turn on|turn off)")
		operation := p1.FindString(line)
		p2 := regexp.MustCompile("([0-9]+)")
		cs := p2.FindAllString(line, 4)
		c := make([]int, len(cs))
		for i := 0; i < 4; i++ {
			c[i], _ = strconv.Atoi(cs[i])
		}
		for x := c[0]; x <= c[2]; x++ {
			for y := c[1]; y <= c[3]; y++ {
				switch operation {
				case "toggle":
					lights[x][y] = !lights[x][y]
				case "turn on":
					lights[x][y] = true
				case "turn off":
					lights[x][y] = false
				}
			}
		}
	}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if lights[x][y] {
				result++
			}
		}
	}
	return result
}

func totalBrightness(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lights := [1000][1000]int{}

	fromRegex := regexp.MustCompile("\\d+,\\d+")
	toRegex := regexp.MustCompile("\\d+,\\d+$")

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		fromStr := fromRegex.FindString(line)
		toStr := toRegex.FindString(line)
		from := strings.Split(fromStr, ",")
		to := strings.Split(toStr, ",")

		fromX, _ := strconv.Atoi(from[0])
		fromY, _ := strconv.Atoi(from[1])
		toX, _ := strconv.Atoi(to[0])
		toY, _ := strconv.Atoi(to[1])

		for x := fromX; x <= toX; x++ {
			for y := fromY; y <= toY; y++ {
				if strings.Contains(line, "off") {
					lights[x][y]--
					if lights[x][y] < 0 {
						lights[x][y] = 0
					}
				} else if strings.Contains(line, "on") {
					lights[x][y]++
				} else {
					lights[x][y] += 2
				}
			}
		}
	}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			total += lights[x][y]
		}
	}

	return total
}

func main() {
	fmt.Println(howManyLit("input"), totalBrightness("input"))
}
