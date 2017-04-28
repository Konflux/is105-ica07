package main

import (
  "net"
  "log"
)

func main ()  {

  //Koble til TCP
  conn, err := net.Dial("tcp", "host:port")
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  //Send melding
  conn.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))
}
