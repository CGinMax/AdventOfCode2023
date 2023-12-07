package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

const (
	Red        = "red"
	RedCubes   = 12
	Green      = "green"
	GreenCubes = 13
	Blue       = "blue"
	BlueCubes  = 14
)

func partOne(reader *bufio.Reader) {

	sum := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("Read file error." + err.Error())
		}

		colonSlice := strings.Split(line, ":")
		gameID, err := strconv.Atoi(strings.Split(colonSlice[0], " ")[1])
		if err != nil {
			fmt.Println("Get game id error.", err.Error())
			continue
		}
		isPossible := true
		subsetsList := strings.Split(strings.TrimSpace(colonSlice[1]), ";")
		for _, subsets := range subsetsList {
			if !isPossible {
				break
			}
			cubes := strings.Split(strings.TrimSpace(subsets), ", ")
			for _, cube := range cubes {
				if strings.Contains(cube, Red) {
					cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
					if cubeCount > RedCubes {
						isPossible = false
						break
					}
				}

				if strings.Contains(cube, Green) {
					cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
					if cubeCount > GreenCubes {
						isPossible = false
						break
					}
				}

				if strings.Contains(cube, Blue) {
					cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
					if cubeCount > BlueCubes {
						isPossible = false
						break
					}
				}

			}
		}
		if isPossible {
			sum += gameID
		}
	}
	fmt.Println("part one sum:", sum)
}

func partTwo(reader *bufio.Reader) {
	sum := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("Read file error." + err.Error())
		}

		thisRed := 0
		thisGreen := 0
		thisBlue := 0

		colonSlice := strings.Split(line, ":")

		gameID, _ := strconv.Atoi(strings.Split(colonSlice[0], " ")[1])

		subsetsList := strings.Split(strings.TrimSpace(colonSlice[1]), "; ")
		for _, subsets := range subsetsList {

			cubes := strings.Split(strings.TrimSpace(subsets), ", ")
			for _, cube := range cubes {
				if strings.Contains(cube, Red) {
					cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
					thisRed = max(thisRed, cubeCount)
				}

				if strings.Contains(cube, Green) {
					cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
					thisGreen = max(thisGreen, cubeCount)
				}

				if strings.Contains(cube, Blue) {
					cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
					thisBlue = max(thisBlue, cubeCount)
				}
			}
		}
		if thisRed == 0 || thisGreen == 0 || thisBlue == 0 {
			fmt.Printf("Game %d is impossible!", gameID)
			continue
		}

		sum += (thisRed * thisBlue * thisGreen)
	}

	fmt.Println("part two sum:", sum)
}

func main() {
	file, err := os.Open("input")
	// file, err := os.Open("inputBasic")
	if err != nil {
		fmt.Println("Open input file error.", err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// partOne(reader)

	partTwo(reader)
}
