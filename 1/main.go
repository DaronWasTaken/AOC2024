package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	INPUT_LINES = 1000
	INPUT_FILE_PATH = "input.txt"
)

func main() {

	file, err := os.Open(INPUT_FILE_PATH)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewReader(file)

	var firstArr [INPUT_LINES]int
	var secondArr [INPUT_LINES]int

	index := 0

	for {
		line, err := buffer.ReadString('\n')
		line, _ = strings.CutSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		elems := strings.Split(line, "   ")
		first, err := strconv.Atoi(elems[0])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(elems[1])
		if err != nil {
			log.Fatal(err)
		}

		firstArr[index] = first
		secondArr[index] = second

		index++
	}

	partOne(firstArr[:], secondArr[:])
	partTwo(firstArr[:], secondArr[:])
}

func partOne(arr1 []int, arr2 []int) {

	var firstArr [INPUT_LINES]int
	var secondArr [INPUT_LINES]int

	copy(firstArr[:], arr1)
	copy(secondArr[:], arr2)

	sort.Ints(firstArr[:])
	sort.Ints(secondArr[:])

	sum := 0

	for i := 0; i < len(firstArr); i++ {
		distance := firstArr[i] - secondArr[i]
		if distance < 0 {
			distance *= -1
		}
		sum += distance
	}

	fmt.Printf("Part I: %d\n", sum)
}

func partTwo(firstArr []int, secondArr []int) {
	countMap := make(map[int]int)
	sum := 0

	for i := 0; i < len(firstArr); i++ {
		countMap[secondArr[i]]++
	}

	for i := 0; i < len(firstArr); i++ {
		sum += firstArr[i] * countMap[firstArr[i]]
	}

	fmt.Printf("Part II: %d\n", sum)
}