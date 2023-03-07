package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 獲取當前時間
	t := time.Now()
	fmt.Printf("%T\n", t)
	// time.Time
	fmt.Println(t)
	// 未經格式化的日期時間
	// 2023-03-06 23:48:17.347652 +0800 CST m=+0.000115760
	fmt.Println("")

	// 獲取指定時間
	t1 := time.Date(2008, 7, 15, 16, 30, 28, 0, time.Local)
	fmt.Println(t1)
	// 2008-07-15 16:30:28 +0800 CST

	// time -----> string之間轉換
	// 格式化時間，會自動轉換成目前文字格式，並呈現日期時間
	f1 := t.Format("2006年1月2日 下午03:04:05")
	fmt.Println(f1)
	// 2023年3月6日 下午11:48:17

	// 格式化日期時間2
	f2 := t.Format("2006/01/02 3:4PM")
	fmt.Println(f2)
	// 2023/03/06 11:50PM

	// string ----> time 字串轉換成時間類型
	s1 := "2023年3月6日"
	// parse裡面的字串格式需和s1的字串格式一致
	// time.Parse("模板", string) ---> time, err
	t2, err := time.Parse("2006年1月2日", s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)
	// 2023-03-06 00:00:00 +0000 UTC

	year, month, day := t.Date()
	fmt.Println(year, month, day)
	// 2023 March 6

	hour, min, sec := t.Clock()
	fmt.Println(hour, min, sec)
	// 0 2 22

	// 各別取得年月日
	y := t.Year()
	m := t.Month()
	d := t.Day()
	hour1 := t.Hour()
	min1 := t.Minute()
	sec1 := t.Second()

	fmt.Printf("%d年%d月%d日 %d時%d分%d秒\n", y, m, d, hour1, min1, sec1)
	// 2023年3月7日 0時7分38秒

	// 顯示今天的星期
	fmt.Println(t.Weekday())
	// Tuesday

	fmt.Println(t.YearDay())
	// 因為2023/3/6所以今年已經過了66天

	// 時間戳: 指定的日期，距離1970年1月1日0點0時0分0秒的時間差值: 秒，納秒
	t3 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStamp := t3.Unix()

	fmt.Println(timeStamp)
	// 顯示指定的時間戳 3600

	timeStamp1 := t.Unix()
	fmt.Println(timeStamp1)
	// 目前時間戳 1678119720

	timeStamp2 := t.UnixNano()
	fmt.Println(timeStamp2)
	// 目前時間戳(納秒) 1678119720786195000
	fmt.Println(len("1678119720786195000"))
	// 總長度: 19

	// 時間間隔

	t4 := t.Add(time.Minute)
	fmt.Println(t)
	// 2023-03-07 00:22:00.786195 +0800 CST m=+0.000094040
	fmt.Println(t4)
	// t4 會比 t1多增加一分鐘
	// 2023-03-07 00:23:00.786195 +0800 CST m=+60.000094040

	t5 := t.Add(time.Hour)
	fmt.Println(t5)
	// 這樣會多一天(24小時) Add(24 * time.Hour)
	// 2023-03-08 00:22:00.786195 +0800 CST m=+86400.000094040

	t6 := t.AddDate(1, 0, 0)
	fmt.Println(t6)
	// 使用AddDate()的方法，可以增加年數或月數或日數
	// 2024-03-07 00:22:00.786195 +0800 CST

	// 時間差距
	d1 := t4.Sub(t)
	fmt.Println(d1)
	// 1m0s
	d2 := t5.Sub(t)
	fmt.Println(d2)
	// 1h0m0s
	d3 := t6.Sub(t)
	fmt.Println(d3)
	// 8784h0m0s

	// 睡眠
	// 讓當前的程序進入睡眠狀態
	fmt.Println("進入睡眠狀態")
	time.Sleep(3 * time.Second)
	// 3秒後進入睡眠狀態
	fmt.Println("睡眠結束")

	// 睡眠[1-10]的隨機秒數
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1 // type == int
	fmt.Printf("睡眠時間 %d 秒\n", randNum)
	// 因為randNum是int，sleep()裡面需要放的是time.Duration()，所以不能只是單純放 randNUm
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("睡醒了")
}
