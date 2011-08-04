// Copyright 2011, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
/*
 *  Filename:    snail.go
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Wed Aug  3 22:17:18 PDT 2011
 *  Description: Print a snail matrix to standard output.
 *  Usage:       snail -n MATRIXSIZE
 */
import (
    "math"
    "fmt"
)

func main() {
    ParseFlags()
    MakeSnailMatrix(opt.n).Print()
}

type SnailMatrix [][]int

func MakeSnailMatrix(n int) SnailMatrix {
    m := make(SnailMatrix, n)
    for i, _ := range m {
        m[i] = make([]int, n)
    }
    SnailDo(n, func(s *Snail) { m[s.i][s.j] = s.count })
    return m
}

func (m SnailMatrix) Print() {
    elmf := fmt.Sprintf("%%%dd", m.FormatWidth())
    for _, row := range m {
        for _, elm := range row {
            fmt.Printf(elmf, elm)
        }
        fmt.Println()
    }
}

func (m SnailMatrix) FormatWidth() int {
    n := len(m)
    return int(math.Ceil(math.Log10(float64(n*n))-0.5)) + 2
}

//  Walks around a snail matrix using the pattern in side lengths;
//  n, n, n, n-1, n-1, n-2, n-2, ..., 2, 2, 1, 1
type Snail struct {
    dir                            Direction
    n, i, j, count, rem, rep, side int
}

//  Execute a function at each point walking around a snail matrix.
func SnailDo(n int, f func(*Snail)) {
    for s := NewSnail(n); !s.Done(); s.Walk() {
        f(s)
    }
}

func NewSnail(n int) *Snail {
    s := new(Snail)
    *s = Snail{n: n, count: 1, rem: n - 1, rep: -1, side: n - 1}
    return s
}

func (s *Snail) Done() bool { return s.side == 0 }

func (s *Snail) Walk() {
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
        s.i--
    case Down:
        s.i++
    case Left:
        s.j--
    case Right:
        s.j++
    }
    // Count
    s.rem--
    s.count++
}

type Direction int

const (
    Right Direction = iota
    Down
    Left
    Up
)

func (d Direction) Rotate() Direction { return (d + 1) % 4 }
