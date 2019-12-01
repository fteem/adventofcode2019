package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func trimNewline(s string) string {
	return strings.Trim(s, "\n")
}

func parseInput(fileName string, out chan<- string) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			close(out)
			break
		}
		check(err)

		out <- trimNewline(line)
	}
}

func main() {
	lines := make(chan string)
	go parseInput("input.txt", lines)

	var totalFuel int
	for line := range lines {
		mass := stringToInt(line)
		fuel := mass/3 - 2
		totalFuel += fuel
	}

	fmt.Println(totalFuel)
}
