package teacher

import "sync"

type Question struct {
	ID         string
	Text       string
	Answer     int
	Answered   bool
	AnsweredBy string // 正確回答者
	Mutex      sync.Mutex
	Attempts   map[string]bool // 記錄嘗試回答的學生
}

// NewQuestion 建立新的問題
func NewQuestion(id, text string, answer int) *Question {
	return &Question{
		ID:       id,
		Text:     text,
		Answer:   answer,
		Attempts: make(map[string]bool),
	}
}

// Lock 加鎖問題，表示有人正在回答
func (q *Question) Lock() {
	q.Mutex.Lock()
}

// Unlock 解鎖問題
func (q *Question) Unlock() {
	q.Mutex.Unlock()
}
