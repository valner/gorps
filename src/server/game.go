package main

import "fmt"

import "math"

import "types"

type EGameState byte
const (
    // Game state
    ETURN EGameState = iota
    EDRAW = iota
    EWIN1 = iota
    EWIN2 = iota
)

var egameStates = [...]string {
    "Player 1 turn",
    "Player 2 turn",
    "Draw",
    "Player 1 win",
    "Player 2 win",
}

func (gameState EGameState) String() string {
    return egameStates[gameState]
}

type gameField struct {
    field [][]Figure
    xlen byte
    ylen byte
}

func (field *gameField) CheckCoord(x, y int) bool{
    return x >= 0 && x < field.xlen && y >= 0 && y < field.ylen
}

func createField(xlen,ylen int) field *gameField {
    field = new(gameField)
    field.xlen = xlen
    field.ylen = ylen
    field.field := make([][]Figure, xlen))
    for i := range f {
        field.field[i] = make([]Figure, ylen)
    }
    return field
}

type DrawBox struct
{
    P1Choose types.EFigureType
    P2Choose types.EFigureType
}

type Figure struct
{
    player byte
    figureType types.EFigureType

}

type Game struct {
    gameField* field
    gameState EGameState
    waitTurn  byte // 1 and 2 - players, 0 - any
}

type Coord struct {
    x,y int
}



func (game *Game) Move(player byte, from, dmove Coord) types.EMoveResult{
    if game.waitTurn != 0 && game.waitTurn != player {
        return types.EWRONG_TURN
    }

    if game.gameState == EP1WIN || game.gameState == EP2WIN {
        return types.EGAME_ENDED
    }

    if !game.field.CheckCoord(from.x, from.y){
        return types.EWRONG_TURN
    }

    if dmove.x < -1 || dmove.x > 1 || dmove.y < -1 || dmove.y > 1 {
        return types.EWRONG_TURN
    }

    if !game.field.CheckCoord(from.x + dmove.x, from.y + dmove.y) {
        return types.EWRONG_TURN

    activeFigure := game.field[from.x][from.y]

    if activeFigure.figureType == ENONE {
        return types.EWRONG_TURN
    }

    if activeFigure.player != player {
        return types.EWRONG_TURN
    }




}

func CreateGame() *game {
    newGame := new(game)
    newGame.gameField = createField(6,8)
    newGame.gameState = EP1TURN
    return newGame
}

func main() {
    g := CreateGame()

    fmt.Printf("game state now is %s", g.gameState)
}




