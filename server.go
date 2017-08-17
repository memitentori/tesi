package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net"
    "bufio"
    "strings"
    "strconv"
)


  var image []byte
  // preparing image
  func init() {
      var err error
      image, err = ioutil.ReadFile("./image.png")
      if err != nil {
          panic(err)
      }
  }
  // Send HTML and push image
  func handlerHtml(w http.ResponseWriter, r *http.Request) {
      pusher, ok := w.(http.Pusher)
      if ok {
          fmt.Println("Push /image")
          pusher.Push("/image", nil)
      }
      w.Header().Add("Content-Type", "text/html")
      fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
  }
  // Send image as usual HTTP request
  func handlerImage(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "image/png")
      w.Write(image)
  }

func Test1(){
  http.HandleFunc("/", handlerHtml)
  http.HandleFunc("/image", handlerImage)
  fmt.Println("start http listening :18443")
  err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
  fmt.Println(err)
}

func main() {

  fmt.Println("Avvio del server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // replace \r\n from the string
    message = strings.Replace(message, "\r\n", "", -1)
    // conver the string into an int
    if n_test, err := strconv.Atoi(message); err == nil {
      // output message received
      fmt.Printf("Esecuzione del test n. %d\n", n_test)
      // select the test
      switch n_test {
        case 1:
          Test1()
        default:
          fmt.Printf("Il test selezionato non è disponibile")
        }
    }
    // sample process for string received
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }

}
