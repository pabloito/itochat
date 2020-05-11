package serverlib

import (
	"bytes"
	"github.com/pabloito/itochat/api"
	"net"
	"testing"
)

func getConn(t *testing.T) (net.Listener, net.Conn, net.Conn){
	l, err := net.Listen("tcp", ":12345")
	if err != nil{
		t.Fatal(err)
	}
	conC, err := net.Dial("tcp", ":12345")
	if err != nil{
		t.Fatal(err)
	}
	conS, err := l.Accept();
	return l, conS, conC
}
func getCm() ClientManager {
	return ClientManager{
		Clients:    make(map[*api.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *api.Client),
		Unregister: make(chan *api.Client),
	}
}

func TestReceive(t *testing.T){
	cm := getCm()
	l, conS, conC := getConn(t)
	defer l.Close()
	defer conS.Close()
	defer conC.Close()
	client := &api.Client{Socket: conS}
	want := []byte{0,0,0,0,0,0,1,1,1,1,1,1,0,1,1,1,0}
	_, err := conC.Write(want)
	if err != nil{
		t.Fatal(err)
	}
	got := cm.receive(client)

	if !bytes.Equal(got,want){
		t.Errorf("got %q, want %q",got,want)
	}
}

func TestSend(t *testing.T){
	cm := getCm()
	l, conS, conC := getConn(t)
	defer l.Close()
	defer conS.Close()
	defer conC.Close()
	client := &api.Client{Socket: conS, Data: make(chan []byte)}
	want := []byte{0,0,0,0,0,0,1,1,1,1,1,1,0,1,1,1,0}
	go cm.send(client)
	client.Data <- want

	got := make([]byte, 4096)
	leng,_ := conC.Read(got)
	got = got[:leng]
	if !bytes.Equal(got,want){
		t.Errorf("got %q, want %q",got,want)
	}

}
func TestRegister(t *testing.T){

}
func TestRegisterNoClient(t *testing.T){

}