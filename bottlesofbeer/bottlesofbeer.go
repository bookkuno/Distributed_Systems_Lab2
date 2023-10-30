package main

import (
	"flag"
	"fmt"
	"sync"
)

var nextAddr string
var wg sync.WaitGroup
var bottleCount int
var songOver chan bool

func main() {
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for the next member of the round.")
	flag.IntVar(&bottleCount, "n", 0, "Bottles of Beer (launches the song if not 0)")
	flag.Parse()

	songOver = make(chan bool)

	if bottleCount == 0 {
		fmt.Println("No bottles to sing, exiting...")
		return
	}

	if bottleCount < 0 {
		fmt.Println("Invalid bottle count, exiting...")
		return
	}

	// Start the song from the last buddy
	if nextAddr == "localhost:8040" {
		go startSinging("Buddy 1", bottleCount)
	} else {
		go listenForSong("Buddy 1")
	}

	wg.Add(1) // Wait for the last buddy to finish the song
	<-songOver
}

func startSinging(name string, bottles int) {
	defer wg.Done()

	if bottles == 0 {
		songOver <- true
		return
	}

	// Sing a verse
	fmt.Printf("%s: %d bottles of beer on the wall, %d bottles of beer.\n", name, bottles, bottles)
	fmt.Printf("%s: Take one down, pass it around...\n", name)

	// Create the next buddy and pass the bottles
	nextBuddy := "Buddy 1"
	if name != "Buddy 3" {
		nextBuddy = "Buddy " + fmt.Sprint(parseBuddyNumber(name)+1)
	}
	go startSinging(nextBuddy, bottles-1)
}

func listenForSong(name string) {
	defer wg.Done()

	// Listen for the song from the previous buddy
	startSinging(name, bottleCount)
}

func parseBuddyNumber(name string) int {
	// Extract the buddy number from the name
	// "Buddy 1" becomes 1
	var num int
	_, err := fmt.Sscanf(name, "Buddy %d", &num)
	if err != nil {
		num = -1 // Invalid buddy name
	}
	return num
}
