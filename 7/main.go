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

type Equation struct {
	result  int
	numbers []int
}

type Operator int

const (
	Add Operator = iota
	Mult
	Conc
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
	result := 0
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

		parts := strings.Split(line, ": ")
		eqRes, _ := strconv.Atoi(parts[0])
		numbersStr := strings.Split(parts[1], " ")
		numbers := strToIntSlice(numbersStr)

		eq := &Equation{result: eqRes, numbers: numbers}
		result += evalEquation(eq)
		result2 += evalEquation2(eq)
	}

	fmt.Printf("PART I: %d\n", result)
	fmt.Printf("PART II: %d\n", result2)
}

func evalEquation(eq *Equation) int {
	operatorsSlice := generateCombinations(len(eq.numbers) - 1)
	for _, operators := range operatorsSlice {
		res := eq.numbers[0]
		for i := 1; i < len(eq.numbers); i++ {
			switch operators[i-1] {
			case Add:
				res = res + eq.numbers[i]
			case Mult:
				res = res * eq.numbers[i]
			}
		}
		if res == eq.result {
			return res
		}
	}
	return 0
}

func evalEquation2(eq *Equation) int {
	operatorsSlice := generateCombinations2(len(eq.numbers) - 1)
	for _, operators := range operatorsSlice {
		res := eq.numbers[0]
		for i := 1; i < len(eq.numbers); i++ {
			switch operators[i-1] {
			case Add:
				res = res + eq.numbers[i]
			case Mult:
				res = res * eq.numbers[i]
			case Conc:
				res, _ = strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(eq.numbers[i]))
			}

		}
		if res == eq.result {
			return res
		}
	}
	return 0
}

func strToIntSlice(slice []string) []int {
	intSlice := make([]int, len(slice))
	for i, str := range slice {
		num, _ := strconv.Atoi(str)
		intSlice[i] = num
	}
	return intSlice
}

// FYI: I made this with outside help. Learn graph theory - as you can see it's very useful
func generateCombinations(size int) [][]Operator {
	var result [][]Operator
	var combination []Operator

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == size {
			combinationCopy := make([]Operator, size)
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}

		combination = append(combination, 0)
		backtrack(pos + 1)
		combination = combination[:len(combination)-1]

		combination = append(combination, 1)
		backtrack(pos + 1)
		combination = combination[:len(combination)-1]
	}

	backtrack(0)
	return result
}

func generateCombinations2(size int) [][]Operator {
	var result [][]Operator
	var combination []Operator

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == size {
			combinationCopy := make([]Operator, size)
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}

		for i := 0; i <= 2; i++ {
			combination = append(combination, Operator(i))
			backtrack(pos + 1)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(0)
	return result
}
