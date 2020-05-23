// Server-side part of the Go websocket sample.
//
// Eli Bendersky [http://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
        "strings"
	"golang.org/x/net/trace"
	"golang.org/x/net/websocket"
)

var (
        wsmap = make(map[string][]*websocket.Conn)
	port = flag.Int("port", 4050, "The server port")
)

func handleWebsocketMessage(ws *websocket.Conn, e string) error {
	// Log the request with net.Trace
	tr := trace.New("websocket.Receive", "receive")
	defer tr.Finish()
	tr.LazyPrintf("Got event %v\n", e)
	return nil
}


func websocketConnection(ws *websocket.Conn) {
	log.Printf("Client connected from %s", ws.RemoteAddr())
	for {
		var event string
		err := websocket.Message.Receive(ws, &event)
                tokens := strings.Split(event, " ")
                var topic string
                switch tokens[0] {
                    case "subscribe":
                        topic = tokens[1]
                        _, prs := wsmap[topic]
                        switch prs {
                            case false:
                                wsmap[topic] = []*websocket.Conn{ws}
                            case true:
                                var present bool = false
                                for _, v := range wsmap[topic] {
                                    if v == ws {
                                        present = true
                                    }
                                }
                                if present == false {
                                    wsmap[topic] = append(wsmap[topic], ws)
                                }
                        }
                    case "publish":
                        topic := tokens[1]
                        msg := strings.Join(tokens[2:], " ")
                        conns := wsmap[topic]
                        for _, val := range conns {
                            websocket.Message.Send(val, msg)
                        }
                    case "list":
                        list := []string{}
                        for k, v := range wsmap {
                            for _, wconn := range v {
                                if wconn == ws {
                                    list = append(list, k)
                                }
                            }
                        }
                        websocket.Message.Send(ws, strings.Join(list, " "))
                }
		if err != nil {
			log.Printf("Receive failed: %s; closing connection...", err.Error())
			if err = ws.Close(); err != nil {
				log.Println("Error closing connection:", err.Error())
			}
			break
		} else {
			if err := handleWebsocketMessage(ws, event); err != nil {
				log.Println(err.Error())
				break
			}
		}
	}
}

func main() {
	flag.Parse()
	http.Handle("/ws", websocket.Handler(websocketConnection))
	log.Printf("Server listening on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
