package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	i := 0
	canal := make(chan bool, 1)
	var err error

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("SERVIDOR ESPERANDO CONEXIONES EN EL PUERTO 8080")

	for {
		con, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		resultado := manejarConCanal(con, canal, i)
		valor := <-canal
		fmt.Println(valor)
		if resultado == "BYE" {
			break
		}
		i++
	}
	fmt.Println("SERVIDOR CERRANDOSE")

}

func manejarConCanal(con net.Conn, canal chan bool, i int) string {
	defer con.Close()
	var resultado string

	fmt.Println("Conexión local ", con.LocalAddr())
	fmt.Println("Conexión remota ", con.RemoteAddr())

	mensaje := fmt.Sprintf("CLIENTE %d CONECTADO \n", i)
	_, err := con.Write([]byte(mensaje))

	if err != nil {
		fmt.Println("Error al enviar mensaje", err)
		canal <- false
		return ""

	}

	lector := bufio.NewReader(con)
	resultado, err = lector.ReadString('\n')
	if err != nil {
		fmt.Println("Error al recibir mensaje:", err)
		canal <- false
		return ""
	}

	resultado = strings.TrimSpace(resultado)
	resultado = strings.ToUpper(resultado)
	fmt.Println("MENSAJE RECIBIDO " + resultado)

	if resultado == "BYE" {
		canal <- false
	} else {
		canal <- true
	}

	return resultado
}
