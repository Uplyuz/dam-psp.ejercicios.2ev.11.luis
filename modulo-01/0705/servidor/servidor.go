package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	i := 0
	var wg sync.WaitGroup
	canal := make(chan bool)

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
		if i%2 == 0 {
			wg.Add(1)
			go manejarCon(con, &wg, i)
		} else {
			go manejarConCanal(con, canal, i)
			mensaje := <-canal
			fmt.Println(mensaje)
		}
		i++
	}
}

func manejarCon(con net.Conn, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	defer con.Close()

	fmt.Println("Conexión local ", con.LocalAddr())
	fmt.Println("Conexión remota ", con.RemoteAddr())

	if con != nil {
		mensaje := fmt.Sprintf("CLIENTE %d CONECTADO \n", i)
		_, err := con.Write([]byte(mensaje))
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
		} else {
			fmt.Println("MENSAJE ENVIADO")
		}
	}
}

func manejarConCanal(con net.Conn, canal chan bool, i int) {
	defer con.Close()

	fmt.Println("Conexión local ", con.LocalAddr())
	fmt.Println("Conexión remota ", con.RemoteAddr())

	if con != nil {
		mensaje := fmt.Sprintf("CLIENTE %d CONECTADO \n", i)
		_, err := con.Write([]byte(mensaje))
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
			canal <- false
			return
		}

		fmt.Println("MENSAJE ENVIADO")
		canal <- true
	}
}
