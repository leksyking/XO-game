package main

import "fmt"

func main(){
	fmt.Println("Welcome to my Game")
	//setup the empty board array
	xoBoard := [3][3]string{}
	
	//setup the first player
	player := "x"
	
for {
	fmt.Println("Current Player is ", player)
	// row value
	var row int
	fmt.Println("Enter the row value:")
	fmt.Scanln(&row)

	if row < 0 || row > 2 {
		fmt.Println("Invalid entry, Please enter new number, 0, 1 or 2")
		continue
	}

	// column value
	var column int
	fmt.Println("Enter the column value:")
	fmt.Scanln(&column)

	if column < 0 || column > 2 {
		fmt.Println("Invalid entry, Please enter new number, 0, 1 or 2")
		continue
	}
	// set value into the board
	if xoBoard[row][column] == "" {
		xoBoard[row][column] = player
	}else {
		fmt.Println("Invalid Entry!, row and column occupied already")
	}

	//check if someone won
	for i := 0; i < 3; i++ {
			//check rows || columns
			if (xoBoard[i][0]==xoBoard[i][1]&&xoBoard[i][1]==xoBoard[i][2]&&xoBoard[i][2]== player || xoBoard[0][i]==xoBoard[1][i]&&xoBoard[1][i]==xoBoard[2][i]&&xoBoard[2][i]== player ){
				fmt.Println("Game is ended, winner is player:", player)
				return
			}
		}

		if 	(xoBoard[0][0] == xoBoard[1][1]&&xoBoard[1][1] == xoBoard[2][2]&&xoBoard[2][2] == player ||
			xoBoard[0][2] == xoBoard[1][1]&&xoBoard[1][1] == xoBoard[2][0]&&xoBoard[2][0] == player){
				fmt.Println("Game is ended, winner is player:", player)
				return
			}

	//display Board 
	fmt.Println(xoBoard[0])
	fmt.Println(xoBoard[1])
	fmt.Println(xoBoard[2]) 
	
	//swap players at the end of each player's session
	if player == "x"{
		player = "o"
	}else{
		player = "x"
	}
	}		
}