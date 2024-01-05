package main

import (
	"fmt"
)

type QuizMap struct {
	data map[int]Quiz
}
type Quiz struct {
	id       int
	question string
	answer   string
}

func (q *Quiz) addQuiz(id int, question, answer string) {
	q.id = id
	q.question = question
	q.answer = answer

}

func (q *Quiz) printQuiz() {
	fmt.Println("Quiz:")

	fmt.Printf("ID: %d\nQuestion: %s\nAnswer: %s\n", q.id, q.question, q.answer)
}

func main() {
	// Creating an instance of the Quiz struct
	quiz := Quiz{}

	quiz.addQuiz(1, "Question1", "ans1")
	quiz.addQuiz(2, "Question2", "ans2")

	quiz.printQuiz()
}
