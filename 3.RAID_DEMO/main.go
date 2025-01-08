package main

import (
	"fmt"
)

// 定義常數
const (
	stripeSize = 16 // 每個stripe的大小，單位是byte
	numDisks   = 4  // RAID中的磁碟數量
	dataLength = 40 // 要寫入的總數據長度
)

type RAID struct {
	disks [][]byte // 每個磁碟的數據
}

// NewRAID 函數，根據不同 RAID 等級初始化 RAID 系統
func NewRAID(level string) *RAID {
	ra := &RAID{
		disks: make([][]byte, numDisks),
	}

	// 初始化每個磁碟為 stripeSize 大小的空字節切片
	for i := 0; i < numDisks; i++ {
		ra.disks[i] = make([]byte, stripeSize)
	}

	switch level {
	case "RAID0":
		fmt.Println("This RAID is RAID0")
		return ra
	case "RAID1":
		fmt.Println("This RAID is RAID1")
		return ra
	case "RAID5":
		fmt.Println("This RAID is RAID5")
		return ra
	case "RAID6":
		fmt.Println("This RAID is RAID6")
		return ra
	case "RAID10":
		fmt.Println("This RAID is RAID0+RAID1")
		return ra
	default:
		return ra
	}
}

// WriteData 函數：將數據寫入 RAID，分配到每個磁碟的條帶中
func (ra *RAID) WriteData(data []byte) {
	// 將數據按照條帶分配到磁碟中
	for i := 0; i < len(data); i++ {
		diskIndex := i % numDisks
		ra.disks[diskIndex][i%stripeSize] = data[i]
	}
}

func (ra *RAID) ClearDisk(diskNum int) {
	// 將指定磁碟清空
	for i := 0; i < stripeSize; i++ {
		ra.disks[diskNum][i] = 0
	}
}

// ReadData 函數：從 RAID 中讀取數據，並轉換為字串
func (ra *RAID) ReadData(length int) string {
	var result []byte
	// 從 RAID 中讀取數據
	for i := 0; i < length; i++ {
		diskIndex := i % numDisks
		result = append(result, ra.disks[diskIndex][i%stripeSize])
	}
	return string(result)
}

func main() {
	// 選擇 RAID 等級
	raidLevel := "RAID5"
	raid := NewRAID(raidLevel)

	str := "Test Data Test Data"
	data := []byte(str)

	raid.WriteData(data)

	// 將磁碟2清零
	raid.ClearDisk(2)

	// 從 RAID 中讀取數據並打印出來
	readData := raid.ReadData(len(data))
	fmt.Printf("讀取的數據：%s\n", readData)
}
