package main

import (
    "flag"
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
    "golang.org/x/net/websocket"
)

var (
    // Port number where the server listens
    serverport = flag.Int("serverport", 4050, "The server port")
)

var base_addr = "ws://localhost"
var origin = "http://localhost"

func read_transmit() {
    addr := fmt.Sprintf("%s:%d/ws", base_addr, *serverport)
    conn, err := websocket.Dial(addr, "", origin)
    if err != nil {
	log.Fatal("websocket.Dial error", err)
    }
    go func() {
        var reply string
        for {
	    err = websocket.Message.Receive(conn, &reply)
	    if err != nil {
		log.Fatal("websocket.JSON.Receive error", err)
	    }

            fmt.Println("received over websocket:", reply)
        }
    }()
        reader := bufio.NewReader(os.Stdin)
    for {
        line, err := reader.ReadString('\n')
        read_line := strings.TrimSuffix(line, "\n")
        mye := read_line
	err = websocket.Message.Send(conn, mye)
	if err != nil {
		log.Fatal("websocket.JSON.Send error", err)
	}
    }
}

func main() {
    flag.Parse()
    read_transmit()
}
