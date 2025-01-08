package student

import (
	"fmt"
	"math/rand"
	"math_question/teacher"
	"sync"
	"time"
)

type Student struct {
	Name string
}

// AnswerQuestion 學生嘗試回答問題
func (s *Student) AnswerQuestion(question *teacher.Question, wg *sync.WaitGroup, correctCh chan string) {
	defer wg.Done()

	thinkTime := rand.Intn(3) + 1
	time.Sleep(time.Duration(thinkTime) * time.Second)

	// 嘗試回答問題
	question.Lock()
	defer question.Unlock()

	// 如果問題已經回答正確，跳過
	if question.Answered {
		return
	}

	var studentAnswer int
	// 將該學生放入Map中表示有回答過
	question.Attempts[s.Name] = true
	// 隨機決定是否回答正確
	isCorrect := rand.Intn(2) == 0 // 50% 機率回答正確
	if isCorrect {
		studentAnswer = question.Answer
	} else {
		studentAnswer = question.Answer + rand.Intn(10) + 1 // 隨機生成錯誤答案，偏移量為10
	}

	// 輸出結果
	fmt.Printf("%s: %s = %d!\n", s.Name, question.Text, studentAnswer)
	if isCorrect {
		question.Answered = true
		question.AnsweredBy = s.Name
		correctCh <- s.Name
	} else {
		fmt.Printf("Teacher:%s, you are wrong.\n", s.Name)
	}
}
