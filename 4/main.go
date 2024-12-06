package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	INPUT_FILE_PATH = "test_input2.txt"

	L  = 0
	R  = 0
	U  = 0
	D  = 0
	RD = 0
	RU = 0
	LU = 0
	LD = 0

	X_ULDR = 0
	X_URDL = 0
	X_X    = 0
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
	index := 0

	box := make([][]rune, 0, 140)

	for {
		line, err := buffer.ReadBytes('\n')
		line, _ = bytes.CutSuffix(line, []byte("\n"))
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		box = append(box, bytes.Runes(line))
		index++
	}

	for i := 0; i < len(box); i++ {
		fmt.Printf("%c\n", box[i])
	}

	result1 := 0
	result2 := 0

	for i, slice := range box {
		for j, r := range slice {
			if r == 'X' {
				result1 += checkAreaForMAS(box, i, j)
			}
			if r == 'A' {
				result2 += x_check(box, i, j)
			}
		}
	}

	fmt.Printf("PART I: %d\n", result1)
	fmt.Printf("L: %d\n", L)
	fmt.Printf("R: %d\n", R)
	fmt.Printf("U: %d\n", U)
	fmt.Printf("D: %d\n", D)
	fmt.Printf("LU: %d\n", LU)
	fmt.Printf("LD: %d\n", LD)
	fmt.Printf("RU: %d\n", RU)
	fmt.Printf("RD: %d\n", RD)

	fmt.Printf("PART II: %d\n", result2)
	fmt.Printf("X_X: %d\n", X_X)
	fmt.Printf("X_ULDR: %d\n", X_ULDR)
	fmt.Printf("X_URDL: %d\n", X_URDL)

}

func x_check(box [][]rune, i, j int) int {
	if i+1 > len(box)-1 {
		return 0
	}
	if j+1 > len(box[0])-1 {
		return 0
	}
	if i-1 < 0 {
		return 0
	}
	if j-1 < 0 {
		return 0
	}

	diagonals := 0
	ul_dr := []rune{box[i-1][j-1], box[i+1][j+1]}
	if string(ul_dr) == "MS" || string(ul_dr) == "SM" {
		X_ULDR++
		diagonals++
	}
	ur_dl := []rune{box[i+1][j-1], box[i-1][j+1]}
	if string(ur_dl) == "MS" || string(ur_dl) == "SM" {
		X_URDL++
		diagonals++
	}

	if diagonals == 2 {
		X_X++
		return 1
	}

	return 0
}

func checkAreaForMAS(box [][]rune, i, j int) int {
	result := checkL(box[i], j) + checkR(box[i], j)
	result += checkD(box, i, j) + checkU(box, i, j)
	result += checkLU(box, i, j) + checkRU(box, i, j)
	result += checkLD(box, i, j) + checkRD(box, i, j)
	return result
}

func checkRD(box [][]rune, i, j int) int {
	for x := 1; x < 4; x++ {
		if i+x > len(box)-1 {
			return 0
		}
		if j+x > len(box[0])-1 {
			return 0
		}
	}

	rd := []rune{box[i+1][j+1], box[i+2][j+2], box[i+3][j+3]}
	if string(rd) == "MAS" {
		RD++
		return 1
	}
	return 0
}

func checkLD(box [][]rune, i, j int) int {
	for x := 1; x < 4; x++ {
		if i+x > len(box)-1 {
			return 0
		}
	}
	for x := 1; x < 4; x++ {
		if j-x < 0 {
			return 0
		}
	}
	ld := []rune{box[i+1][j-1], box[i+2][j-2], box[i+3][j-3]}
	if string(ld) == "MAS" {
		LD++
		return 1
	}
	return 0
}

func checkLU(box [][]rune, i, j int) int {
	for x := 1; x < 4; x++ {
		if i-x < 0 {
			return 0
		}
	}
	for x := 1; x < 4; x++ {
		if j-x < 0 {
			return 0
		}
	}
	lu := []rune{box[i-1][j-1], box[i-2][j-2], box[i-3][j-3]}
	if string(lu) == "MAS" {
		LU++
		return 1
	}
	return 0
}

func checkRU(box [][]rune, i, j int) int {
	for x := 1; x < 4; x++ {
		if i-x < 0 {
			return 0
		}
		if j+x > len(box[0])-1 {
			return 0
		}
	}

	ru := []rune{box[i-1][j+1], box[i-2][j+2], box[i-3][j+3]}
	if string(ru) == "MAS" {
		RU++
		return 1
	}
	return 0
}

func checkR(row []rune, index int) int {
	for i := 1; i < 4; i++ {
		if index+i > len(row)-1 {
			return 0
		}
	}

	right := []rune{row[index+1], row[index+2], row[index+3]}
	if string(right) == "MAS" {
		R++
		return 1
	}

	return 0
}

func checkL(row []rune, index int) int {
	for i := 1; i < 4; i++ {
		if index-i < 0 {
			return 0
		}
	}

	left := []rune{row[index-1], row[index-2], row[index-3]}
	if string(left) == "MAS" {
		L++
		return 1
	}

	return 0
}

func checkD(box [][]rune, i, j int) int {
	for x := 1; x < 4; x++ {
		if i+x > len(box)-1 {
			return 0
		}
	}

	down := []rune{box[i+1][j], box[i+2][j], box[i+3][j]}
	if string(down) == "MAS" {
		D++
		return 1
	}
	return 0
}

func checkU(box [][]rune, i, j int) int {
	for x := 1; x < 4; x++ {
		if i-x < 0 {
			return 0
		}
	}
	up := []rune{box[i-1][j], box[i-2][j], box[i-3][j]}
	if string(up) == "MAS" {
		U++
		return 1
	}
	return 0
}
