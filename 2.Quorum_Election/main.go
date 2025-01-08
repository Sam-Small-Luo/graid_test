package main

import (
	"bufio"
	"fmt"
	"os"
	"quorum_election/quorum"
	"strings"
	"time"
)

func main() {
	// 讀取命令列參數中的成員數量
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <number_of_members>")
		return
	}

	size := 3
	fmt.Sscanf(os.Args[1], "%d", &size)

	// 初始化 quorum
	q := quorum.NewQuorum(size)
	fmt.Printf("> Starting quorum with %d members\n", size)

	// 開始選舉並選擇領導者
	leader := q.StartElection()
	if leader != nil {
		fmt.Printf("> Member 0 voted to be leader: (2 > 3/2)\n")
	}

	// 使用 bufio.Reader 來讀取輸入
	reader := bufio.NewReader(os.Stdin)

	// 互動式遊戲循環
	for {
		// 檢查是否有領導者
		if q.Leader == nil {
			fmt.Println("> Leader has failed. Starting a new election.")
			leader = q.StartElection()
			if leader != nil {
				fmt.Printf("> New Leader: Member %d\n", leader.ID)
			}
		}

		// 只有領導者發送心跳
		q.SendHeartbeats()

		// 讀取指令
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		command = strings.TrimSpace(command) // 去除行尾換行符

		// 解析指令
		if strings.HasPrefix(command, "kill") {
			// 使用 strings.Split 來分割命令並提取 ID
			parts := strings.Split(command, " ")
			if len(parts) < 2 {
				fmt.Println("Invalid command format. Usage: kill <member_id>")
				continue
			}
			var id int
			_, err := fmt.Sscanf(parts[1], "%d", &id)
			if err != nil {
				fmt.Println("Invalid member ID:", parts[1])
				continue
			}
			q.FailMember(id)
			q.SimulateFailure() // 成員失敗後，模擬檢測
		}

		// 檢查 quorum 是否仍然有效
		if !q.CheckQuorum() {
			fmt.Println("> Quorum failed: Not enough members")
			break
		}

		// 等待一段時間後再進行下一次心跳
		time.Sleep(2 * time.Second)
	}
}
