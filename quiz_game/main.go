package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv",
		"a csv file in the format of \"question, answer\"")
	timePtr := flag.Int("limit", 30,
		"the time limit for the quiz in seconds")
	flag.Parse()
	fmt.Println(*timePtr)
	timer1 := time.NewTimer(time.Duration(*timePtr) * time.Second)
	records := readCsvFile(*csvPtr)
	problems := castToProblems(records)
	numberOfCorrectAnswers := showQuestionGetAnswerCount(*timer1, problems)
	fmt.Printf("\nYou scored %d out of %d\n", numberOfCorrectAnswers, len(problems))
}