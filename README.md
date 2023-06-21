# TicTacToe

Simple [Golang](https://go.dev) implementation of [TicTacToe](https://en.wikipedia.org/wiki/tic-tac-toe#gameplay).

## Exposed Constants
```golang
// Player may be X or O
const (
    X Player = "X"
    O Player = "O"
)
```
```golang
// BoardSize defines the length and width of the board
const BoardSize = 3
```

## Exposed Methods

```golang
// NewTicTacToe creates a new instance of a TicTacToe game
func NewTicTacToe() TicTacToe
```
```golang
// Move places the player's symbol in the given row and column and errors if move is invalid
func (t *TicTacToe) Move(player Player, row, column int) error
```
```golang
// GetTurn returns the turn of the TicTacToe game
func (t *TicTacToe) GetTurn() Player
```
```golang
// GetWinner returns the winner of the TicTacToe game
func (t *TicTacToe) GetWinner() *Player
```
```golang
// GetBoard returns the TicTacToe board
func (t *TicTacToe) GetBoard() [BoardSize][BoardSize]*Player
```
