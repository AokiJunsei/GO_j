package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func sendMove(choice string) {
	values := url.Values{}
	values.Set("choice", choice)

	resp, err := http.PostForm("http://localhost:8080/move", values)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("じゃんけんを始めます")

	for {
		fmt.Print("Enter your choice (rock, paper, scissors): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		
		if choice != "rock" && choice != "paper" && choice != "scissors" {
			fmt.Println("Invalid choice. Please enter again.")
			continue
		}

		
		sendMove(choice)
	}
}
