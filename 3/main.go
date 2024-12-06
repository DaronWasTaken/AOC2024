package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	data, err := os.ReadFile(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile(`(mul\(\d+,\d+\))|(do\(\))|(don\'t\(\))`)
	if err != nil {
		log.Fatal(err)
	}

	match := r.FindAllString(string(data), -1)

	doParsingSlice := make([]string, len(match))

	doFlag := true
	for _, str := range match {
		switch str {
		case "do()":
			doFlag = true
		case "don't()":
			doFlag = false
		default:
			if doFlag {
				doParsingSlice = append(doParsingSlice, str)
			}
		}
	}

	fmt.Printf("PART I: %d\n", getExprResult(match))
	fmt.Printf("PART II: %d\n", getExprResult(doParsingSlice))
}


// [mul(x,y), ...]
func getExprResult(data []string) int {
	var sb strings.Builder
	for _, str := range data {
		sb.WriteString(str)
	}

	r, err := regexp.Compile(`\d+,\d+`)
	if err != nil {
		log.Fatal(err)
	}

	match := r.FindAllString(sb.String(), -1)
	result := 0

	for _, str := range match {
		digits := strings.Split(str, ",")
		assert(len(digits) == 2)
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		result += x * y
	}

	return result
}

func assert(expr bool) {
	if !expr {
		panic(expr)
	}
}
