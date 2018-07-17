package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// learn https://github.com/gorilla/websocket/blob/master/examples/chat/client.go
// learn http://blog.engineer.adways.net/entry/advent_calendar/01
// learn https://gowebexamples.com/websockets/

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, // see https://gowalker.org/github.com/gorilla/websocket#hdr-Origin_Considerations
}

type client struct {
	socket *websocket.Conn   // WebSocketのコネクション
	send   chan []byte       // メッセージをブラウザに送信するchannel
	room   *room             // クライアントが属するチャットルーム
}

func (c *client) read() {
	for {
		log.Println("called ReadMessage")
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		log.Println("called WriteMessage")
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

// チャットルーム
type room struct {
	forward chan []byte      // 他のすべてのクライアントに送信するメッセージを持つ
	join    chan *client     // 入室するクライアント
	leave   chan *client     // 退室するクライアント
	clients map[*client]bool // 入室中のすべてのクライアント
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//入室
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// すべてのクライアント(ユーザー)にメッセージを送信する
			for client := range r.clients {
				select {
				case client.send <- msg:
					// 送信
				default:
					// 失敗
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

// チャットルームの初期化処理
func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize), // バッファサイズを設定
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }() // この関数が終了する時に呼ばれる
	go client.write()                    // 別のスレッドで書き出す
	client.read()                        // メインスレッドで読み込み
}

func main() {
	r := NewRoom()
	http.Handle("/room", r)

	go r.run()

	log.Println("start server")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}