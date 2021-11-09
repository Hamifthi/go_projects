package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)


func readCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Unable to open specified file")
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read csv content")
	}
	return records
}


func castToProblems(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{
			question: record[0],
			answer: strings.TrimSpace(record[1]),
		}
	}
	return problems
}


func readStdIn(answerCh chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	enteredAnswer, err := reader.ReadString('\n')
	enteredAnswer = strings.TrimSpace(strings.Replace(
		enteredAnswer, "\n", "", 1))
	if err != nil {
		log.Fatalf("can't get answer from you")
	}
	answerCh <- enteredAnswer
}


func showQuestionGetAnswerCount(timer time.Timer, problems []problem) int {
	numberOfCorrectAnswers := 0
	answerCh := make(chan string)
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i, problem.question)
		go readStdIn(answerCh)
		select {
			case <-timer.C:
				return numberOfCorrectAnswers
			case answer := <-answerCh:
				if answer == problem.answer {
					numberOfCorrectAnswers++
				}
		}
	}
	return numberOfCorrectAnswers
}