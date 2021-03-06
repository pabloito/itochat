package main

import (
	"fmt"
	"net"
	"github.com/pabloito/itochat/api"
	"github.com/pabloito/itochat/serverlib"
)
func main(){
	fmt.Printf("Server\n")
	startServerMode()
}

func startServerMode() {
    fmt.Println("Starting server...")
    listener, error := net.Listen("tcp", ":12345")
    if error != nil {
        fmt.Println(error)
    }
    manager := serverlib.ClientManager{
        Clients:    make(map[*api.Client]bool),
        Broadcast:  make(chan []byte),
        Register:   make(chan *api.Client),
        Unregister: make(chan *api.Client),
    }
    go manager.Start()
    go manager.RegisterLoop(listener)
	select {}
}