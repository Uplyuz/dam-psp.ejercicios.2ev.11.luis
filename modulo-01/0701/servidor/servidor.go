package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, error := net.Listen("tcp", ":8090")
	if error != nil {
		log.Fatal(error)
	}
	defer listener.Close()

	if listener != nil {
		fmt.Println("SERVIDOR ESPERANDO CONEXIONES EN EL PUERTO 8080")
	}

	con, error := listener.Accept()
	if error != nil {
		log.Print(error)
		return
	}

	defer con.Close()

	mensaje := "SERVIDOR-01\n"
	if con != nil {
		_, err := con.Write([]byte(mensaje))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("MENSAJE ENVIADO")
	}

	fmt.Println("SERVIDOR FINALIZADO")

}
