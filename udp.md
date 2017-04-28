# UDP pakker

## Lokal UDP pakke
Dette er et eksempel på en UDP-pakke sendt innen lokal node.
```
# TODO
```

---

## UDP pakke over nettverk
Dette er et eksempel på en UDP-pakke sendt over nettverket.

```
0000   14 2d 27 6e 56 d1 44 1c a8 4e 01 a3 08 00 45 00
0010   00 38 56 38 00 00 80 11 82 77 0a e4 22 96 0a e4
0020   29 a8 cb bd 40 00 00 24 f9 bc 4d c3 b8 74 65 20
0030   46 72 20 35 2e 35 20 31 34 3a 34 35 20 46 6c c3
0040   a5 6b 6c 79 70 61
```

### Ethernet II Header
```
[14 2d 27 6e 56 d1]  Ethernet II Destination
[44 1c a8 4e 01 a3]  Ethernet II Source
[08 00]              Ethernet II Type (0x0800 == IPv4)
```

### IPv4 Header
```
[45]                 0100 (IPv4 versjon 4)
                     0101 (IPv4 header length (20 bytes: 5?)
[00]                 0000 00 - Differentiated Services Codepoint: Default (0)
                     00 - Explicit Congestion Notification: Not ECN-Capable Transport (0)
[00 38]              Total length (56 = IPv4 header + UDP header + data)
[56 38]              Identification
[00]                 Flags
                          0... .... - reserved bit
                          .0.. .... - don't fragment
                          ..0. .... - more fragments
[00]                 Fragment offset
[80]                 Time to live (128)
[11]                 Protocol (17 = UDP)
[82 77]              Header checksum
[0a e4 22 96]        Source
[0a e4 29 a8]        Destination
```

### User Datagram Protocol

```
[cb bd]              Source port (52157)
[40 00]              Destination port (16384)
[00 24]              Length (36 = UDP header + data)
[f9 bc]              Checksum
```

### Data

`[]bytes("Møte Fr 5.5 14:45 Flåklypa")`
