
package main

import (
    "fmt"
)

func CheckWin(board [][]string) bool{
    gameOver := false 

    //Check for all columns
    if board[0][0] == board[0][1] && board[0][0] == board[0][2]{
        fmt.Println("Game Over")
        gameOver = true
    }else if board[1][0] == board[1][1] && board[1][0] == board[1][2]{
        fmt.Println("Game Over")
        gameOver = true
    }else if board[2][0] == board[2][1] && board[2][0] == board[2][2]{
        fmt.Println("Game Over")
        gameOver = true
    //Check for rows
    }else if board[0][0] == board[1][0] && board[0][0] == board[2][0]{
        fmt.Println("Game Over")
        gameOver = true
    }else if board[0][1] == board[1][1] && board[0][1] == board[2][1]{
        fmt.Println("Game Over")
        gameOver = true
    }else if board[0][2] == board[1][2] && board[0][2] == board[2][2]{
        fmt.Println("Game Over")
        gameOver = true
    //Check for Diagonals
    }else if board[0][0] == board[1][1] && board[0][0] == board[2][2]{
        fmt.Println("Game Over")
        gameOver = true
    }else if board[0][2] == board[1][1] && board[1][1] == board[2][0]{
        fmt.Println("Game Over")
        gameOver = true
    }

    return gameOver
}

