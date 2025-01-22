package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
)

func main() {
	mensajes := [2]string{"bye", "hi"}

	mensajeEnviado := mensajes[rand.Intn(len(mensajes))]

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

	_, err := con.Write([]byte(mensajeEnviado + "\n"))
	if err != nil {
		fmt.Println("Error al enviar mensaje")
		return
	}
}
