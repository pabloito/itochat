package serverlib

//todo: add make command.
//unexport fields and add getters.

import(
    "fmt"
    "github.com/pabloito/itochat/api"
    "net"
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
            manager.addConection(connection)
        case connection := <-manager.Unregister:
            manager.removeConnection(connection)
        case message := <-manager.Broadcast:
            manager.broadcastMessage(message)
        }
    }
}

func (manager *ClientManager) addConection(connection *api.Client) {
    manager.Clients[connection] = true
    fmt.Println("Added new connection!")
}

func (manager *ClientManager) removeConnection(connection *api.Client) {
    if _, ok := manager.Clients[connection]; ok {
        connection.Socket.Close()
        close(connection.Data)
        delete(manager.Clients, connection)
        fmt.Println("A connection has terminated!")
    }
}
func (manager *ClientManager) broadcastMessage(message []byte) {
    for client := range manager.Clients {
        select {  //check if client closed
        case client.Data <- message:
        default:
            close(client.Data)
            delete(manager.Clients, client)
        }
    }
}

func (manager *ClientManager) receiveLoop(client *api.Client) { //todo: terminate connection
    exit := false
    for !exit {
        ok, message := manager.receive(client)
        if ok {
            fmt.Println("RECEIVED: " + string(message))
            manager.Broadcast <- message
        }else{
            fmt.Println("Closing Receive Loop")
            exit=true
        }
    }
}

func (manager *ClientManager) receive(client *api.Client) (bool,[]byte) {
    message := make([]byte, 4096)
    l, err := client.Socket.Read(message)
    if err != nil {
        manager.Unregister <- client
        return false,[]byte("")
    }
    return true, message[:l]
}
func (manager *ClientManager) sendLoop(client *api.Client) {//todo: terminate connection
    exit := false
    for !exit{
        ok := manager.send(client)
        if(!ok){
            fmt.Println("Closing Send Loop")
            exit = true
        }
    }
}
func (manager *ClientManager) send(client *api.Client) bool {
    select {
    case message, ok := <-client.Data:
        if ok {
            client.Socket.Write(message)
            return true
        }
    }
    return false
}

func (manager *ClientManager) RegisterLoop(listener net.Listener) {
    for {
        manager.register(listener)
    }
}

func (manager *ClientManager) register(listener net.Listener) {
    connection, err := listener.Accept()
    if err != nil {
        fmt.Println(err)
    }
    client := &api.Client{Socket: connection, Data: make(chan []byte)}
    manager.Register <- client
    go manager.receiveLoop(client)
    go manager.sendLoop(client)
}




