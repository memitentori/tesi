package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
)


func main() {


  // connect to this socket
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  for {
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Indicare il numero di test che si vuole effettuare: (1)")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("E' stato selezionato il test n. "+message)

  }
}
