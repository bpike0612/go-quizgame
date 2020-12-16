package main

import "testing"

func TestQuizBuilder(t *testing.T) {
	director := Director{}

	a := &ArithmeticBuilder{}
	director.SetBuilder(a)
	director.Construct()

	quiz := a.GetQuiz()

	if quiz.FileName != "problems.csv" {
		t.Errorf("Quiz file name should be problems.csv and instead is set to %s\n", quiz.FileName)
	}

	if quiz.PassingScore != 80 {
		t.Errorf("Passingscore should be set to 80 and instead is set to %d\n", quiz.PassingScore)
	}

	if quiz.Pass != false {
		t.Errorf("Pass should be set to false by default and instead is set to %t\n", quiz.Pass)
	}

}
func TestQuizProduct(t *testing.T) {
	director := Director{}

	a := &ArithmeticBuilder{}
	director.SetBuilder(a)
	director.Construct()

	a.SetPass(true)

	if a.GetFileName().FileName != "problems.csv" {
		t.Errorf("Filename should be problems.csv and instead is %s\n", a.GetFileName().FileName)
	}
}
