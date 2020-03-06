package api

import (
    "fmt"
    "net"
)

type Client struct {
    Socket net.Conn
    Data   chan []byte
}

func (client *Client) ListenLoop() {
    for {
        message := make([]byte, 4096)
        length, err := client.Socket.Read(message)
        if err != nil {
            defer client.Socket.Close()
            panic("Error in netSocket")
        }
        if length > 0 {
            fmt.Println("RECEIVED: " + string(message))
        }
    }
}