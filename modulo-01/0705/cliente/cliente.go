package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	con, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		log.Fatal(error)
	}
	defer con.Close()

	lector := bufio.NewReader(con)
	mensaje, er := lector.ReadString('\n')

	if er != nil {
		log.Println("Error al leeer el mensaje")
		return
	}
	fmt.Println("MENSAJE RECIBIDO XXXXX")
	fmt.Println(mensaje)
	fmt.Println("CLIENTE FINALIZADO ")

}
