package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	file, err := os.Open("problems.csv")

	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)
	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()
	// Checks for the error
	if err != nil {
		fmt.Println("Error reading records")
	}

	timer1 := time.NewTimer(1 * time.Second)

	correct := 0

	// Loop to iterate through
	// and print each of the string slice
	for _, eachrecord := range records {
		fmt.Println(eachrecord[0])

		answerCh := make(chan string)
		go func() {
			// var then variable name then variable type
			var answer string

			// Taking input from user
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer1.C:
			fmt.Println("Time's up!")
			return
		case answer := <-answerCh:
			if answer == eachrecord[1] {
				correct++
			}
		}

	}

	fmt.Println("You got", correct, "out of", len(records))
}
