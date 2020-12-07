package quiz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuiz_PrepareQuiz(t *testing.T) {
	q := New()
	m, _ := q.PrepareQuiz("../test_quiz.csv")

	var key string
	var value string
	for k, v := range m {
		key = k
		value = v
	}

	a := assert.New(t)
	a.NotNil(m)
	a.Equal("5+5", key)
	a.Equal("10", value)

}

func TestQuiz_PrepareQuiz2(t *testing.T) {
	q := New()
	m, err := q.PrepareQuiz("")
	a := assert.New(t)
	a.NotNil(m)
	a.Error(err, "Couldn't open the csv file open : no such file or directory")
}

func TestQuiz_TakeQuiz(t *testing.T) {
	q := New()
	m := make(map[string]string)
	m["5+5"] = "10"
	correct, total := q.TakeQuiz(m)
	a := assert.New(t)
	a.Equal(float64(0), correct)
	a.Equal(float64(1), total)
}

func TestQuiz_Grade(t *testing.T) {
	q := New()
	pass, score := q.Grade(float64(1), float64(1), 80)
	a := assert.New(t)
	a.Equal(true, pass)
	a.Equal(float64(100), score)
}

func TestQuiz_Grade2(t *testing.T) {
	q := New()
	pass, score := q.Grade(float64(1), float64(2), 80)
	a := assert.New(t)
	a.Equal(false, pass)
	a.Equal(float64(50), score)
}
