package main

import (
	"fmt"
	"time"
)

func print_map(b [10][10]int) {
	print("   ")
	for i := 0; i <= 9; i++ {
		print(" ", i)
	}
	print("\n")
	print("  -----------------------\n")
	for i := 0; i <= 9; i++ {
		print(i, " |")
		for j := 0; j <= 9; j++ {
			if b[i][j] == 0 {
				print(" +")
			} else if b[i][j] == 1 {
				print(" 0")
			} else if b[i][j] == 2 {
				print(" @")
			} else {
				print("ERROR")
			}
		}
		print(" |\n")
	}
	print("  -----------------------\n")

}

func win_check(b [10][10]int, x int, y int) int {
	var last_stone int = b[x][y]
	var start_x, start_y, end_x, end_y int = x, y, x, y

	for start_x-1 >= 0 && b[start_x-1][y] == last_stone {
		start_x--
	}
	for end_x+1 <= 9 && b[end_x+1][y] == last_stone {
		end_x++
	}
	if end_x-start_x+1 >= 5 {
		return last_stone
	}

	start_x = x
	end_x = x
	for start_y-1 >= 0 && b[x][start_y-1] == last_stone {
		start_y--
	}
	for end_y+1 <= 9 && b[x][end_y+1] == last_stone {
		end_y++
	}
	if end_y-start_y+1 >= 5 {
		return last_stone
	}

	start_y = y
	end_y = y

	for start_x-1 >= 0 && start_y-1 >= 0 && b[start_x-1][start_y-1] == last_stone {
		start_x--
		start_y--
	}
	for end_x+1 <= 9 && end_y+1 <= 9 && b[end_x+1][end_y+1] == last_stone {
		end_x++
		end_y++
	}
	if end_x-start_x+1 >= 5 {
		return last_stone
	}

	start_x = x
	start_y = y
	end_x = x
	end_y = y
	for start_x-1 >= 0 && end_y+1 <= 9 && b[start_x-1][end_y+1] == last_stone {
		start_x--
		end_y++
	}
	for end_x+1 <= 9 && start_y-1 >= 0 && b[end_x+1][start_y-1] == last_stone {
		end_x++
		start_y--
	}
	if end_x-start_x+1 >= 5 {
		return last_stone
	}

	return 0
}
func main() {
	var board [10][10]int
	var x int
	var y int
	var turn int

	print_map(board)
	turn = 1
	for {
		print("please enter \"x y\" coordinate >>")
		fmt.Scanln(&x, &y)
		if x < 0 || y < 0 || x >= 10 || y >= 10 {
			println("error, out of bound!")
			time.Sleep(2 * time.Second)
			continue
		} else if board[x][y] != 0 {
			println("error, already used!")
			time.Sleep(2 * time.Second)
			continue
		} else {
			if turn%2 == 1 {
				board[x][y] = 1
			} else {
				board[x][y] = 2
			}
		}
		print_map(board)

		if win_check(board, x, y) == 1 {
			println("player 1 wins!")
			break
		} else if win_check(board, x, y) == 2 {
			println("player 2 wins!")
			break
		}

		if turn == 100 {
			println("draw!")
			break
		}
		turn = turn + 1
	}
}
