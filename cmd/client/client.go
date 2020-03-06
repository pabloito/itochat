package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
    "github.com/pabloito/itochat/clientlib"
	"github.com/pabloito/itochat/api"
)
func main(){
	fmt.Printf("Client\n")
	startClientMode("tcp", "localhost:12345")
}

func startClientMode(network, address string) {
    connection, error := net.Dial(network,address)
    if error != nil {
        fmt.Printf("Error opening Socket, program will exit\n'%s'\n",error)
        os.Exit(-1)
    }
    client := &api.Client{Socket: connection}
    go client.ListenLoop()
    writeLoop(client)
}

func writeLoop(client *api.Client){
    exitDue := false
    reader := bufio.NewReader(os.Stdin)

    for !exitDue {
        str, _ := reader.ReadString('\n')
        str = str[:len(str)-1]
        command := lib.CheckCommand(str)
        exitDue = executeCommand(command,client)
    }
}

func executeCommand(command *lib.Command, client *api.Client) bool{
    switch command.T {
    case lib.Send:
        client.Socket.Write([]byte(strings.TrimRight(command.Msg, "\n")))
        fmt.Printf("message '%s' Sent!\n",command.Msg)
    case lib.Exit:
        fmt.Printf("Exiting program!\n")
        return true
    case lib.Invalid:
        fmt.Printf("Invalid Command '%s'\n",command.Str)
    }
    return false
}