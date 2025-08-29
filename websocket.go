package chanjs

import (
    "encoding/json"
    "github.com/gorilla/websocket"
)

type EventMessage struct {
    Name      string          `json:"name"`
    MessageID string          `json:"message_id"`
    Payload   json.RawMessage `json:"payload"`
}

func parseEventMessage(msgBytes []byte) EventMessage {
    var evt EventMessage
    _ = json.Unmarshal(msgBytes, &evt)
    return evt
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
