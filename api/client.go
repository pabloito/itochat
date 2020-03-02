package api

import (
    "fmt"
    "net"
)

type Client struct {
    Socket net.Conn
    Data   chan []byte
}

func (client *Client) Receive() {
    for {
        message := make([]byte, 4096)
        length, err := client.Socket.Read(message)
        if err != nil {
            client.Socket.Close()
            break
        }
        if length > 0 {
            fmt.Println("RECEIVED: " + string(message))
        }
    }
}