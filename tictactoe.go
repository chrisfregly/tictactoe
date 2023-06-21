package tictactoe

import "fmt"

type Team string

const (
	X Team = "X"
	O Team = "O"

	BoardSize = 3
)

// TicTacToe - https://en.wikipedia.org/wiki/tic-tac-toe#gameplay
type TicTacToe struct {
	turn   Team
	winner *Team
	board  [BoardSize][BoardSize]*Team
}

// NewTicTacToe creates a new instance of a TicTacToe game
func NewTicTacToe() TicTacToe {
	return TicTacToe{
		turn:   X,
		winner: nil,
		board:  [BoardSize][BoardSize]*Team{},
	}
}

// MarkLocation places the Team's symbol in the given row and column
func (t *TicTacToe) MarkLocation(team Team, row, column int) error {
	if t.winner != nil {
		return fmt.Errorf("tictactoe: game is already over")
	}
	if team != t.turn {
		return fmt.Errorf("tictactoe: not %s's turn", team)
	}
	if row < 0 || row >= BoardSize || column < 0 || column >= BoardSize {
		return fmt.Errorf("tictactoe: row and/or column is out of bounds")
	}
	if t.board[row][column] != nil {
		return fmt.Errorf("tictactoe: location to mark is not empty")
	}
	t.board[row][column] = &team
	t.winner = getWinner(t.board)
	if t.winner == nil {
		t.turn = getNextTurn(t.turn)
	}
	return nil
}

// GetTurn returns the turn of the TicTacToe game
func (t *TicTacToe) GetTurn(team Team, row, col int) Team {
	return t.turn
}

// GetWinner returns the winner of the TicTacToe game
func (t *TicTacToe) GetWinner(team Team, row, col int) *Team {
	return t.winner
}

// GetBoard returns the TicTacToe board
func (t *TicTacToe) GetBoard(team Team, row, col int) [BoardSize][BoardSize]*Team {
	return t.board
}

// getWinner checks rows, columns, and diagonals for three in a row and returns a pointer to team or nil
func getWinner(board [BoardSize][BoardSize]*Team) *Team {
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
func getNextTurn(currentTurn Team) Team {
	if currentTurn == X {
		return O
	}
	return X
}
