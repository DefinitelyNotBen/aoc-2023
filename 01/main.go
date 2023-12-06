package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("aoc.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)
	var fileLines []string

	for fs.Scan() {
		fileLines = append(fileLines, findDigits(fs.Text()))
	}

	f.Close()

	sum := 0
	for _, num := range fileLines {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	fmt.Println(sum)

}

func findDigits(s string) string {
	var s1 string
	var s2 string
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			s1 = string(s[i])
			break
		}
	}

	for j := (len(s) - 1); j >= 0; j-- {
		if _, err := strconv.Atoi(string(s[j])); err == nil {
			s2 = string(s[j])
			break
		}
	}

	return s1 + s2
}
