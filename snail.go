// Copyright 2011, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
/*
 *  Filename:    snail.go
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Wed Aug  3 22:17:18 PDT 2011
 *  Description: Print a snail matrix to standard output.
 *  Usage:       snail [-n=N]
 */
import (
    "math"
    "fmt"
)

func main() {
    parseFlags()
    MakeSnailMatrix(opt.n).Print()
}

//  A matrix with entries increasing in a contracting spiral.
type SnailMatrix [][]int

//  Create and initialize an n by n snail matrix.
func MakeSnailMatrix(n int) SnailMatrix {
    m := make(SnailMatrix, n)
    for i, _ := range m {
        m[i] = make([]int, n)
    }
    SnailDo(n, func(s *Snail) { m[s.I][s.J] = s.Count })
    return m
}

//  Print the matrix with aligned columns to standard out.
func (m SnailMatrix) Print() {
    elmf := fmt.Sprintf("%%%dd", m.numWidth())
    for _, row := range m {
        for _, elm := range row {
            fmt.Printf(elmf, elm)
        }
        fmt.Println()
    }
}

func (m SnailMatrix) numWidth() int {
    n := len(m)
    return int(math.Ceil(math.Log10(float64(n*n))-0.5)) + 2
}

//  Walks around a snail matrix expoiting the pattern in side lengths;
//  N, N, N, N-1, N-1, N-2, N-2, ..., 2, 2, 1, 1
type Snail struct {
    dir                            Direction
    N, I, J, Count, rem, rep, side int
}

//  Execute a function at each point walking around a snail matrix.
func SnailDo(n int, f func(*Snail)) {
    for s := newSnail(n); !s.done(); s.walk() {
        f(s)
    }
}

func newSnail(n int) *Snail {
    s := new(Snail)
    *s = Snail{N: n, Count: 1, rem: n - 1, rep: -1, side: n - 1}
    return s
}

func (s *Snail) done() bool { return s.side == 0 }

func (s *Snail) walk() {
    if s.rem == 0 {
        // Turn
        if s.rep++; s.rep == 2 {
            s.side--
            s.rep = 0
        }
        s.rem = s.side
        s.dir = s.dir.Rotate()
    }
    // Move
    switch s.dir {
    case Up:
        s.I--
    case Down:
        s.I++
    case Left:
        s.J--
    case Right:
        s.J++
    }
    // Count
    s.rem--
    s.Count++
}

type Direction int

const (
    Right Direction = iota
    Down
    Left
    Up
)

//  Return the direction clockwise of d.
func (d Direction) Rotate() Direction { return (d + 1) % 4 }
