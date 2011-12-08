// Copyright 2011, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
/*  Filename:    snail.go
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Wed Aug  3 22:17:18 PDT 2011
 *  Description: Print a snail matrix to standard output.
 *  Usage:       snail [-n=N]
 */
import (
	"strings"
	"math"
	"fmt"
)

func main() { fmt.Print(MakeSnailMatrix(opt.n)) }

//  A matrix with increasing entries in a contracting spiral.
type SnailMatrix [][]int

//  Create and initialize an n by n snail matrix.
func MakeSnailMatrix(n int) SnailMatrix {
	m := make(SnailMatrix, n)
	for i, _ := range m {
		m[i] = make([]int, n)
	}
	m.fill(n)
	return m
}

func (m SnailMatrix) fill(n int) { SnailDo(n, func(s *Snail) { m[s.I][s.J] = s.count }) }

//  Print the matrix with aligned columns to standard out.
func (m SnailMatrix) String() string {
	formats, rows, cols := m.rowFormats(), make([]string, len(m)), make([]string, len(m))
	for i, row := range m {
		for j, elm := range row {
			cols[j] = fmt.Sprintf(formats[j], elm)
		}
		rows[i] = fmt.Sprintln(strings.Join(cols, ""))
	}
	return strings.Join(rows, "")
}

//  Return the length of one side of m.
func (m SnailMatrix) Side() int { return len(m) }

//  Return a slice containing m.Side() integer format strings.
func (m SnailMatrix) rowFormats() []string {
	base, numwidth, gap := "%%%dd", m.numWidth(), m.leftGap()
	formats := []string{fmt.Sprintf(base, numwidth-gap), fmt.Sprintf(base, numwidth)}
	for i := 0; i < m.Side()-2; i++ {
		formats = append(formats, formats[len(formats)-1])
	}
	return formats
}

//  The base width with which to format numbers.
func (m SnailMatrix) numWidth() int { n := m.Side(); return width(n * n) }

// The size of the gap on the left column of width m.numWidth().
func (m SnailMatrix) leftGap() int {
	largestLeft := 1
	if n := m.Side(); n > 1 {
		largestLeft = 4*n - 3 - 1
	}
	return m.numWidth() - width(largestLeft) + 1
}

//  Return the number of decimal digits needed for x plus 1 for padding.
func width(x int) int { return int(math.Log10(float64(x))) + 2 }

//  Walks a snail pattern expoiting the side-length pattern; N, N, N, N-1, N-1, ..., 2, 2, 1, 1
type Snail struct {
	dir                   Direction // Direction the snail is facing.
	N, I, J               int       // Snail size, row, column.
	count, side, rep, rem int       // step counter, side-step counter, side-step repetition counter, remaining side-steps
}

//  Execute a function at each point walking around a snail matrix.
func SnailDo(n int, f func(*Snail)) {
	switch n {
	case 1:
		f(newSnail(1))
	default:
		for s := newSnail(n); s.side > 0; s.walk() {
			f(s)
		}
	}
}

func newSnail(n int) *Snail { return &Snail{N: n, rem: n - 1, rep: -1, side: n - 1} }
func (s *Snail) walk()      { s.turn(); moves[s.dir](s); s.rem--; s.count++ }

func (s *Snail) turn() {
	if s.rem == 0 {
		if s.rep++; s.rep == 2 {
			s.side--
			s.rep = 0
		}
		s.rem = s.side
		s.dir = s.dir.Rotate()
	}
}

var moves = []func(s *Snail){
	Up: func(s *Snail) { s.I-- }, Down: func(s *Snail) { s.I++ },
	Left: func(s *Snail) { s.J-- }, Right: func(s *Snail) { s.J++ },
}

type Direction int

const (
	Right Direction = iota // Clockwise of Up.
	Down                   // Clockwise of Right.
	Left                   // Clockwise of Down.
	Up                     // Clockwise of Left.
)

//  Return the direction clockwise of d.
func (d Direction) Rotate() Direction { return (d + 1) % 4 }
