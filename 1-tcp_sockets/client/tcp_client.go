package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
)

func main() {
    // cria buffer de entrada
    input := bufio.NewScanner(os.Stdin)         

    fmt.Print("Qual site/IP e porta deseja se conectar? ")
    // lê o buffer e passa o conteúdo pra host
    input.Scan()                                
    host := input.Text()                        
    // cria conexão tcp na string (host:porta) especificada
    conn, _ := net.Dial("tcp", host)
    // com a conexão estabelecida, faça a requisição
    fmt.Print("\nDigite a requisição: ")
    input.Scan()
    // envia a requisição a conn
    fmt.Fprint(conn, input.Text()+"\n\n")
    // cria buffer para ler a requisição. Não sei ao certo a diferença entre Scanner e Reader mas foi assim que funfou
    response := bufio.NewReader(conn)    
    for {
        // lê a próxima linha da string. se chegou em EOF ou algum outro erro, termina a conexão
        if message, err := response.ReadString('\n'); err != nil {
            break 
        } else {
            fmt.Print(message)
        }
    }
    conn.Close()
}