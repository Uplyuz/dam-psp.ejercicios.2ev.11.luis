package main

import (
	"fmt"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error con la udp address")
		return
	}

	con, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error iniciando el servidor udp")
		return
	}
	defer con.Close()

	if con != nil {
		fmt.Println("SERVIDOR ESPERANDO CONEXIONES EN EL PUERTO 8080")
	}

	buffer := make([]byte, 10)
	n, clientAddr, err := con.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer el cliente")
		return
	}

	fmt.Println("MENSAJE DEL CLIENTE ", clientAddr, string(buffer[:n]))

	fmt.Println("MENSAJE NO ENVIADO AL CLIENTE ", clientAddr)

	fmt.Println("SERVIDOR FINALIZADO")

}
