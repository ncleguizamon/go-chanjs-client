package chanjs

import (
    "time"
    "github.com/gorilla/websocket"
    "log"
)



type WebSocketConn struct {
    conn *websocket.Conn
}



type AsyncClient struct {
    opts   Options
    ws     *WebSocketConn
    dedup  *DedupCache
}



func NewClient(opts Options) (*AsyncClient, error) {
    wsc, err := NewWebSocketConn(opts)
    if err != nil {
        return nil, err
    }

    var dc *DedupCache
    if !opts.DedupCacheDisable {
        dc = NewDedupCache(opts.DedupCacheMaxSize, opts.DedupCacheTTL)
    }

    client := &AsyncClient{
        opts:  opts,
        ws:    wsc,
        dedup: dc,
    }
    go client.runHeartbeat()
    return client, nil
}
    

func (c *AsyncClient) runHeartbeat() {
    ticker := time.NewTicker(c.opts.HeartbeatInterval)
    defer ticker.Stop()
    for range ticker.C {
        c.ws.SendPing()
    }
}

func (c *AsyncClient) Close() error {
    return c.ws.Close()
}

func (c *AsyncClient) ListenEvent(pattern string, callback func(interface{})) {
   log.Println("👂 Iniciando escuta de mensagens WebSocket..." + pattern)
    c.ws.Listen(func(evt EventMessage) {
        log.Println("EventMessage" + evt.MessageID)
        if matchPattern(pattern, evt.Name) {
            if c.dedup != nil && c.dedup.IsDuplicate(evt.MessageID) {
                return
            }
            callback(evt.Payload)
        }
    })
}
func (c *AsyncClient) SendMessage(message string) error  {
    // Assuming c.ws has a Send or Write method; replace with your actual method
    err := c.ws.Send(message)
    if err != nil {
        log.Println("Error sending message:", err)
        return err
    }
    log.Println("Message sent successfully")
    return nil
}
