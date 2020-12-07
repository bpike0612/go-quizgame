package quiz

type iQuiz interface {
	PrepareQuiz(file string) ([]string, error)
	TakeQuiz(map[string]string) (float64, float64)
	Grade(correct float64, total float64, passingScore int) (bool, float64)
}
