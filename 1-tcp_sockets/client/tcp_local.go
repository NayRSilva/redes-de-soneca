package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
    input := bufio.NewScanner(os.Stdin)
	// connect to this socket
    conn, err := net.Dial("tcp", "127.0.0.1:8081")
    if err != nil {
        panic(err)
	}
    for {
        fmt.Print("Mensagem da requisição: ")
		input.Scan()
		if request := input.Text(); request == "STOP" {
			break
		} else {
			fmt.Fprintf(conn, request+"\n")
		}
        // listen for reply
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Resposta do servidor: "+message)
	}
	conn.Close()
}