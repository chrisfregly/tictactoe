package tictactoe

import "fmt"

type Player string

const (
	X Player = "X"
	O Player = "O"

	BoardSize = 3
)

// TicTacToe - https://en.wikipedia.org/wiki/tic-tac-toe#gameplay
type TicTacToe struct {
	turn   Player
	winner *Player
	board  [BoardSize][BoardSize]*Player
}

// NewTicTacToe creates a new instance of a TicTacToe game
func NewTicTacToe() TicTacToe {
	return TicTacToe{
		turn:   X,
		winner: nil,
		board:  [BoardSize][BoardSize]*Player{},
	}
}

// Move places the player's symbol in the given row and column and errors if move is invalid
func (t *TicTacToe) Move(player Player, row, column int) error {
	if t.IsGameOver() {
		return fmt.Errorf("tictactoe: game is already over")
	}
	if player != t.turn {
		return fmt.Errorf("tictactoe: not %s's turn", player)
	}
	if row < 0 || row >= BoardSize || column < 0 || column >= BoardSize {
		return fmt.Errorf("tictactoe: row and/or column is out of bounds")
	}
	if t.board[row][column] != nil {
		return fmt.Errorf("tictactoe: location %d,%d is not empty", row, column)
	}
	t.board[row][column] = &player
	t.winner = getWinner(t.board)
	if t.winner == nil {
		t.turn = getNextTurn(t.turn)
	}
	return nil
}

// GetTurn returns the turn of the TicTacToe game
func (t *TicTacToe) GetTurn() Player {
	return t.turn
}

// GetWinner returns the winner of the TicTacToe game
func (t *TicTacToe) GetWinner() *Player {
	return t.winner
}

// GetBoard returns the TicTacToe board
func (t *TicTacToe) GetBoard() [BoardSize][BoardSize]*Player {
	return t.board
}

// IsGameOver returns whether or not the game is over
func (t *TicTacToe) IsGameOver() bool {
	return t.winner != nil || isBoardFull(t.board)
}

// getWinner checks rows, columns, and diagonals for three in a row and returns a pointer to team or nil
func getWinner(board [BoardSize][BoardSize]*Player) *Player {
	for i := 0; i < BoardSize; i++ {
		if board[i][0] != nil && board[i][1] != nil && board[i][2] != nil &&
			*board[i][0] == *board[i][1] && *board[i][1] == *board[i][2] {
			return board[i][0]
		}
		if board[0][i] != nil && board[1][i] != nil && board[2][i] != nil &&
			*board[0][i] == *board[1][i] && *board[1][i] == *board[2][i] {
			return board[0][i]
		}
	}
	if board[0][0] != nil && board[1][1] != nil && board[2][2] != nil &&
		*board[0][0] == *board[1][1] && *board[1][1] == *board[2][2] {
		return board[0][0]
	}
	if board[2][0] != nil && board[1][1] != nil && board[0][2] != nil &&
		*board[2][0] == *board[1][1] && *board[1][1] == *board[0][2] {
		return board[2][0]
	}
	return nil
}

// getNextTurn returns the opposite of currentTurn
func getNextTurn(currentTurn Player) Player {
	if currentTurn == X {
		return O
	}
	return X
}

// isBoardFull returns whether or not the board is full
func isBoardFull(board [BoardSize][BoardSize]*Player) bool {
	for _, row := range board {
		for _, player := range row {
			if player == nil {
				return false
			}
		}
	}
	return true
}
