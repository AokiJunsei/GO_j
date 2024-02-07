package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	//"strings"
	"time"
)

var choices = []string{"rock", "paper", "scissors"}

func handleMove(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm()
	playerChoice := r.Form.Get("choice")

	
	rand.Seed(time.Now().UnixNano())
	serverChoice := choices[rand.Intn(len(choices))]

	
	result := ""
	if playerChoice == serverChoice {
		result = "It's a draw!"
	} else if (playerChoice == "rock" && serverChoice == "scissors") ||
		(playerChoice == "paper" && serverChoice == "rock") ||
		(playerChoice == "scissors" && serverChoice == "paper") {
		result = "You win!"
	} else {
		result = "You lose!"
	}

	
	fmt.Fprintf(w, "Server choice: %s\nPlayer choice: %s\nResult: %s", serverChoice, playerChoice, result)
}

func main() {
	
	http.HandleFunc("/move", handleMove)

	
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
