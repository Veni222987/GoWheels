package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v4/net"
)

type Message struct {
	Name      string `json:"name"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

const MessageContent = "长度为十个字的一句话长度为十个字的一句话长度为十个字的一句话"

func testWebsocket() {
	// 开始监听本地回环网卡流量
	interfaceName := "lo"
	// 获取网络接口的统计信息
	stats, err := net.IOCounters(true)
	if err != nil {
		log.Fatal(err)
	}
	var recvBeg, sentBeg uint64
	for _, stat := range stats {
		if stat.Name == interfaceName {
			recvBeg, sentBeg = stat.BytesRecv, stat.BytesSent
			log.Printf("开始接收字节数: %d, 发送字节数: %d", recvBeg, sentBeg)
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/websocket"}
	log.Printf("连接到 %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("连接错误:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("读取错误,协程关闭:", err)
				return
			}
			log.Printf("服务器的消息: %s", message)
		}
	}()

	for i := range 10 {
		b, err := json.Marshal(Message{
			Name:      "User1",
			Msg:       MessageContent,
			Timestamp: time.Now().UnixMilli(),
		})
		if err != nil {
			log.Println(i, err)
			continue
		}
		err = c.WriteMessage(websocket.TextMessage, b)
		if err != nil {
			log.Println(i, err)
		}
	}
	c.Close()

	log.Println("结束")
	stats, err = net.IOCounters(true)
	if err != nil {
		log.Fatal(err)
	}
	for _, stat := range stats {
		if stat.Name == interfaceName {
			log.Printf("结束接收字节数: %d, 发送字节数: %d", stat.BytesRecv, stat.BytesSent)
			log.Println("websocket总IO字节数：", stat.BytesRecv+stat.BytesSent-recvBeg-sentBeg)
		}
	}
}

func testHTTP() {
	// 开始监听本地回环网卡流量
	interfaceName := "lo"
	// 获取网络接口的统计信息
	stats, err := net.IOCounters(true)
	if err != nil {
		log.Fatal(err)
	}
	var recvBeg, sentBeg uint64
	for _, stat := range stats {
		if stat.Name == interfaceName {
			recvBeg, sentBeg = stat.BytesRecv, stat.BytesSent
			log.Printf("开始接收字节数: %d, 发送字节数: %d", recvBeg, sentBeg)
		}
	}

	for i := range 10 {
		// TODO 这里补充http请求的
		url := "http://127.0.0.1:8080/http"

		// 准备POST请求的JSON数据
		b, _ := json.Marshal(Message{
			Name:      "User1",
			Msg:       MessageContent,
			Timestamp: time.Now().UnixMilli(),
		})

		jsonData, err := json.Marshal(b)
		if err != nil {
			fmt.Println("无法序列化JSON数据:", err)
			return
		}

		// 发送POST请求
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("POST请求失败:", i, err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("响应状态码:", resp.StatusCode)

		// 读取和打印响应内容
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("无法读取响应:", i, err)
			return
		}
		fmt.Println("响应内容:", string(body))
	}

	log.Println("结束")
	stats, err = net.IOCounters(true)
	if err != nil {
		log.Fatal(err)
	}
	for _, stat := range stats {
		if stat.Name == interfaceName {
			log.Printf("结束接收字节数: %d, 发送字节数: %d", stat.BytesRecv, stat.BytesSent)
			log.Println("http总IO字节数：", stat.BytesRecv+stat.BytesSent-recvBeg-sentBeg)
		}
	}
}

func main() {
	testWebsocket()
	testHTTP()
}
