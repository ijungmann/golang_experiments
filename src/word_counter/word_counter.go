package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, _ := ioutil.ReadFile("./test.txt")
	//Test
	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	wordMap := make(map[string]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")
		for i := 0; i < len(words); i++ {
			wordMap[words[i]] += 1
		}

	}
	for word, count := range wordMap {
		fmt.Printf("%s : %d\n", word, count)
	}
}
