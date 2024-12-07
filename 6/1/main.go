package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	INPUT_FILE_PATH = "test_input.txt"
	PLAYER_CHAR     = '^'
	WALL_CHAR       = '#'
	MARK_CHAR       = 'X'

	PLAYER_UP    = '^'
	PLAYER_RIGHT = '>'
	PLAYER_DOWN  = 'v'
	PLAYER_LEFT  = '<'

	GAME_TIME = 50 * time.Millisecond
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Player struct {
	themap    [][]byte
	x, y      int
	direction Direction
}

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
	var themap [][]byte
	index := 0

	player := &Player {direction: UP}

	for {
		line, err := buffer.ReadBytes('\n')
		line, _ = bytes.CutSuffix(line, []byte{'\n'})
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if bytes.ContainsRune(line, PLAYER_CHAR) {
			x := bytes.IndexRune(line, PLAYER_CHAR)
			player.x = x
			player.y = index
		}

		themap = append(themap, line)
		index++
	}

	player.themap = themap

	for player_move(player) {
		// printMap(player.themap)
	}
	// printMap(player.themap)
	fmt.Println("END OF GAME")
	fmt.Printf("PART I: %d\n", sumOfMarks(themap))
}

func sumOfMarks(themap [][]byte) int {
	result := 0
	for _, line := range themap {
		for _, char := range line {
			if char == byte(MARK_CHAR) {
				result++
			}
		}
	}
	return result
}

func player_move(self *Player) bool {
	res := _tryMove(self)
	switch res {
	case -1:
		self.themap[self.y][self.x] = byte(MARK_CHAR)
		return false
	case 0:
		if self.direction == 3 {
			self.direction = 0
		} else {
			self.direction++
		}
		player_changeChar(self)
		return true
	case 1:
		return true
	}
	panic("incorrect tryMove result")
}

func player_changeChar(self *Player) {
	switch self.direction {
	case UP:
		PLAYER_CHAR = PLAYER_UP
	case RIGHT:
		PLAYER_CHAR = PLAYER_RIGHT
	case DOWN:
		PLAYER_CHAR = PLAYER_DOWN
	case LEFT:
		PLAYER_CHAR = PLAYER_LEFT
	}
}

func _tryMove(p *Player) int {
	switch p.direction {
	case UP:
		if isOutofBound(p.y-1, p.x, p.themap) {
			return -1
		}
		if isBlocked(p.y-1, p.x, p.themap) {
			return 0
		}
		p.themap[p.y][p.x] = byte(MARK_CHAR)
		p.y--
		p.themap[p.y][p.x] = byte(PLAYER_CHAR)
		return 1
	case RIGHT:
		if isOutofBound(p.y, p.x+1, p.themap) {
			return -1
		}
		if isBlocked(p.y, p.x+1, p.themap) {
			return 0
		}
		p.themap[p.y][p.x] = byte(MARK_CHAR)
		p.x++
		p.themap[p.y][p.x] = byte(PLAYER_CHAR)
		return 1
	case DOWN:
		if isOutofBound(p.y+1, p.x, p.themap) {
			return -1
		}
		if isBlocked(p.y+1, p.x, p.themap) {
			return 0
		}
		p.themap[p.y][p.x] = byte(MARK_CHAR)
		p.y++
		p.themap[p.y][p.x] = byte(PLAYER_CHAR)
		return 1
	case LEFT:
		if isOutofBound(p.y, p.x-1, p.themap) {
			return -1
		}
		if isBlocked(p.y, p.x-1, p.themap) {
			return 0
		}
		p.themap[p.y][p.x] = byte(MARK_CHAR)
		p.x--
		p.themap[p.y][p.x] = byte(PLAYER_CHAR)
		return 1
	}
	panic("Incorrect direction")
}

func isBlocked(y, x int, themap [][]byte) bool {
	if themap[y][x] == byte(WALL_CHAR) {
		return true
	}
	return false
}

func isOutofBound(y, x int, themap [][]byte) bool {
	if y > (len(themap)-1) || y < 0 || x > (len(themap[y])-1) || x < 0 {
		return true
	}
	return false
}

func printMap(themap [][]byte) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	for _, line := range themap {
		for _, char := range line {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}

	time.Sleep(GAME_TIME)
}
