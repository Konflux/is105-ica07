package main

import (
    "crypto/tls"
    "net"
    "log"
    "../check"
)



func main() {

    // Vi laster inn TLS-nøkkelparet fra fil
    cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
    check.Check(err)

    config := &tls.Config{Certificates: []tls.Certificate{cer}}

    // Åpne for tilkoblinger via TCP
    port := "16443"
    l, err := tls.Listen("tcp", ":" + port, config)
    check.Check(err)
    defer l.Close()

    println("Listening on port " + port)

    for {
        // Accept() vil blokkere loopen til en innkommende tilkobling skjer
        conn, err := l.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    // Vi lager et buffer for å lese bytes inn til
    buffer := make([]byte, 1024)

    // Av en eller annen grunn vil Read i dette tilfellet ikke gi error
    // om bufferet er for stort
    rlen, err := conn.Read(buffer)
    check.Check(err)

    msg := string(buffer[0:rlen])
    println(msg)

    // Vi sender en melding tilbake for å la klienten vite at vi har mottatt
    // meldingen
    n, err := conn.Write([]byte("Affirmative: " + msg))
    if err != nil {
        log.Println(n, err)
        return
    }
}
