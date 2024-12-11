package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Args struct{}

type HelloWorld struct{}

func (h *HelloWorld) SayHello(args *Args, reply *string) error {
	*reply = "Hola Mundo desde RPC!"
	return nil
}

func main() {
	helloWorld := new(HelloWorld)
	rpc.Register(helloWorld)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor RPC en ejecución...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error de conexión:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}