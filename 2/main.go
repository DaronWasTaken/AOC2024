package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	INPUT_FILE_PATH = "test_input.txt"
)

func main() {

	inputFile := flag.String("f", INPUT_FILE_PATH, "Name of the input file")
	flag.Parse()

	INPUT_FILE_PATH = *inputFile

	fmt.Printf("Using file: %s\n", INPUT_FILE_PATH)
	file, err := os.Open(INPUT_FILE_PATH)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewReader(file)

	result1 := 0
	result2 := 0

	for {
		line, err := buffer.ReadString('\n')
		line, _ = strings.CutSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		dataStrings := strings.Split(line, " ")
		data := make([]int, len(dataStrings))
		for i, str := range dataStrings {
			number, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			data[i] = number
		}

		// fmt.Printf("%v [%v]\n", data, isSafe(data))
		// fmt.Printf("%v [%v]\n", data, isSafeWithDampener(data))

		//PART I
		if isSafe(data) {
			result1++
		}

		//PART II
		if isSafeWithDeletion(data) {
			result2++
		}
	}

	fmt.Printf("PART I: %d\n", result1)
	fmt.Printf("PART II: %d\n", result2)
}

func isSafe(data []int) bool {
	increase, decrease := false, false
	for i := 1; i < len(data); i++ {
		diff := data[i] - data[i-1]
		if diff > 0 {
			increase = true
		} else if diff < 0 {
			decrease = true
		} else {
			return false
		}

		if increase && decrease {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func isSafeWithDeletion(data []int) bool {

	if isSafe(data) {
		return true
	}

	for i := 0; i < len(data); i++ {
		dataCopy := make([]int, len(data))
		copy(dataCopy, data)
		if i == len(dataCopy)-1 {
			dataCopy = dataCopy[:i]
		} else {
			dataCopy = append(dataCopy[:i], dataCopy[i+1:]...)
		}
		if isSafe(dataCopy) {
			return true
		}
	}
	return false
}
