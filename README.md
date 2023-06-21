# TicTacToe

Simple [Golang](https://go.dev) implementation of [TicTacToe](https://en.wikipedia.org/wiki/tic-tac-toe#gameplay).

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
func (t *TicTacToe) GetTurn(player Player, row, col int) Player
```
```golang
// GetWinner returns the winner of the TicTacToe game
func (t *TicTacToe) GetWinner(player Player, row, col int) *Player
```
```golang
// GetBoard returns the TicTacToe board
func (t *TicTacToe) GetBoard(player Player, row, col int) [BoardSize][BoardSize]*Player
```
