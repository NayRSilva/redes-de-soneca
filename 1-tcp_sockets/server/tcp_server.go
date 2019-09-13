package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

func main() {

    fmt.Println("Launching server...")

    // listen on all interfaces
    ln, _ := net.Listen("tcp", ":8081")

    for {
        conn, _ := ln.Accept()
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    fmt.Println("Conexão aberta com "+ conn.RemoteAddr().String())
    for {
        // will listen for message to process ending in newline (\n)
        message, _ := bufio.NewReader(conn).ReadString('\n')
        // output message received
        fmt.Print("Mensagem recebida:", message)
        if strings.TrimSpace(message) == "STOP" {
            break
        }
        // sample process for string received
        newmessage := strings.ToUpper(message)
        // send new string back to client
        conn.Write([]byte(newmessage + "\n"))
    }
    fmt.Println("Conexão fechada com "+ conn.RemoteAddr().String())
}