package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
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

	mensaje := "123456789-123456789"
	_, err = con.Write([]byte(mensaje))
	if err != nil {
		fmt.Println("Error al enviar mensaje:", err)
		return
	}

	fmt.Println("MENSAJE ENVIADO AL SERVIDOR: " + mensaje)

	go contarSegundos()

	con.SetReadDeadline(time.Now().Add(5 * time.Second))

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

func contarSegundos() {
	segundos := 0
	for {
		time.Sleep(1 * time.Second)
		segundos++
		fmt.Println("Segundos transcurridos: ", segundos)
	}
}
