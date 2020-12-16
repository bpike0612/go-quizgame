package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"
)

type QuizBuilder interface {
	PrepareQuiz() QuizBuilder
	AskQuestions() QuizBuilder
	Grade() QuizBuilder
	GiveResults() QuizBuilder
	GetFileName() QuizProduct
	SetFileName() QuizBuilder
	GetPassingScore() QuizProduct
	SetPassingScore() QuizBuilder
	SetPass(bool) QuizProduct
	GetQuiz() QuizProduct
}

//Director
type Director struct {
	builder QuizBuilder
}

func (d *Director) Construct() {
	d.builder.SetPassingScore().SetFileName().PrepareQuiz().AskQuestions().Grade().GiveResults()
}

func (d *Director) SetBuilder(b QuizBuilder) {
	d.builder = b
}

//Product
type QuizProduct struct {
	FileName           string
	Problems           map[string]string
	PassingScore       int
	CorrectAnswerTotal float64
	TotalQuestions     float64
	TestScore          float64
	Pass               bool
}

//A Builder of type arithmetic
type ArithmeticBuilder struct {
	q QuizProduct
}

func (a *ArithmeticBuilder) GetPassingScore() QuizProduct {
	return a.q
}

func (a *ArithmeticBuilder) SetPassingScore() QuizBuilder {
	a.q.PassingScore = 80
	return a
}

func (a *ArithmeticBuilder) GetFileName() QuizProduct {
	return a.q
}

func (a *ArithmeticBuilder) SetFileName() QuizBuilder {
	a.q.FileName = "problems.csv"
	return a
}

func (a *ArithmeticBuilder) GetQuiz() QuizProduct {
	return a.q
}

func (a *ArithmeticBuilder) SetPass(pass bool) QuizProduct {
	a.q.Pass = pass
	return a.q
}

// PrepareQuiz opens and parses the quiz file.
func (a *ArithmeticBuilder) PrepareQuiz() QuizBuilder {
	//var records []string
	var m map[string]string
	m = make(map[string]string)

	// Open the quizName
	csvfile, err := os.Open(a.q.FileName)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the quizName
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()

		// records = append(records, record...)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m[record[0]] = record[1]
		a.q.Problems = m
		//a.q.Problems[record[0]] = record[1]
	}
	//return q.problems, err
	//return err
	return a
}

// AskQuestions takes a map of the parsed csv file and interacts with the user,
// displaying each question to the terminal allowing the user to respond.
// Returns correctAnswerTotal and totalQuestions
func (a *ArithmeticBuilder) AskQuestions() QuizBuilder {
	fmt.Println("-----------BEGIN--------------")
	//var correct float64
	timer := time.NewTimer(30 * time.Second)

	// TODO: Refactor this to be easier for testing bpike_20201215
	// label
problemloop:
	for k, v := range a.q.Problems {
		fmt.Printf("%s=", k)
		// channel
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			a.q.TotalQuestions = float64(len(a.q.Problems))
			return a
			break problemloop
		case answer := <-answerCh:
			if v == answer {
				a.q.CorrectAnswerTotal++
			}
		}
	}
	//a.q.CorrectAnswerTotal = correct
	a.q.TotalQuestions = float64(len(a.q.Problems))
	return a
}

// Grade determines if the user passed the problems based on the number of correct answers
// compared to the total number of questions asked. The passingscore variable sets the
// pass or fail threshold.
func (a *ArithmeticBuilder) Grade() QuizBuilder {
	a.q.TestScore = (a.q.CorrectAnswerTotal / a.q.TotalQuestions) * 100
	if a.q.TestScore < float64(a.q.PassingScore) {
		return a
	}
	a.q.Pass = true
	return a
}

func (a *ArithmeticBuilder) GiveResults() QuizBuilder {
	fmt.Println("\n------------END---------------")
	fmt.Printf("Total Questions: %d\n", int(a.q.TotalQuestions))
	fmt.Printf("Correct Answers: %d\n", int(a.q.CorrectAnswerTotal))
	fmt.Printf("Passing Score: %d %%\n", a.q.PassingScore)
	fmt.Printf("Score: %.2f %%\n", math.RoundToEven(a.q.TestScore))
	fmt.Printf("Pass: %t\n", a.q.Pass)
	return a
}
