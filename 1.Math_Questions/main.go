package main

import (
	"fmt"
	"math_question/student"
	"math_question/teacher"
	"sync"
	"time"
)

func main() {
	students := []student.Student{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
		{Name: "D"},
		{Name: "E"},
	}

	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(3 * time.Second) 

	questionsCh := make(chan *teacher.Question, 10)

	// 問題的thread
	go teacher.AskQuestions(questionsCh)

	for question := range questionsCh {
		correctCh := make(chan string, 1)
		var wg sync.WaitGroup

		// 回答thread
		for _, student := range students {
			wg.Add(1)
			go student.AnswerQuestion(question, &wg, correctCh)
		}

		go func() {
			wg.Wait()
			close(correctCh)

			question.Lock()
			defer question.Unlock()
			if question.Answered {
				//一旦正確則將回答過的人印出，除了自己答對的以外
				fmt.Printf("Teacher: %s, %s you are right!\n", question.ID, question.AnsweredBy)
				for name, answered := range question.Attempts {
					if answered && name != question.AnsweredBy {
						fmt.Printf("%s:%s %s, you win.\n", name, question.ID, question.AnsweredBy)
					}
				}
			} else {
				fmt.Printf("Teacher: Boooo~ %s answer is %d.\n", question.ID, question.Answer)
			}
		}()
	}
}
