package main
/*
Name : M Furqon Fahlevi
SID	 : 1301194214
*/
import (
	"fmt"
	"math/rand"
	"time"
)

const N = 20

var player[N] int
var dealer[N] int

func main() {
	var playerdice, playersum, dealerdice, dealersum, keep, points, diceindex int

	welcomemessage()
	instruction()

	rand.Seed(time.Now().UnixNano())
	fmt.Print("Player dice\t: ")
	i := 0
	for i < 7 {
		playerdice = throwDice()
		fmt.Print(playerdice, " ")
		player[i+1] = playerdice
		playersum = playersum + playerdice
		i++
	}
	fmt.Println("")
	fmt.Println("Player sum\t:" ,playersum)
	fmt.Println("")
	fmt.Print("Dealer dice\t: ")
	j := 0
	for j < 7 &&  dealersum < 18 {
		dealerdice = throwDice()
		fmt.Print(dealerdice, " ")
		dealer[j+1] = dealerdice
		dealersum = dealersum + dealerdice
		j++
	}
	fmt.Println()
	fmt.Println("Dealer sum\t:" ,dealersum)
	fmt.Println("\nDo you want to keep all the dice : 1. Yes 2. No (Input the number)")
	fmt.Scanln(&keep)
	if keep == 2 {
		fmt.Print("Choose which dice : (end your choice with 0) \n")
		for k := 1; k <= 7; k++ {
			fmt.Printf("Dice (%d) = %d\n", k, player[k])
		}
		playersum = 0
		fmt.Scan(&diceindex)
		for diceindex != 0 {
			playersum = playersum + player[diceindex]
			fmt.Scan(&diceindex)
		}
		countPoints(playersum, dealersum, &points)
		fmt.Printf("Result : the player choose which dice he/she wants to use = %d, the dealer dice = %d, so the dealer lost. The player gets %d points", playersum, dealersum, points)
	} else {
		countPoints(playersum, dealersum, &points)
		fmt.Printf("Result : the player keeps all the dice = %d, the dealer dice = %d, so the dealer lost. The player gets %d points", playersum, dealersum, points)
	}	
}

func throwDice() int {
	return rand.Intn(6) + 1
}

func countPoints(playersum, dealersum int, points *int) {
	if playersum == 21 && dealersum > 21 {
		*points = 150
	} else if playersum == 21 && dealersum < 21 {
		*points = 100
	} else if  playersum == 21 && dealersum == 21 {
		*points = 70
	} else if playersum < 21 && dealersum > 21 {
		*points = 30
	} else if playersum > dealersum && playersum < 21 {
		*points = 10
	}
}

func WinorNot(points int) string {
	var win string
	if points == 0 {
		win = "lost"
	} else {
		win = "win"
	}
	return win
}

func welcomemessage() {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------Welcome to Easy 21 Game! :D---------------------")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("----------------MADE BY : M FURQON FAHLEVI----------------------")
}

func instruction() {
	fmt.Println("The Play :")
	fmt.Println("Begin by the player rolls 7 dice and select which dice you will keep to achive no more than 21.")
	fmt.Println("You can keep the value of the dice or choosing them")
	fmt.Println("")
	fmt.Println("How to win the game : ")
	fmt.Println("1. The Player wins 150 points if your dice equals 21 and The Dealer is above 21.")
	fmt.Println("2. The Player wins 100 points if your dice equals 21 and The Dealer is less than 21.")
	fmt.Println("3. The Player wins 70 points if both are equals 21.")
	fmt.Println("4. The Player wins 30 points if your dice is less than 21 and The Dealer is above 21.")
	fmt.Println("5. The Player wins 10 points if your dice above The Dealer dice and less than 21.")
	fmt.Println("You won't get any points if you aren't follow the instruction.")
	fmt.Println("")
}