package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type quiz struct {
	csv.Reader
}

func New() quiz {
	q := quiz{}
	return q
}

// open and parse quiz file
func (q quiz) PrepareQuiz(file string) (map[string]string, error) {
	//var records []string
	var m map[string]string
	m = make(map[string]string)

	// Open the file
	csvfile, err := os.Open(file)
	if err != nil {
		return m, err
	}

	// Parse the file
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
	}
	return m, err
}

// TakeQuiz() takes a map of the parsed csv file and interacts with the user,
// displaying each question to the terminal allowing the user to respond.
func (q quiz) TakeQuiz(records map[string]string) (float64, float64) {
	var correct float64
	for k, v := range records {
		fmt.Printf("%s=", k)
		var answer string
		fmt.Scanln(&answer)
		if v == answer {
			correct++
		}
	}
	var total float64
	total = float64(len(records))
	return correct, total
}

// Grade() determines if the user passed the quiz based on the number of correct answers
// compared to the total number of questions asked. The passingscore variable sets the
// pass or fail threshold.
func (q quiz) Grade(correct float64, total float64, passingScore int) (bool, float64) {
	var score float64
	score = (correct / total) * 100
	if score < float64(passingScore) {
		return false, score
	}
	return true, score
}
