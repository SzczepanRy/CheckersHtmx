package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

type BoardI interface {
	ReadBackgroundBoard()
	ReadPlayerBoard()
}

type Board struct {
	BackgroundBoard [8][8]int
	PlayerBoard     [8][8]int
}

func InitBoard() Board {
	var newBoard Board
	newBoard.BackgroundBoard = [8][8]int{
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
	}
	newBoard.PlayerBoard = [8][8]int{
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{2, 0, 2, 0, 2, 0, 2, 0},
		{0, 2, 0, 2, 0, 2, 0, 2},
	}

	return newBoard

}
func (b *Board) ReadBackgroundBoard() {

	for _, val := range b.BackgroundBoard {
		fmt.Println(val)
	}

}
func (b *Board) ReadPlayerBoard() {

	for _, val := range b.PlayerBoard {

		fmt.Println(val)

	}

}

// select board

func (b *Board) SelectBoard(x int, y int, prev *Coordinates) {
	//player
	//right
	prev.x = x
	prev.y = y

	if y-1 >= 0 && x+1 <= 7 {
		if b.PlayerBoard[y-1][x+1] == 0 {
			b.PlayerBoard[y-1][x+1] = 3
		} else if y-2 >= 0 && x+2 <= 7 {
			if b.PlayerBoard[y-2][x+2] == 0 {
				if y-3 >= 0 && x+3 <= 7 {

					if (b.PlayerBoard[y-3][x+3] == 1 || b.PlayerBoard[y-3][x+3] == 2) && b.PlayerBoard[y-4][x+4] == 0 {

						b.PlayerBoard[y-4][x+4] = 3
					}
				}
				b.PlayerBoard[y-2][x+2] = 3
			}

		}

	} //left

	if y-1 >= 0 && x-1 >= 0 {
		if b.PlayerBoard[y-1][x-1] == 0 {
			b.PlayerBoard[y-1][x-1] = 3
		} else if y-2 >= 0 && x-2 >= 0 {
			if b.PlayerBoard[y-2][x-2] == 0 {

				if y-3 >= 0 && x-3 >= 0 {
					if (b.PlayerBoard[y-3][x-3] == 1 || b.PlayerBoard[y-3][x-3] == 2) && b.PlayerBoard[y-4][x-4] == 0 {

						b.PlayerBoard[y-4][x-4] = 3
					}
				}
				b.PlayerBoard[y-2][x-2] = 3
			}
		}

	}
}
func (b *Board) ClearBoard() {
	for y, arr := range b.PlayerBoard {
		for x, val := range arr {
			if val == 3 {
				b.PlayerBoard[y][x] = 0
			}
		}

	}
}

func (b *Board) MovePiece(x int, y int, prev *Coordinates) {
	// check if the place
	fmt.Println(prev.x)
	if prev.x != -1 && prev.y != -1 {

		if x > prev.x {
			// right

			if y-1 >= 0 && x+1 <= 7 {
				if b.PlayerBoard[prev.y-1][prev.x+1] == 1 {
					b.PlayerBoard[prev.y-1][prev.x+1] = 0
				}
			}
		} else if x < prev.x {
			// left

			if y-1 >= 0 && x-1 >= 0 {
				if b.PlayerBoard[prev.y-1][prev.x-1] == 1 {
					b.PlayerBoard[prev.y-1][prev.x-1] = 0
				}
			}
		} else {
			log.Fatal("MOVE PIECE FALED")
		}
		b.PlayerBoard[y][x] = 2

		b.PlayerBoard[prev.y][prev.x] = 0
		for prev.x > x && prev.y > y {
			fmt.Println(prev.y, prev.x)
			if b.PlayerBoard[prev.y][prev.x] != 2 {
				b.PlayerBoard[prev.y][prev.x] = 0
			}
			prev.y -= 1
			prev.x -= 1
		}
		b.ClearBoard()
	}
}

// emeny
func (b *Board) EnemyMove() {
	fmt.Println("Enemy move Begin")
	allEnemyCoords := []Coordinates{}

	for y, arr := range b.PlayerBoard {
		for x, _ := range arr {
			if b.PlayerBoard[y][x] == 1 {
				allEnemyCoords = append(allEnemyCoords, Coordinates{x: x, y: y})
			}

			if y+2 <= 7 && x-2 >= 0 {
				if b.PlayerBoard[y][x] == 1 && b.PlayerBoard[y+1][x-1] == 2 && b.PlayerBoard[y+2][x-2] == 0 {
					//there is a player piece on the right
					fmt.Println("right  hitt")
					b.PlayerBoard[y][x] = 0
					b.PlayerBoard[y+1][x-1] = 0
					b.PlayerBoard[y+2][x-2] = 1

					allEnemyCoords = []Coordinates{}
					return
				}
			}

			if y+2 <= 7 && x+2 <= 7 {
				if b.PlayerBoard[y][x] == 1 && b.PlayerBoard[y+1][x+1] == 2 && b.PlayerBoard[y+2][x+2] == 0 {
					//there is a player piece on the  left

					fmt.Println("left hitt")
					b.PlayerBoard[y][x] = 0
					b.PlayerBoard[y+1][x+1] = 0
					b.PlayerBoard[y+2][x+2] = 1

					allEnemyCoords = []Coordinates{}
					return
				}
			}
		}

	}

  randPiece := Coordinates{}

	if len(allEnemyCoords) > 1 {

		randPiece = allEnemyCoords[rand.IntN(len(allEnemyCoords))]
	} else {
		randPiece = allEnemyCoords[0]
	}

	randDir := rand.IntN(2)

	//	fmt.Println("random dir", randDir)

	//	fmt.Println("random piece ", randPiece)
	if randDir == 0 {
		//left
		if randPiece.x-1 >= 0 && randPiece.y+1 <= 7 {
			if b.PlayerBoard[randPiece.y+1][randPiece.x-1] == 0 {
				fmt.Println("should be placed 0")
				b.PlayerBoard[randPiece.y][randPiece.x] = 0
				b.PlayerBoard[randPiece.y+1][randPiece.x-1] = 1
				allEnemyCoords = []Coordinates{}
				return

			} else if randPiece.y+2 <= 7 && randPiece.x-2 >= 0 {
				if b.PlayerBoard[randPiece.y+2][randPiece.x-2] == 0 {
					/*
						      // idk if this should be included ..
										if y-3 >= 0 && x-3 >= 0 {
											if (b.PlayerBoard[y-3][x-3] == 1 || b.PlayerBoard[y-3][x-3] == 2) && b.PlayerBoard[y-4][x-4] == 0 {

												b.PlayerBoard[y-4][x-4] = 3
											}
										}
					*/

					fmt.Println("should be jupmped 0")
					b.PlayerBoard[randPiece.y][randPiece.x] = 0
					b.PlayerBoard[randPiece.y+2][randPiece.x-2] = 1
					allEnemyCoords = []Coordinates{}
					return
				}
			}

		}

	} else {
		//right
		if randPiece.x+1 <= 7 && randPiece.y+1 <= 7 {
			if b.PlayerBoard[randPiece.y+1][randPiece.x+1] == 0 {

				fmt.Println("should be placed 1")
				b.PlayerBoard[randPiece.y][randPiece.x] = 0
				b.PlayerBoard[randPiece.y+1][randPiece.x+1] = 1

				allEnemyCoords = []Coordinates{}
				return
			} else if randPiece.y+2 <= 7 && randPiece.x+2 <= 7 {
				if b.PlayerBoard[randPiece.y+2][randPiece.x+2] == 0 {
					/*
						      // idk if this should be included ..
										if y-3 >= 0 && x-3 >= 0
											if (b.PlayerBoard[y-3][x-3] == 1 || b.PlayerBoard[y-3][x-3] == 2) && b.PlayerBoard[y-4][x-4] == 0 {

												b.PlayerBoard[y-4][x-4] = 3
											}
										}
					*/

					fmt.Println("should be jupmped 1")
					b.PlayerBoard[randPiece.y][randPiece.x] = 0
					b.PlayerBoard[randPiece.y+2][randPiece.x+2] = 1
					allEnemyCoords = []Coordinates{}
					return

				}
			}

		}

	}

	fmt.Println("Faled")
	b.EnemyMove()
}
