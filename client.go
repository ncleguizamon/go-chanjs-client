package chanjs

import (
    "context"
    "time"
)

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
    c.ws.Listen(func(evt EventMessage) {
        if matchPattern(pattern, evt.Name) {
            if c.dedup != nil && c.dedup.IsDuplicate(evt.MessageID) {
                return
            }
            callback(evt.Payload)
        }
    })
}
