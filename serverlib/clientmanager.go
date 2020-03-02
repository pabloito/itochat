package serverlib

//todo: add make command.
//unexport fields and add getters.

import(
    "fmt"
    "github.com/pabloito/itochat/api"
)

type ClientManager struct {
    Clients    map[*api.Client]bool
    Broadcast  chan []byte
    Register   chan *api.Client
    Unregister chan *api.Client
}

func (manager *ClientManager) Start() {
    for {
        select {
        case connection := <-manager.Register:
            manager.Clients[connection] = true
            fmt.Println("Added new connection!")
        case connection := <-manager.Unregister:
            if _, ok := manager.Clients[connection]; ok {
                close(connection.Data)
                delete(manager.Clients, connection)
                fmt.Println("A connection has terminated!")
            }
        case message := <-manager.Broadcast:
            for connection := range manager.Clients {
                select {
                case connection.Data <- message:
                default:
                    close(connection.Data)
                    delete(manager.Clients, connection)
                }
            }
        }
    }
}

func (manager *ClientManager) Receive(client *api.Client) {
    for {
        message := make([]byte, 4096)
        length, err := client.Socket.Read(message)
        if err != nil {
            manager.Unregister <- client
            client.Socket.Close()
            break
        }
        if length > 0 {
            fmt.Println("RECEIVED: " + string(message))
            manager.Broadcast <- message
        }
    }
}
func (manager *ClientManager) Send(client *api.Client) {
    defer client.Socket.Close()
    for {
        select {
        case message, ok := <-client.Data:
            if !ok {
                return
            }
            client.Socket.Write(message)
        }
    }
}
