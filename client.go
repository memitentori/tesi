package main

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "net/http"

    "golang.org/x/net/http2"
)

func Request() {
    client := http.Client{

        Transport: &http2.Transport{
            DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
                return net.Dial(netw, addr)
            },
        },
    }

    resp, err := client.Get("https://golang.com")
    if err != nil {
        log.Fatal(err)
        return
    }

    fmt.Printf("resp: %#v\n", resp)

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
        return
    }

    fmt.Println(string(body))
}
