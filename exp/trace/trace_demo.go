package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("无法创建trace文件:", err)
		return
	}
	defer f.Close()

	// 启动trace
	err = trace.Start(f)
	if err != nil {
		fmt.Println("无法启动trace:", err)
		return
	}
	defer trace.Stop()

	time.Sleep(1 * time.Second)

	// 示例：打印一些信息
	fmt.Println("Hello, World!")
}
