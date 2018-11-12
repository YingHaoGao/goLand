package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
	"io/ioutil"
)

func main() {

	showAlert()
	createFile( "C:/Users/Administrator/Desktop/Reminder.txt", " " )

	for i := 0; i <= 100; i++ {
		time.Sleep(30 * time.Millisecond)

		idx := 0

		setKey( idx, i )
		setMouse()
	}
}

/* 创建文件 */
func createFile( name,content string ) {
	data :=  []byte(content)
	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}
}

/* 模拟点击键盘 */
func setKey(idx int, i int) {
	keysArr := []string{ "up", "right", "down", "left", "enter", "escape", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r" }

	if i % len(keysArr) >= len(keysArr) {
		idx = 0
	}else{
		idx = i % len(keysArr)
	}

	robotgo.KeyTap(  keysArr[idx] )
}

/* 设置鼠标定位 */
func setMouse() {
	x := ( rand.Intn(100) * 9 * 2 ) + 100
	y := ( rand.Intn(100) * 9 ) + 100

	robotgo.MoveMouse( x, y )
}

/* 显示提示 */
func showAlert() {
	robotgo.ShowAlert( "Reminder", string("congratulations! Your computer is poisoned! You can restart your computer ^_^Y") )
}
