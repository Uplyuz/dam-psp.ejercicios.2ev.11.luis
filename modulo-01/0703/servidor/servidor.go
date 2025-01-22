package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	listener, error := net.Listen("tcp", ":8080")
	if error != nil {
		log.Fatal(error)
	}

	defer listener.Close()

	if listener != nil {
		fmt.Println("SERVIDOR ESPERANDO CONEXIONES EN EL PUERTO 8080")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		con, error := listener.Accept()
		if error != nil {
			log.Println(error)
			continue
		}
		go manejarCon(con, &wg, i)
	}

	wg.Wait()

	fmt.Println("SERVIDOR DETENIDO, ALCANZADAS LAS 10 CONEXIONES")
	fmt.Println("SERVIDOR FINALIZADO")

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
			fmt.Println(err)
		}
		fmt.Println("MENSAJE ENVIADO")
	}
}
