package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// csvFilename is a pointer to a string
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
	timeLimit := flag.Int("limit", 30, "the time limit for your quiz in seconds")

	// parse command line arguments
	flag.Parse()

	// open file
	file, err := os.Open(*csvFilename)
	if err != nil {

		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))

	}

	// read the csv file - NewReader takes in an io.Reader
	r := csv.NewReader(file)
	// parse the csv file
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the csv file: %s\n", *csvFilename))
	}

	// fmt.Println(lines)
	problems := parseLines(lines)
	// fmt.Println(problems)

	// declare a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// initialize score counter
	correct := 0

	// define a label
problemLoop:
	// print problems to user and get response
	for i, p := range problems {
		answerChan := make(chan string)
		fmt.Printf("problem #%d: %s = \n", i+1, p.q)

		// spawn a new go routine to handle blocking code [fmt.Scanf]
		go func() {
			// check answer
			var answer string
			fmt.Scanf("%s\n", &answer) // Scanf scans text from std input and we store the value in answer with a pointer
			// send answer to channel
			answerChan <- answer
		}()

		// use a select statement to check when we have a message from the channel
		select {
		case <-timer.C:
			fmt.Println("\nTime up!")
			break problemLoop
		case answer := <-answerChan:
			if answer == p.a {
				// fmt.Println("correct!")
				correct++
			}
		}
	}

	// pritn final score
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))

	// loop through lines and fill the slice
	for i, line := range lines {
		res[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return res
}
