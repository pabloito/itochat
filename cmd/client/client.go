package main

import (
	"bufio"
	"fmt"
	"github.com/pabloito/itochat/api"
	"github.com/pabloito/itochat/clientlib"
	"net"
	"os"
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
		_, err := client.Socket.Write([]byte(command.Msg))
		if err != nil{
			fmt.Println(err)
			return false
		}
        fmt.Printf("message '%s' Sent!\n",command.Msg)
    case lib.Exit:
        fmt.Printf("Exiting program!\n")
        return true
    case lib.Invalid:
        fmt.Printf("Invalid Command '%s'\n",command.Str)
    }
    return false
}