package main

import (
	"win"
	"syscall"
	"strconv"
)

func _TEXT(_str string) *uint16{
	return syscall. StringToUTF16Ptr(_str)
}

func _toString(_n int32) string{
	return strconv.Itoa(int(_n))
}

func main() {
	var hwnd win.HWND
	cxScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	cyScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.MessageBox(hwnd,_TEXT("大家好，我屏幕的宽度为：" + _toString(cxScreen) + " 高度为：" + _toString(cyScreen)),_TEXT("Golang 窗口测试"),win.MB_OK)
	win.MessageBox(hwnd,_TEXT("大家好，我屏幕的宽度为：" + _toString(cxScreen) + " 高度为：" + _toString(cyScreen)),_TEXT("Golang 窗口测试"),win.MB_OK)
	win.MessageBox(hwnd,_TEXT("大家好，我屏幕的宽度为：" + _toString(cxScreen) + " 高度为：" + _toString(cyScreen)),_TEXT("Golang 窗口测试"),win.MB_OK)
	win.MessageBox(hwnd,_TEXT("大家好，我屏幕的宽度为：" + _toString(cxScreen) + " 高度为：" + _toString(cyScreen)),_TEXT("Golang 窗口测试"),win.MB_OK)
	win.MessageBox(hwnd,_TEXT("大家好，我屏幕的宽度为：" + _toString(cxScreen) + " 高度为：" + _toString(cyScreen)),_TEXT("Golang 窗口测试"),win.MB_OK)
}