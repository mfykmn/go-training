package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	s := &websocket.Server{Handler: socket}
	http.HandleFunc("/socket", s)
	http.ListenAndServe(":3000", s)
}

func socket(conn *websocket.Conn) {

	// Add this request user to participants list.
	participation := participants.PushBack(conn)
	// FIXME: It's better to embed more information to participants list.
	//        For example, hmm..., yes, like "ID"

	rand.Seed(time.Now().Unix())
	id := fmt.Sprintf("#%02x%02x%02x", rand.Intn(255), rand.Intn(255), rand.Intn(255))
	logger := log.New(os.Stdout, fmt.Sprintf("[%s]\t", id), 0)

	defer func() {
		conn.Close()
		participants.Remove(participation) // clean up
		logger.Println("Exited loop")
	}()

	// Sturct for decoding message from client side
	msg := struct {
		Text string
		Type string
	}{}

	// {{{ FIXME: Tell who this request user is.
	ev := &Event{Type: "CONNECT", Text: "yourself", User: id}
	b, _ := json.Marshal(ev)
	conn.Write(b)
	// }}}

	// This loop keeps alive unless any error raises.
	for {

		if err := websocket.JSON.Receive(conn, &msg); err != nil {
			if err == io.EOF {
				logger.Println("Connection closed:", err)
			} else {
				logger.Println("Unexpected error:", err)
			}
			return // Exit from this loop
		}

		switch msg.Type {
		case "KEEPALIVE":
		// do nothing
		default:
			// event := &Event{
			// 	Type: "MESSAGE",
			// 	Text: msg.Text,
			// 	User: id,
			// }
			// b, _ := json.Marshal(event)
			// // Publish to all participants.
			// for e := participants.Front(); e != nil; e = e.Next() {
			// 	// FIXME: Type assertion validation :p
			// 	// FIXME: Error handling on Write :p
			// 	e.Value.(*websocket.Conn).Write(b)
			// }
			// FIXME: It's better to make some func to separate this common process :p
		}

		// continue: Keep waiting for message from this connection.
	}

}
