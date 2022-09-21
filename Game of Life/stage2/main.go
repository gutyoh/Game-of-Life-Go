package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Universe struct {
	state [][]int
}

func (u *Universe) nextGeneration() {
	newState := make([][]int, len(u.state))

	for i := range u.state {
		newState[i] = make([]int, len(u.state[i]))
		copy(newState[i], u.state[i])
	}

	size := len(u.state)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			neighbors := 0

			neighbors += u.state[(i-1+size)%size][(j-1+size)%size]
			neighbors += u.state[(i-1+size)%size][j]
			neighbors += u.state[(i-1+size)%size][(j+1)%size]

			neighbors += u.state[i][(j+1)%size]
			neighbors += u.state[(i+1)%size][(j+1)%size]

			neighbors += u.state[(i+1)%size][j]
			neighbors += u.state[(i+1)%size][(j-1+size)%size]
			neighbors += u.state[i][(j-1+size)%size]

			cell := u.state[i][j]
			if cell == 1 {
				if neighbors < 2 || neighbors > 3 {
					newState[i][j] = 0
				}
			} else {
				if neighbors == 3 {
					newState[i][j] = 1
				}
			}
		}
	}

	u.state = newState
}

func (u Universe) String() string {
	rows := []string{}

	for _, row := range u.state {
		r := ""
		for _, cell := range row {
			if cell == 1 {
				r += "O"
			} else {
				r += " "
			}
		}
		rows = append(rows, r)
	}
	return strings.Join(rows, "\n")
}

func NewUniverse(size, seed int) *Universe {
	rand.Seed(int64(seed))

	state := make([][]int, 0, size)

	for i := 0; i < size; i++ {
		row := make([]int, 0, size)
		for j := 0; j < size; j++ {
			row = append(row, rand.Intn(2))
		}
		state = append(state, row)
	}
	return &Universe{state}
}

func main() {
	var n int
	var s int
	var m int

	_, err := fmt.Scanf("%d %d %d", &n, &s, &m)
	if err != nil {
		log.Fatal(err)
	}

	universe := NewUniverse(n, s)
	for i := 0; i < m; i++ {
		universe.nextGeneration()
	}
	fmt.Println(universe)
}
