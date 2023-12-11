package main

import (
	"fmt"
	"os"
	"os/exec"
)

var board [3][3]string
var currentPlayer string

func main() {
	initializeBoard()
	currentPlayer = "X"

	for {
		clearScreen()
		printBoard()

		if isBoardFull() {
			fmt.Println("It's a draw!")
			break
		}

		fmt.Printf("Player %s's turn\n", currentPlayer)
		row, col := getPlayerMove()

		if isValidMove(row, col) {
			board[row][col] = currentPlayer

			if hasCurrentPlayerWon() {
				clearScreen()
				printBoard()
				fmt.Printf("Player %s wins!\n", currentPlayer)
				break
			}

			switchPlayer()
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----")
		}
	}
}

func getPlayerMove() (int, int) {
	var row, col int
	fmt.Print("Enter row (0-2): ")
	fmt.Scan(&row)
	fmt.Print("Enter column (0-2): ")
	fmt.Scan(&col)
	return row, col
}

func isValidMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != " " {
		return false
	}
	return true
}

func hasCurrentPlayerWon() bool {
	// Check rows, columns, and diagonals
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func clearScreen() {
	cmd := exec.Command("clear") // for Linux and MacOS
	cmd.Stdout = os.Stdout
	cmd.Run()
}
