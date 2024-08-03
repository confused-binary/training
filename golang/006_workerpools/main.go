package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func worker(profile string, command string, results chan<- string) {
	fmt.Println("aws --profile", profile, command)
	n := rand.Intn((10-3)+1) + 3
	fmt.Printf("%s sleeping %d seconds...\n", profile, n)
	time.Sleep(time.Duration(n) * time.Second)
	results <- profile + " done!"
}

func main() {
	profiles := []string{"profile1", "profile2"}
	argCommand := strings.Join(os.Args[1:], " ")

	command := make(chan string, len(profiles))
	results := make(chan string, len(profiles))

	for _, profile := range profiles {
		go worker(profile, argCommand, results)
	}

	command <- argCommand
	close(command)

	for i := 1; i <= len(profiles); i++ {
		x := <-results
		fmt.Println(x)
	}
}
