package chanjs

import "time"

type Options struct {
    SocketURL               string
    ChannelRef              string
    ChannelSecret           string
    HeartbeatInterval       time.Duration
    EnableBinaryTransport   bool
    DedupCacheDisable       bool
    DedupCacheMaxSize       int
    DedupCacheTTL           time.Duration
}
