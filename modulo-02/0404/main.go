package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	var imagen bytes.Buffer

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
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewReader(conn)

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
			fmt.Println("Detectado fin de cabeceras HTTP.")
			break
		}

		fmt.Println("CABECERA >>" + texto)
	}

	_, err = scan.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error al descartar la primera línea:", err)
		os.Exit(1)
	}

	_, err = io.Copy(&imagen, scan)
	if err != nil {
		fmt.Println("Error al leer contenido binario:", err)
		os.Exit(1)
	}

	file, err := os.Create("salida.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = imagen.WriteTo(file)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
	}
	fmt.Println("Imagen guardada como salida.png")
}
