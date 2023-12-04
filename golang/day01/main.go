package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var documents []string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please inpt caleration document:")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err.Error())
			break
		}

		documents = append(documents, input)
	}

	sum := 0
	for _, line := range documents {
		first := 0
		second := 0
		for i := 0; i < len(line); i++ {
			if '0' <= line[i] && line[i] <= '9' {
				first, _ = strconv.Atoi(string(line[i]))
				break

			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				second, _ = strconv.Atoi(string(line[i]))
				break
			}
		}
		sum += first*10 + second
	}
	fmt.Println("sum is ", sum)
}
