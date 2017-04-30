package main

import (
    "crypto/tls"
    "../check"
)


func main() {
    conf := &tls.Config{
         // TLS-sertifikatet vi bruker er selvsignert
         // og vil derfor ikke kunne bli verifisert
         InsecureSkipVerify: true,
    }

    // Vi åpner en sikker TCP-tilkobling over localhost via TLS
    conn, err := tls.Dial("tcp", "127.0.0.1:16443", conf)
    check.Check(err)
    defer conn.Close()

    // Så sender vi meldingen vår til server
    n, err := conn.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))
    check.Check(err)

    // Vi vil også lese meldingen vi får i retur
    buf := make([]byte, 1024)
    // conn.Read() vil i dette tilfellet alltid lese over EOF, så
    // vi hopper over errorhandling her.
    n, _ = conn.Read(buf)

    println(string(buf[:n]))
}
