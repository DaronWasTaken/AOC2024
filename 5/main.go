package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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

	beforeMap := make(map[int][]int) // key has to be before each value
	var orders [][]int

	for {
		line, err := buffer.ReadString('\n')
		line, _ = strings.CutSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			key, _ := strconv.Atoi(parts[0])
			value, _ := strconv.Atoi(parts[1])
			beforeMap[key] = append(beforeMap[key], value)
		}

		if strings.Contains(line, ",") {
			strSlice := strings.Split(line, ",")
			orders = append(orders, strToIntSlice(strSlice))
		}
	}

	// fmt.Println(beforeMap)
	// fmt.Println(orders)

	var correctOrders [][]int
	var incorrectOrders [][]int

	for _, order := range orders {
		if isOrderCorrect(order, beforeMap) {
			correctOrders = append(correctOrders, order)
		} else {
			incorrectOrders = append(incorrectOrders, order)
		}
	}

	// fmt.Println("=== Correct orders ===")
	// fmt.Println(correctOrders)

	result := 0

	for _, order := range correctOrders {
		result += order[len(order)/2]
	}

	fmt.Printf("PART I: %d\n", result)

	result2 := 0

	for _, order := range incorrectOrders {
		rearrangeOrder(order, beforeMap)
		result2 += order[len(order)/2]
	}

	fmt.Printf("PART II: %d\n", result2)
}

func rearrangeOrder(order []int, beforeMap map[int][]int) {
	for i := len(order) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			// fmt.Printf("Is [%d] in the beforeMap of [%d]?\n", order[j], order[i])
			if slices.Contains(beforeMap[order[i]], order[j]) {
				// fmt.Printf("Before: %v\n", order)
				temp := order[i]
				order[i] = order[j]
				order[j] = temp
				// fmt.Printf("After: %v\n", order)
				if j > 1 {
					j--
				}
			}
		}
	}
}

func isOrderCorrect(order []int, beforeMap map[int][]int) bool {
	// fmt.Printf("Starting checking order: %v\n", order)
	for i := len(order) - 1; i > 0; i-- {
		// fmt.Printf("Checking if [%d] is not breaking rules - map: %v\n", order[i], beforeMap[order[i]])
		for j := 0; j < i; j++ {
			// fmt.Printf("Is [%d] in the beforeMap of [%d]?\n", order[j], order[i])
			if slices.Contains(beforeMap[order[i]], order[j]) {
				return false
			}
		}
	}
	return true
}

func strToIntSlice(slice []string) []int {
	intSlice := make([]int, len(slice))
	for i, x := range slice {
		num, _ := strconv.Atoi(x)
		intSlice[i] = num
	}
	return intSlice
}
