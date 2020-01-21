package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func howManyNice(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		rule1, _ := regexp.MatchString("[aeiou].*[aeiou].*[aeiou]", line)

		rule2 := false
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				rule2 = true
			}
		}
		rule3, _ := regexp.MatchString("ab|cd|pq|xy", line)
		if rule1 && rule2 && !rule3 {
			result++
		}
	}
	return result
}

func newRules(filename string) int {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		rule1 := false
		for i := 0; i < len(line)-1; i++ {
			pair := string([]byte{line[i], line[i+1]})
			index := strings.Index(line[i+2:], pair)
			if index >= 0 {
				rule1 = true
				break
			}
		}
		rule2 := false
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				rule2 = true
				break
			}
		}
		if rule1 && rule2 {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(howManyNice("input"))
	fmt.Println(newRules("input"))
}
