package main

import (
	"flag"
	"fmt"
)

func main() {
	quizType := flag.String("type", "arithmetic", "Type of quiz")
	quizFile := flag.String("file", "problems.csv", "Name of quiz file")
	passingScore := flag.Int("pass", 80, "Passing score")
	// timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	fmt.Println(*quizType)
	fmt.Println(*quizFile)
	fmt.Println(*passingScore)

	director := Director{}

	a := &ArithmeticBuilder{}
	director.SetBuilder(a)
	director.Construct()

}
