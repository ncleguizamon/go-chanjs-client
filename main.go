package main

import (
    "fmt"
    "time"
    "go-chanjs-client/chanjs"
)

func main() {
    opts  := chanjs.Options{
    SocketURL:             "wss://financiacion-int-qa.apps.ambientesbc.com/saf/api/v1/adfsender/ws/?channel=17f8684a4ba933d29fa80e13a886f3d5.c6d26fe8f3a64e36b893a72cc1bc47ea",
    ChannelRef:            "17f8684a4ba933d29fa80e13a886f3d5.c6d26fe8f3a64e36b893a72cc1bc47ea",
    ChannelSecret:         "Auth::SFMyNTY.g2gDaANtAAAAQTE3Zjg2ODRhNGJhOTMzZDI5ZmE4MGUxM2E4ODZmM2Q1LmM2ZDI2ZmU4ZjNhNjRlMzZiODkzYTcyY2MxYmM0N2VhbQAAAAZmcm9uZDFtAAAAC1VTRVJARE9NQUlObgYAbHZm95gBYgABUYA.l8GCYS8tgqansoxGvP2OrY4WTkNZh3s3fhfRL0vKCTw",
    HeartbeatInterval:     200 * time.Millisecond,
    EnableBinaryTransport: false,
    InsecureSkipVerify: true,
}

    client, err := chanjs.NewClient(opts)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
   client.SendMessage(opts.ChannelSecret)
    //if err != nil {
   // log.Println("Failed to send message:", err)
     //}

    client.ListenEvent("event.some", func(payload interface{}) {
        fmt.Println("Evento recibido:", payload)
    })

    // Mantener el programa corriendo (por ejemplo, usando un canal vacío)
  //  select {}
}

