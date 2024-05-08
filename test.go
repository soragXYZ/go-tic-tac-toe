package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	sizeTab int    = 3
	player1 string = "X"
	player2 string = "O"
)

// checkWin returns a bool, if a player has won or not
func checkWIn(symbol string, tab [sizeTab][sizeTab]string) bool {

	if // columns
	tab[0][0] == symbol &&
		tab[0][1] == symbol &&
		tab[0][2] == symbol {
		return true
	} else if tab[1][0] == symbol &&
		tab[1][1] == symbol &&
		tab[1][2] == symbol {
		return true
	} else if tab[2][0] == symbol &&
		tab[2][1] == symbol &&
		tab[2][2] == symbol {
		return true
	} else if // lines
	tab[0][0] == symbol &&
		tab[1][0] == symbol &&
		tab[2][0] == symbol {
		return true
	} else if tab[0][1] == symbol &&
		tab[1][1] == symbol &&
		tab[2][1] == symbol {
		return true
	} else if tab[0][2] == symbol &&
		tab[1][2] == symbol &&
		tab[2][2] == symbol {
		return true
	} else if // diagonals
	tab[0][0] == symbol &&
		tab[1][1] == symbol &&
		tab[2][2] == symbol {
		return true
	} else if tab[0][2] == symbol &&
		tab[1][1] == symbol &&
		tab[2][0] == symbol {
		return true
	} else {
		return false
	}
}

func main() {

	// initialize
	var tab [sizeTab][sizeTab]string
	for i := 0; i < sizeTab; i++ {
		for j := 0; j < sizeTab; j++ {
			tab[i][j] = "_"
		}
	}
	var turn int

	// play until draw
	for turn < sizeTab*sizeTab {

		var player string

		if turn%2 == 0 {
			player = player1
		} else {
			player = player2
		}

		fmt.Println("Player", player, ", your turn")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entrez un nombre entier : ")
		scanner.Scan()

		nbr, err := strconv.Atoi(scanner.Text()) // convert input as int

		// verify user input
		if err != nil || nbr >= sizeTab*sizeTab || nbr < 0 {
			fmt.Println("Fail, try again")
		}

		var line int = nbr % sizeTab
		var column int = nbr / sizeTab

		// verify if box is empty
		if tab[line][column] != "_" {
			fmt.Println("Incorrect, try again")
			continue
		}

		tab[line][column] = player
		for _, i := range tab {
			fmt.Println(i)
		}

		// check if victory
		var victory bool = checkWIn(player, tab)
		if victory {
			fmt.Println("Player", player, "has won")
			break
		}

		turn++
	}
	fmt.Println("thanks for playing")

}
