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
        fmt.Println("RECEIVED:" + string(client.Listen()))
    }
}

func (client *Client) Listen() []byte{
    message := make([]byte, 4096)
    _, err := client.Socket.Read(message)
    if err != nil {
        defer client.Socket.Close()
        panic("Error in netSocket")
    }
    return message
}