package teacher

import (
	"fmt"
	"math/rand"
	"time"
)

var mathSymbols = []string{"+", "-", "*", "/"}

// AskQuestions 持續產生問題並傳遞到通道
func AskQuestions(questionsCh chan *Question) {
	questionID := 1
	for {
		// 每秒生成一道新問題
		a, b := rand.Intn(101), rand.Intn(101)
		c := mathSymbols[rand.Intn(len(mathSymbols))]
		questionText := fmt.Sprintf("Q%d: %d %s %d", questionID, a, c, b)

		// 計算正確答案
		var correctAnswer int
		switch c {
		case "+":
			correctAnswer = a + b
		case "-":
			correctAnswer = a - b
		case "*":
			correctAnswer = a * b
		case "/":
			if b != 0 {
				correctAnswer = a / b
			} else {
				continue // 避免除以零的情況
			}
		}

		// 新問題
		question := NewQuestion(fmt.Sprintf("Q%d", questionID), questionText, correctAnswer)

		fmt.Printf("Teacher: %s = ?\n", questionText)
		questionsCh <- question
		questionID++

		time.Sleep(1 * time.Second) // 每秒生成一道新問題
	}
}
