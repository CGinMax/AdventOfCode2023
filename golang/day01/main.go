package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func onlyDigital(documents []string) {
	sum := 0
	for _, line := range documents {
		first := 0
		second := 0
		for i := 0; i < len(line); i++ {
			if '0' <= line[i] && line[i] <= '9' {
				first = int(line[i] - '0')
				break

			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				second = int(line[i] - '0')
				break
			}
		}
		sum += first*10 + second
	}
	fmt.Println("sum:", sum)
}

func digitalAndLetters(documents []string) {

	letters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0
	for _, line := range documents {
		first := 0
		firstInedx := len(line)
		second := 0
		secondIndex := -1
		for num, numStr := range letters {
			numIndex := strings.Index(line, strconv.Itoa(num+1))
			numStrIndex := strings.Index(line, numStr)
			if numIndex >= 0 && numIndex < firstInedx {
				first = num + 1
				firstInedx = numIndex
			}
			if numStrIndex >= 0 && numStrIndex < firstInedx {
				first = num + 1
				firstInedx = numStrIndex
			}
		}

		for num, numStr := range letters {
			numIndex := strings.LastIndex(line, strconv.Itoa(num+1))
			numStrIndex := strings.LastIndex(line, numStr)
			if numIndex >= 0 && numIndex > secondIndex {
				second = num + 1
				secondIndex = numIndex
			}
			if numStrIndex >= 0 && numStrIndex > secondIndex {
				second = num + 1
				secondIndex = numStrIndex
			}
		}

		sum += first*10 + second
	}
	fmt.Println("sum:", sum)
}

func main() {
	var documents []string
	file, err := os.Open("input")
	// file, err := os.Open("input_basic")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err.Error())
			return

		}

		documents = append(documents, input)
	}

	onlyDigital(documents)
	digitalAndLetters(documents)

}
