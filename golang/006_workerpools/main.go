package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func worker(profile string, command <-chan string, results chan<- string) {
	// for j := range jobs {
	// cmd := strings.Join(command, " ")
	cmd := command
	fmt.Println("Profile", profile, "started job -", cmd)
	time.Sleep(time.Second)
	fmt.Println("Profile", profile, "finished job -", cmd)
	results <- "done!"
	// }
}

func main() {
	profiles := []string{"profile1", "profile2"}
	argCommand := strings.Join(os.Args[1:], " ")
	command := make(chan string, len(profiles))
	results := make(chan string, len(profiles))

	for _, profile := range profiles {
		go worker(profile, command, results)
	}

	command <- argCommand
	close(command)

	for i := 1; i <= len(profiles); i++ {
		<-results
	}

	// for j := 1; j <= numJobs; j++ {
	// 	jobs <- j
	// }
	// close(jobs)

	// for a := 1; a <= numJobs; a++ {
	// 	<-results
	// }
}
