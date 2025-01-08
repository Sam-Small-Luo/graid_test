package member

import "fmt"

// Member 結構表示 quorum 中的單一成員
type Member struct {
	ID             int
	IsLeader       bool
	IsAlive        bool
	Voted          bool
	Heartbeat      map[int]bool // 記錄其他成員的心跳回應
	WantToBeLeader bool         // 標記是否有意成為領導者
}

// NewMember 創建一個新的成員，並設置其 ID
func NewMember(id int) *Member {
	return &Member{
		ID:             id,
		IsLeader:       false,
		IsAlive:        true,
		Voted:          false,
		WantToBeLeader: false,
		Heartbeat:      make(map[int]bool),
	}
}

// DeclareIntentToBeLeader 設置成員的意圖為成為領導者
func (m *Member) DeclareIntentToBeLeader() {
	m.WantToBeLeader = true
	fmt.Printf("> Member %d: I want to be leader\n", m.ID)
}

// SendHeartbeat 發送心跳信號（只有領導者可以發送）
func (m *Member) SendHeartbeat(members []*Member) {
	if m.IsLeader {
		for _, member := range members {
			if member.IsAlive && member.ID != m.ID {
				member.ReceiveHeartbeat(m)
			}
		}
	}
}

// ReceiveHeartbeat 接收來自其他成員的心跳信號
func (m *Member) ReceiveHeartbeat(sender *Member) {
	if m.IsAlive {
		fmt.Printf("Member %d: Heartbeat received from Member %d\n", m.ID, sender.ID)
		m.Heartbeat[sender.ID] = true
	}
}

// ElectLeader 根據投票結果嘗試選舉領導者
func (m *Member) ElectLeader(members []*Member) *Member {
	if !m.WantToBeLeader {
		// 只有表明有意成為領導者的成員才會被選舉
		return nil
	}

	votes := 0
	total := 0
	for _, member := range members {
		if member.IsAlive && member.ID != m.ID {
			total++
			if member.VoteFor(m) {
				votes++
			}
		}
	}

	if votes > total/2 {
		m.IsLeader = true
		fmt.Printf("> Member %d: Accept member %d to be leader\n", m.ID, m.ID)
		return m
	}

	// 顯示投票結果並計算
	fmt.Printf("> Member %d voted to be leader: (%d > %d/%d)\n", m.ID, votes, total, 2)
	return nil
}

// VoteFor 投票選擇某個成員為領導者
func (m *Member) VoteFor(member *Member) bool {
	// 簡單的投票邏輯：如果成員還活著且尚未投票，就進行投票
	if m.IsAlive && !m.Voted {
		m.Voted = true
		return true
	}
	return false
}

// Fail 將成員標記為失敗（模擬成員無反應）
func (m *Member) Fail() {
	m.IsAlive = false
	fmt.Printf("Member %d: I'm dead now (failure)\n", m.ID)
}

// KickOut 從 quorum 中移除一個成員
func (m *Member) KickOut() {
	m.IsAlive = false
	fmt.Printf("Member %d: Kicked out of quorum\n", m.ID)
}
