package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var documents []string
	file, err := os.Open("input")
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
		fmt.Printf("first:%d, second:%d\n", first, second)
		sum += first*10 + second
	}
	fmt.Println("sum is", sum)
}
