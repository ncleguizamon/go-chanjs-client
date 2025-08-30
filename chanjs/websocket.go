package chanjs

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "crypto/tls"
    "log"
)

type EventMessage struct {
    Name      string          `json:"name"`
    MessageID string          `json:"message_id"`
    Payload   json.RawMessage `json:"message_data"`
}

func parseEventMessage(msgBytes []byte) EventMessage {
    var evt EventMessage
    _ = json.Unmarshal(msgBytes, &evt)
    return evt
}

func NewWebSocketConn(opts Options) (*WebSocketConn, error) {
    websocket.DefaultDialer.TLSClientConfig = &tls.Config{
        InsecureSkipVerify: opts.InsecureSkipVerify, // ⚠️ true só para testes!
    }
   log.Println("👂 Iniciando escuta de mensagens WebSocket..." + opts.SocketURL )

   conn, resp,  err := websocket.DefaultDialer.Dial(opts.SocketURL, nil)
    
    if err != nil {
        log.Fatalf("Erro ao conectar: %v (HTTP status: %v)", err, resp.Status)
        return nil, err
    }
    return &WebSocketConn{conn: conn}, nil
}

func (w *WebSocketConn) SendPing() error {
   // log.Println("👂 SendPing" )
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

func (w *WebSocketConn) Send(message string) error {
    // Asumiendo que w.conn es *websocket.Conn de gorilla/websocket
    err := w.conn.WriteMessage(websocket.TextMessage, []byte(message))
    if err != nil {
        log.Println("Error enviando mensaje:", err)
        return err
    }
    log.Println("Mensaje enviado:", message)
    return nil
}



func (w *WebSocketConn) Close() error {
    return w.conn.Close()
}
