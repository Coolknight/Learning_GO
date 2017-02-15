// An implementation of Chess horse jump
package main

import (
  "bytes"
  "fmt"
  //"os"
  //"time"
  "strconv"
)

// Field.s represents a two-dimensional field of cells.
// Field.iter represents the number of horse jumps in a board
// Field.sols represents the number of solutions found.
type Field struct {
  s    [][]int
  iter int
  sols int
}

// NewField returns an empty field of the specified width and height.
func NewField(a int) *Field {
  s := make([][]int, a)
  for i := range s {
    s[i] = make([]int, a)
  }
  for i := range s {
    for j := range s[i] {
      s[i][j] = 0
    }
  }
  return &Field{s: s, iter: 0, sols: 0}
}

// Set sets the state of the specified position.
func (f *Field) Set(x, y int, b bool) {
  if b {
    f.iter++
    f.s[y][x] = f.iter
  } else {
    f.iter--
    f.s[y][x] = 0
  }
}

// Occupied reports whether a position is occupied or not.
// If the x or y coordinates are outside the field boundaries we act like if they were occupied.
func (f *Field) Occupied(x, y int) bool {
  return x < 0 || x > 7 || y < 0 || y > 7 || f.s[y][x] != 0
}

// Next tries to occupy the given position. If it is not then we iterate. When we are done we print the board.
func (f *Field) Next(x, y int) {
  if ! f.Occupied(x, y) {
    f.Set(x,y,true)
    if ! f.Success() {
      f.Next(x-1, y-2)
      f.Next(x-1, y+2)
      f.Next(x+1, y-2)
      f.Next(x+1, y+2)
      f.Next(x-2, y-1)
      f.Next(x-2, y+1)
      f.Next(x+2, y-1)
      f.Next(x+2, y+1)
    } else {
      fmt.Print("\033[H\033[2J")
      fmt.Print("\x0c", f) // Clear screen and print field.
      f.sols++
    }
    f.Set(x,y,false)
  }
}

// Success returns true if all board has been visited.
func (f *Field) Success() bool {
  var ret bool
  ret = true
  for y := range f.s {
    for x := range f.s[y] {
      if ! f.Occupied(x, y) {
        ret = false
      }
    }
  }
return ret
}

// String returns the game board as a string.
func (f *Field) String() string {
  var buf bytes.Buffer
  var s string
  for i:=0; i<len(f.s)*3+1; i++ {
    buf.WriteByte('#')
  }
  buf.WriteByte('\n')
  for y := range f.s {
    buf.WriteByte('|')
    for x := range f.s[y] {
      s = strconv.Itoa(f.s[y][x]) + "|"
      if len(s) == 3 {
        buf.WriteString(strconv.Itoa(f.s[y][x]) + "|")
      } else {
        buf.WriteString("0" + strconv.Itoa(f.s[y][x]) + "|")
      }
    }
    buf.WriteByte('\n')
  }
  for i:=0; i<len(f.s)*3+1; i++ {
    buf.WriteByte('#')
  }
  buf.WriteByte('\n')
  buf.WriteString(strconv.Itoa(f.sols))
  return buf.String()
}

func main() {
  f := NewField(8) // Here is where you set the size of the board
  f.Next(0, 0)
}
