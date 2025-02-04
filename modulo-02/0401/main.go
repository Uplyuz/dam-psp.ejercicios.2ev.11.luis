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

	if len(args) < 2 {
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

	response, _ := reader.ReadString('\n')

	fmt.Println(response) // 220 MailHog at your service

	// Enviar comando HELO
	fmt.Fprintf(writer, "HELO localhost\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 MailHog at your service

	// Enviar MAIL FROM (Remitente)
	from := "alumno@fempa.local"
	fmt.Fprintf(writer, "MAIL FROM:<%s>\r\n", from)
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK

	// Enviar RCPT TO (Destinatario)
	to := "demo@fempa.local"
	fmt.Fprintf(writer, "RCPT TO:<%s>\r\n", to)
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK

	// Iniciar la transmisión del cuerpo del mensaje con DATA
	fmt.Fprintf(writer, "DATA\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 354 Start mail input; end with <CRLF>.

	// Enviar el contenido del correo (asunto, cuerpo)
	subject := "Prueba de envio de correo: PSP"
	body := "Cuerpo del mensaje de correo."
	fmt.Fprintf(writer, "Subject: %s\r\n\r\n%s\r\n.\r\n", subject,
		body)
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 250 OK: queued as 12345

	// Cerrar la sesión SMTP con el comando QUIT
	fmt.Fprintf(writer, "QUIT\r\n")
	writer.Flush()
	response, _ = reader.ReadString('\n')
	fmt.Println(response) // 221 Bye

	// El mensaje ha sido enviado correctamente a MailHog
	fmt.Println("Email enviado correctamente")
}
