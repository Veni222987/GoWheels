package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Name      string `json:"name"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// 协议升级
	// ? 协议升级为什么是在服务端做的?
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		resMsg := "收到:" + string(p)
		log.Println(string(resMsg))
		if err := conn.WriteMessage(messageType, []byte(resMsg)); err != nil {
			log.Println(err)
			return
		}
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只接受POST请求", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "无法读取请求体", http.StatusBadRequest)
		return
	}

	msg := &Message{}
	json.Unmarshal(body, msg)

	resMsg := "收到:" + msg.Msg
	log.Println(string(resMsg))

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resMsg)
}

func main() {
	http.HandleFunc("/websocket", websocketHandler)
	http.HandleFunc("/http", httpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
