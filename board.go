package main

import (
	"fmt"
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
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
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
	if b.PlayerBoard[y][x] == 5 {
		//king
		fmt.Println("KING")
		tempX := x
		tempY := y
		for tempX > 0 && tempY > 0 {
			tempX -= 1
			tempY -= 1
			if b.PlayerBoard[tempY][tempX] == 0 {
				b.PlayerBoard[tempY][tempX] = 3
			}
		}

		temp2X := x
		temp2Y := y
		for temp2X < 7 && temp2Y > 0 {
			temp2X += 1
			temp2Y -= 1
			if b.PlayerBoard[temp2Y][temp2X] == 0 {
				b.PlayerBoard[temp2Y][temp2X] = 3
			}

		}

		temp3X := x
		temp3Y := y
		for temp3X > 0 && temp3Y < 7 {
			temp3X -= 1
			temp3Y += 1
			if b.PlayerBoard[temp3Y][temp3X] == 0 {
				b.PlayerBoard[temp3Y][temp3X] = 3
			}

		}
		temp4X := x
		temp4Y := y
		for temp4X < 7 && temp4Y < 7 {
			temp4X += 1
			temp4Y += 1
			if b.PlayerBoard[temp4Y][temp4X] == 0 {
				b.PlayerBoard[temp4Y][temp4X] = 3
			}

		}
	} else {
		//pawn
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
		/*
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
			}*/
		if b.PlayerBoard[prev.y][prev.x] == 5 {

			b.PlayerBoard[y][x] = 5
		} else {

			b.PlayerBoard[y][x] = 2
		}

		b.PlayerBoard[prev.y][prev.x] = 0
		for prev.x > x && prev.y > y {
			fmt.Println(prev.y, prev.x)
			if b.PlayerBoard[prev.y][prev.x] != 2 && b.PlayerBoard[prev.y][prev.y] != 5 {
				b.PlayerBoard[prev.y][prev.x] = 0
			}
			prev.y -= 1
			prev.x -= 1
		}
		for prev.x < x && prev.y > y {
			fmt.Println(prev.y, prev.x)
			if b.PlayerBoard[prev.y][prev.x] != 2 && b.PlayerBoard[prev.y][prev.y] != 5 {
				b.PlayerBoard[prev.y][prev.x] = 0
			}
			prev.y -= 1
			prev.x += 1
		}
		for prev.x > x && prev.y < y {
			fmt.Println(prev.y, prev.x)
			if b.PlayerBoard[prev.y][prev.x] != 2 && b.PlayerBoard[prev.y][prev.y] != 5 {
				b.PlayerBoard[prev.y][prev.x] = 0
			}
			prev.y += 1
			prev.x -= 1
		}
		for prev.x < x && prev.y < y {
			fmt.Println(prev.y, prev.x)
			if b.PlayerBoard[prev.y][prev.x] != 2 && b.PlayerBoard[prev.y][prev.y] != 5 {
				b.PlayerBoard[prev.y][prev.x] = 0
			}
			prev.y += 1
			prev.x += 1
		}

		b.ClearBoard()
	}
}

// emeny
func (b *Board) EnemyMove() {
	allEnemyCoords := []Coordinates{}
	allEnemyKingis := []Coordinates{}

	for y, arr := range b.PlayerBoard {
		for x, _ := range arr {
			if b.PlayerBoard[y][x] == 1 {
				allEnemyCoords = append(allEnemyCoords, Coordinates{x: x, y: y})
			}
			if b.PlayerBoard[y][x] == 4 {
				allEnemyCoords = append(allEnemyCoords, Coordinates{x: x, y: y})
			}
			// check the easy hits in pawns
			if y+2 <= 7 && x-2 >= 0 {
				if b.PlayerBoard[y][x] == 1 && b.PlayerBoard[y+1][x-1] == 2 && b.PlayerBoard[y+2][x-2] == 0 {
					//there is a player piece on the right
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
	checkKing := rand.IntN(2)

	if checkKing == 0 {
		if len(allEnemyKingis) > 0 {
			randKing := allEnemyKingis[rand.IntN(len(allEnemyKingis))]
			randKingDir := rand.IntN(4)

			switch randKingDir {
			case 0:
				//b-l

				temp := false
				tempX := randKing.x
				tempY := randKing.y
				for {
					if tempY <= 7 && tempX >= 0 {
						tempX -= 1
						tempY += 1
						if b.PlayerBoard[tempY][tempY] == 1 {
							temp = true
						}

						if b.PlayerBoard[tempY][tempY] == 0 && temp {
							b.PlayerBoard[tempY][tempX] = 4
							b.PlayerBoard[randKing.y][randKing.x] = 0
							break
						}
					} else {
						break
					}
				}
				return
			case 1:
				//b-r
				temp := false
				tempX := randKing.x
				tempY := randKing.y
				for {
					if tempY <= 7 && tempX <= 7 {
						tempX += 1
						tempY += 1
						if b.PlayerBoard[tempY][tempY] == 1 {
							temp = true
						}

						if b.PlayerBoard[tempY][tempY] == 0 && temp {
							b.PlayerBoard[tempY][tempX] = 4
							b.PlayerBoard[randKing.y][randKing.x] = 0
							break
						}
					} else {
						break
					}
				}
				return

			case 2:
				//t-l
				temp := false
				tempX := randKing.x
				tempY := randKing.y
				for {
					if tempY >= 0 && tempX >= 0 {
						tempX -= 1
						tempY -= 1
						if b.PlayerBoard[tempY][tempY] == 1 {
							temp = true
						}

						if b.PlayerBoard[tempY][tempY] == 0 && temp {
							b.PlayerBoard[tempY][tempX] = 4
							b.PlayerBoard[randKing.y][randKing.x] = 0
							break
						}
					} else {
						break
					}
				}
				return

			case 3:
				//t-r
				temp := false
				tempX := randKing.x
				tempY := randKing.y
				for {
					if tempY >= 0  && tempX <= 7 {
						tempX += 1
						tempY -= 1
						if b.PlayerBoard[tempY][tempY] == 1 {
							temp = true
						}

						if b.PlayerBoard[tempY][tempY] == 0 && temp {
							b.PlayerBoard[tempY][tempX] = 4
							b.PlayerBoard[randKing.y][randKing.x] = 0
							break
						}
					} else {
						break
					}
				}
				return

			}

		}
	}

	//	fmt.Println("random dir", randDir)

	//	fmt.Println("random piece ", randPiece)
	if randDir == 0 {
		if randPiece.x-1 >= 0 && randPiece.y+1 <= 7 {
			if b.PlayerBoard[randPiece.y+1][randPiece.x-1] == 0 {
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

					b.PlayerBoard[randPiece.y][randPiece.x] = 0
					b.PlayerBoard[randPiece.y+2][randPiece.x-2] = 1
					allEnemyCoords = []Coordinates{}
					return
				}
			}

		}

	} else {
		if randPiece.x+1 <= 7 && randPiece.y+1 <= 7 {
			if b.PlayerBoard[randPiece.y+1][randPiece.x+1] == 0 {

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

					b.PlayerBoard[randPiece.y][randPiece.x] = 0
					b.PlayerBoard[randPiece.y+2][randPiece.x+2] = 1
					allEnemyCoords = []Coordinates{}
					return

				}
			}

		}

	}
	//fix some errors here
	//when the last piece is at the end of the board than this function locks in a loop
	//temporary solution ..
	if len(allEnemyCoords) == 0 {
		fmt.Println("___________________")
		fmt.Println(randDir, randPiece)
		fmt.Println("Enemy turn skipped ")
		return

	}

	fmt.Println("Faled")
	b.EnemyMove()
}

func (b *Board) CountPieces() (EnemyCount int, PlayerCount int) {
	numE := 0
	numP := 0
	for _, arr := range b.PlayerBoard {
		for _, val := range arr {
			if val == 1 || val == 4 {
				numE += 1
			}
			if val == 2 || val == 5 {

				numP += 1
			}
		}

	}
	return numE, numP
}

func (b *Board) CheckForKing() {
	for x, val := range b.PlayerBoard[0] {
		if val == 2 {
			b.PlayerBoard[0][x] = 5

		}
	}
	for x, val := range b.PlayerBoard[7] {
		if val == 1 {
			b.PlayerBoard[0][x] = 4

		}
	}

}
