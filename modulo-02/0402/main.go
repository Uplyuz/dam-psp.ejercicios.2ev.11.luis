package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
		log.Fatal("Error al conectar:", err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	// Leer respuesta del servidor
	response, _ := reader.ReadString('\n')
	fmt.Println(response) // 220 MailHog at your service

	// Enviar HELO
	fmt.Fprintf(writer, "HELO localhost\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK

	// Enviar MAIL FROM
	from := "alumno@fempa.local"
	fmt.Fprintf(writer, "MAIL FROM:<%s>\r\n", from)
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK

	// Enviar RCPT TO
	to := "demo@fempa.local"
	fmt.Fprintf(writer, "RCPT TO:<%s>\r\n", to)
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK

	fmt.Fprintf(writer, "DATA\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response)
	subject := "Prueba de envío de correo: PSP"
	fmt.Fprintf(writer, "Subject: %s\r\n", subject)
	fmt.Fprintf(writer, "Content-Type: text/html; charset=\"UTF-8\"\r\n")
	fmt.Fprintf(writer, "\r\n")

	htmlBody := `
		<html>
			<body>
				<h1>Hola, esto es un correo de prueba</h1>
				<p>Este es un mensaje enviado desde  <b>MailHog</b>.</p>
			</body>
		</html>`
	fmt.Fprintf(writer, "%s\r\n", htmlBody)

	fmt.Fprintf(writer, ".\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK

	// Enviar QUIT
	fmt.Fprintf(writer, "QUIT\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 221 Bye

	fmt.Println("Email enviado correctamente en HTML")
}
