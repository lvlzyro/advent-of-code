package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fiveZeroes(filename string) int {
	input, _ := ioutil.ReadFile(filename)
	i := 0
	for {
		data := strings.TrimSpace(string(input)) + strconv.Itoa(i)
		result := md5.Sum([]byte(data))
		str := hex.EncodeToString(result[:])
		if str[0:5] == "00000" {
			return i
		}
		i++
	}
}

func sixZeroes(filename string) int {
	input, _ := ioutil.ReadFile(filename)
	i := 0
	for {
		data := strings.TrimSpace(string(input)) + strconv.Itoa(i)
		result := md5.Sum([]byte(data))
		str := hex.EncodeToString(result[:])
		if str[0:6] == "000000" {
			return i
		}
		i++
	}
}

func main() {
	fmt.Println(fiveZeroes("input"))
	fmt.Println(sixZeroes("input"))
}
