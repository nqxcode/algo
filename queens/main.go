package main

import "fmt"

const Queen = '✓'
const Noop = '-'

type Board struct {
	board   [][]rune
	results [][]string
}

func (b *Board) M() int {
	return len(b.board)
}

func (b *Board) N() int {
	if len(b.board) == 0 {
		return 0
	}
	return len(b.board[0])
}

func (b *Board) Init(m, n int) {
	b.board = make([][]rune, m)
	for j := 0; j < m; j++ {
		b.board[j] = make([]rune, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.board[i][j] = Noop
		}
	}
}

func (b *Board) IsAllowed(row, col int) bool {
	for i := 0; i < row; i++ {
		if b.board[i][col] == Queen {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b.board[i][j] == Queen {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j < b.N(); i, j = i-1, j+1 {
		if b.board[i][j] == Queen {
			return false
		}
	}

	return true
}

func (b *Board) doStep(row int) {
	if b.M() == row {
		var result []string
		for i := 0; i < b.M(); i++ {
			result = append(result, string(b.board[i]))
		}

		b.results = append(b.results, result)
		return
	}

	for i := 0; i < b.N(); i++ {
		if b.IsAllowed(row, i) {
			b.board[row][i] = Queen
			b.doStep(row + 1)
			b.board[row][i] = Noop
		}
	}
}

func (b *Board) Run() {
	b.doStep(0)
}

func (b *Board) PrintResults() {
	for i, result := range b.results {
		fmt.Println(fmt.Sprintf("Solution: %d", i+1))
		for _, row := range result {
			fmt.Println(row)
		}
		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("Total: %d", len(b.results)))
}

func main() {
	board := new(Board)
	board.Init(8, 8)
	board.Run()
	board.PrintResults()
}
