package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
	"github.com/pabloito/itochat/api"
)
func main(){
	fmt.Printf("Client\n")
	startClientMode()
}

func startClientMode() {
    fmt.Println("Starting client...")
    connection, error := net.Dial("tcp", "localhost:12345")
    if error != nil {
        fmt.Println(error)
    }
    client := &api.Client{Socket: connection}
    go client.Receive()
    for {
        reader := bufio.NewReader(os.Stdin)
        message, _ := reader.ReadString('\n')
        connection.Write([]byte(strings.TrimRight(message, "\n")))
    }
}
