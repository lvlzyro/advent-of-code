package main

import "io/ioutil"

import "fmt"

func onlySanta(filename string) int {
	input, _ := ioutil.ReadFile(filename)
	x, y := 0, 0
	m := make(map[[2]int]bool, len(input))

	m[[2]int{0, 0}] = true

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}
		m[[2]int{x, y}] = true
	}
	return len(m)
}

func roboSanta(filename string) int {
	input, _ := ioutil.ReadFile(filename)
	x, y := make(map[bool]int, 2), make(map[bool]int)
	m := make(map[[2]int]bool, len(input))

	m[[2]int{0, 0}] = true

	for i := 0; i < len(input); i++ {
		isIt := i%2 == 0

		switch input[i] {
		case '^':
			y[isIt]--
		case '>':
			x[isIt]++
		case 'v':
			y[isIt]++
		case '<':
			x[isIt]--
		}
		m[[2]int{x[isIt], y[isIt]}] = true
	}
	return len(m)
}

func main() {
	fmt.Println(onlySanta("input"))
	fmt.Println(roboSanta("input"))
}
