package main

import (
	"fmt"
	"net/rpc"
)

type Args struct{}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer client.Close()

	var reply string
	args := Args{}
	err = client.Call("HelloWorld.SayHello", args, &reply)
	if err != nil {
		fmt.Println("Error al llamar a la función RPC:", err)
		return
	}

	fmt.Println(reply)
}