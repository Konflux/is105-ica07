package main

import (
  "net"
  "../check"
)

func main ()  {

  //Koble til TCP
  conn, err := net.Dial("tcp", "127.0.0.1:16385")
  check.Check(err)
  defer conn.Close()

  //Send melding
  conn.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))

  // Vi vil også lese en melding vi får i retur
  buf := make([]byte, 1024)
  // conn.Read() vil i dette tilfellet alltid lese over EOF og få feil, så
  // vi hopper over errorhandling her.
  n, _ := conn.Read(buf)

  println(string(buf[:n]))
}
