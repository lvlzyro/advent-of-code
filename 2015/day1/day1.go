package main

import "io/ioutil"

import "fmt"

func whereIsSanta(text []byte) (int, int) {
	floor, basementpos := 0, 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			if basementpos == 0 {
				basementpos = i + 1
			}
		}
	}
	return floor, basementpos
}

func main() {
	input, _ := ioutil.ReadFile("input")

	fmt.Println(whereIsSanta(input))
}
