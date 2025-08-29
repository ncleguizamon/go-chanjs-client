package chanjs

import (
    "github.com/gorilla/websocket"
)

type WebSocketConn struct {
    conn *websocket.Conn
}

func NewWebSocketConn(opts Options) (*WebSocketConn, error) {
    conn, _, err := websocket.DefaultDialer.Dial(opts.SocketURL, nil)
    if err != nil {
        return nil, err
    }
    return &WebSocketConn{conn: conn}, nil
}

func (w *WebSocketConn) SendPing() error {
    return w.conn.WriteMessage(websocket.PingMessage, nil)
}

func (w *WebSocketConn) Listen(handler func(EventMessage)) {
    for {
        _, msgBytes, err := w.conn.ReadMessage()
        if err != nil {
            return
        }
        evtMsg := parseEventMessage(msgBytes)
        handler(evtMsg)
    }
}

func (w *WebSocketConn) Close() error {
    return w.conn.Close()
}
