package main

import (
    "fmt"
    "time"
    "go-chanjs-client/chanjs"
)

func main() {
    opts  := chanjs.Options{
    SocketURL:             "wss://host/api/v1/adfsender/ws/",
    ChannelRef:            "17f8684a4ba933d29fa80e13",
    ChannelSecret:         "SFMyNTY.g2gDaANtAAAAQTE3Zjg2ODRhNGJhOTMzZDI5ZmE4MGUxM2E4ODZmM2Q1LjY5Yzk2YWI3NjU4NjQ1YTI5OGZjMWU5NjQyYmZhNzExbQAAAAZmcm9uZDFtAAAAC1VTRVJARE9NQUlObgYAhK19B5kBYgABUYA.t1L88xr1F6mFcJVR6YN2IkVpa_1Bi0VCIIPEmWneEjo",
    HeartbeatInterval:     200 * time.Millisecond,
    EnableBinaryTransport: false,
    InsecureSkipVerify: true,
}

    client, err := chanjs.NewClient(opts)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
   //client.SendMessage(opts.ChannelSecret)
    //if err != nil {
   // log.Println("Failed to send message:", err)
     //}

    client.ListenEvent("event.some", func(payload interface{}) {
        fmt.Println("Evento recibido:", payload)
    })

//client.Close()
    // Mantener el programa corriendo (por ejemplo, usando un canal vacío)
  //  select {}
}

