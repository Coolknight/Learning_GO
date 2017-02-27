// An implementation MiniMax for ticktacktoe
package main

import (
	"bytes"
	"fmt"
	"time"
	//"strconv"
	"math/rand"
)

const DEPTH = 3
const SIZE = 3
const DEBUG = false

// Field.s represents the board.
type Field struct {
	s    [][]byte
	last byte
}

// NewField returns an empty field of the specified width and height.
func NewField(a int) *Field {
	s := make([][]byte, a)
	for i := range s {
		s[i] = make([]byte, a)
	}
	for i := range s {
		for j := range s[i] {
			s[i][j] = ' '
		}
	}
	return &Field{s: s, last: '#'}
}

func (f *Field) Copy(c *Field) {
	for i := range c.s {
		copy(f.s[i], c.s[i])
	}
	f.last = c.last
}

func (f *Field) Playable(x, y int) bool {
	return f.s[x][y] == ' '
}

// Set sets the state of the specified position.
func (f *Field) Set(x, y int, b byte) bool {
	if f.Playable(x, y) {
		f.s[x][y] = b
		f.last = b
		return true
	} else {
		return false
	}
}

// Value returns the value of a given board.
func (f *Field) Value() int {
	result := 0
	temp := 0
	values := [SIZE][SIZE]int{}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if f.s[i][j] == 'O' {
				values[i][j] = 2
			} else if f.s[i][j] == ' ' {
				values[i][j] = 1
			} else {
				values[i][j] = -2
			}
		}
	}

	for i := 0; i < SIZE; i++ {
		temp = values[i][0] + values[i][1] + values[i][2]
		switch temp {
		case 6:
			result += 10000
		case -6:
			result -= 10000
		default:
			result += temp
		}
		temp = values[0][i] + values[1][i] + values[2][i]
		switch temp {
		case 6:
			result += 10000
		case -6:
			result -= 10000
		default:
			result += temp
		}
	}
	temp = values[0][0] + values[1][1] + values[2][2]
	switch temp {
	case 6:
		result += 10000
	case -6:
		result -= 10000
	default:
		result += temp
	}
	temp = values[0][2] + values[1][1] + values[2][0]
	switch temp {
	case 6:
		result += 10000
	case -6:
		result -= 10000
	default:
		result += temp
	}
	return result
}

func (f *Field) isFinished() bool {
	full := true
	for i := range f.s {
		for j := range f.s[i] {
			if f.s[i][j] == ' ' {
				full = false
			}
		}
	}
	return f.Value() > 5000 || f.Value() < -5000 || full
}

// Minimax implements the namesake algorithm. Maxs our value. Mins others value.
func (f *Field) Minimax(depth int, max bool) int {
	if (depth == 0) || (f.isFinished()) {
		return f.Value()
	}

	var bestValue int
	try := NewField(len(f.s))
	try.Copy(f)
	best := NewField(len(f.s))
	if max {
		bestValue = -999999
		for i := range try.s {
			for j := range try.s[i] {
				if try.Playable(i, j) {
					try.Copy(f)
					try.Set(i, j, 'O')
					if DEBUG {
						fmt.Println("Maximizando=", bestValue, "ultimo jugador=", try.last)
						fmt.Println(try, "valor", try.Value())
					}
					v := try.Minimax(depth-1, false)
					if v > bestValue {
						bestValue = v
						best.Copy(try)
					}
				}
			}
		}
	} else { //MIN
		bestValue = 999999
		for i := range f.s {
			for j := range f.s[i] {
				if try.Playable(i, j) {
					try.Copy(f)
					try.Set(i, j, 'X')
					if DEBUG {
						fmt.Println("Minimizando=", bestValue, "ultimo jugador=", try.last)
						fmt.Println(try, "valor", try.Value())
					}
					v := try.Minimax(depth-1, true)
					if v < bestValue {
						bestValue = v
						best.Copy(try)
					}
				}
			}
		}
	}
	if depth == DEPTH {
		f.Copy(best)
	}
	return bestValue
}

// String returns the game board as a string.
func (f *Field) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(f.s)*2+1; i++ {
		buf.WriteByte('#')
	}
	buf.WriteByte('\n')
	for y := range f.s {
		buf.WriteByte('|')
		for x := range f.s[y] {
			buf.WriteByte(f.s[y][x])
			buf.WriteByte('|')
		}
		buf.WriteByte('\n')
	}
	for i := 0; i < len(f.s)*2+1; i++ {
		buf.WriteByte('#')
	}
	buf.WriteByte('\n')
	return buf.String()
}

func main() {
	play := false
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var n int
	f := NewField(SIZE)      // Here is where you set the size of the board
	if random.Intn(2) == 0 { //I start
		play = true
	}
	fmt.Print("\033[H\033[2J", f)
	for !f.isFinished() {
		if play {
			fmt.Print("Enter Move: ")
			fmt.Scanf("%d", &n)
			fmt.Println(n)
			switch n {
			case 1:
				f.Set(2, 0, 'X')
			case 2:
				f.Set(2, 1, 'X')
			case 3:
				f.Set(2, 2, 'X')
			case 4:
				f.Set(1, 0, 'X')
			case 5:
				f.Set(1, 1, 'X')
			case 6:
				f.Set(1, 2, 'X')
			case 7:
				f.Set(0, 0, 'X')
			case 8:
				f.Set(0, 1, 'X')
			case 9:
				f.Set(0, 2, 'X')
			default:
				fmt.Print("ERROR")
			}
			fmt.Print("\033[H\033[2J", f)
		}
		f.Minimax(DEPTH, true)
		fmt.Print("\033[H\033[2J", f)
		play = true
	}
}
