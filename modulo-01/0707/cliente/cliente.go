package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error al resolver la dirrecion del servidor", err)
		return
	}

	con, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error al conectar con el servidor", err)
		return
	}

	defer con.Close()

	mensaje := "Hola servidor"
	_, err = con.Write([]byte(mensaje))
	if err != nil {
		fmt.Println("Error al enviar mensaje:", err)
		return
	}

	lector := bufio.NewReader(con)
	mensaje, err = lector.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el mensaje del servidor", err)
		return
	}

	fmt.Println("MENSAJE RECIBIDO XXXXX")
	fmt.Println("MENSAJE RECIBIDO DESDE EL SERVIDO")
	fmt.Println(mensaje)
	fmt.Println("CLIENTE FINALIZADO ")

}
