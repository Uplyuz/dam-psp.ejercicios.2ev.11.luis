package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Error, debes introducir la dirección del servidor y el puerto de acceso.")
		return
	}

	server := args[1] + ":" + args[2]

	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	peticion := "GET /images/hog.png HTTP/1.1\r\n" +
		"Host: " + server + "\r\n" +
		"Connection: close\r\n\r\n"

	_, err = conn.Write([]byte(peticion))

	scan := bufio.NewReader(conn)

	binario := false

	for {
		texto, err := scan.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error al leer la respuesta:", err)
			os.Exit(1)
		}
		if strings.TrimSpace(texto) == "" {
			binario = true

		}

		if binario {
			fmt.Println("BINARIO>>" + texto)
		} else {
			fmt.Println("CABECERA>>" + texto)
		}
	}

}
