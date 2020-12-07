package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"quizgame/quiz"
)

func main() {

	quizFile := flag.String("quiz", "problems.csv", "Name of quiz file")
	passingScore := flag.Int("pass", 80, "Passing score")
	flag.Parse()
	q := quiz.New()
	records, err := q.PrepareQuiz(*quizFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	fmt.Println("-----------BEGIN--------------")
	correct, total := q.TakeQuiz(records)
	pass, score := q.Grade(correct, total, *passingScore)
	fmt.Println("------------END---------------")
	fmt.Printf("Total Questions: %d\n", int(total))
	fmt.Printf("Correct Answers: %d\n", int(correct))
	fmt.Printf("Passing Score: %d %%\n", *passingScore)
	fmt.Printf("Score: %.2f %%\n", math.RoundToEven(score))
	fmt.Printf("Pass: %t\n", pass)
}
