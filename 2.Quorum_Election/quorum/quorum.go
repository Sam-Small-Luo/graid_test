package quorum

import (
	"fmt"
	"quorum_election/member"
)

// Quorum 結構表示整個 quorum 的成員
type Quorum struct {
	Members []*member.Member
	Leader  *member.Member
}

// NewQuorum 創建一個指定數量的成員的 quorum
func NewQuorum(size int) *Quorum {
	members := make([]*member.Member, size)
	for i := 0; i < size; i++ {
		members[i] = member.NewMember(i)
	}
	return &Quorum{Members: members}
}

// StartElection 開始選舉過程並選擇領導者
func (q *Quorum) StartElection() *member.Member {
	// 初始問候語
	for _, m := range q.Members {
		if m.IsAlive {
			fmt.Printf("> Member %d: Hi\n", m.ID)
		}
	}

	// 所有成員表明選擇成為領導者的意圖
	for _, m := range q.Members {
		if m.IsAlive {
			m.DeclareIntentToBeLeader()
		}
	}

	// 嘗試選舉領導者
	var leader *member.Member
	for _, m := range q.Members {
		if m.IsAlive {
			leader = m.ElectLeader(q.Members)
			if leader != nil {
				q.Leader = leader
				return leader
			}
		}
	}
	return nil
}

// SendHeartbeats 由領導者發送心跳信號
func (q *Quorum) SendHeartbeats() {
	if q.Leader != nil {
		q.Leader.SendHeartbeat(q.Members)
	}
}

// FailMember 使指定 ID 的成員失敗（模擬無反應）
func (q *Quorum) FailMember(id int) {
	for _, m := range q.Members {
		if m.ID == id {
			m.Fail()
			if m.IsLeader {
				q.Leader = nil // 領導者失敗後需要重新選舉
			}
			break
		}
	}
}

// KickOutMember 將指定 ID 的成員移出 quorum
func (q *Quorum) KickOutMember(id int) {
	for _, m := range q.Members {
		if m.ID == id {
			m.KickOut()
			if m.IsLeader {
				q.Leader = nil // 領導者被移除後需要重新選舉
			}
			break
		}
	}
}

// CheckQuorum 檢查 quorum 是否仍然有效
func (q *Quorum) CheckQuorum() bool {
	activeMembers := 0
	for _, m := range q.Members {
		if m.IsAlive {
			activeMembers++
		}
	}
	// 如果活著的成員少於一半，則 quorum 失敗
	return activeMembers > len(q.Members)/2
}

// SimulateFailure 模擬失敗檢測並將失敗的成員移除
func (q *Quorum) SimulateFailure() {
	for _, m := range q.Members {
		if m.IsAlive {
			// 檢查領導者是否檢測到成員的失敗
			for _, other := range q.Members {
				if other.IsAlive && other.ID != m.ID && !other.Heartbeat[m.ID] {
					fmt.Printf("> Member %d: failed heartbeat with Member %d\n", m.ID, other.ID)
				}
			}
		}
	}
}
