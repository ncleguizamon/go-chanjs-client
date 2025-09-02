package chanjs

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "crypto/tls"
    "net/url"
    "log"
    "fmt"
)

type EventPayload struct {
	Message string `json:"message"`
}

type EventMessage struct {
	ID        string       // "7c1b3e5b-285e-4ebd-adb2-8b531c5c797f"
	UserID    string       // "95ca4275-599a-46c7-97e9-46e664d96b1a"
	EventName string       // "event.some"
	Payload   EventPayload // {"message": "..."}
}

/*func parseEventMessage(msgBytes []byte) EventMessage {
    var evt EventMessage
    _ = json.Unmarshal(msgBytes, &evt)
    return evt
}
    */

    func ParseEventMessage(data []byte) (*EventMessage, error) {
    var raw []json.RawMessage
   
    if err := json.Unmarshal(data, &raw); err != nil {
        return nil, fmt.Errorf("error unmarshalling raw JSON array: %w", err)
    }

    if len(raw) != 4 {
        return nil, fmt.Errorf("expected array of 4 elements, got %d", len(raw))
    }

    var em EventMessage

    if err := json.Unmarshal(raw[0], &em.ID); err != nil {
        return nil, fmt.Errorf("error parsing ID: %w", err)
    }
    if err := json.Unmarshal(raw[1], &em.UserID); err != nil {
        return nil, fmt.Errorf("error parsing UserID: %w", err)
    }
    if err := json.Unmarshal(raw[2], &em.EventName); err != nil {
        return nil, fmt.Errorf("error parsing EventName: %w", err)
    }
    if em.EventName != "AuthOk" {
     firstChar := raw[3][0]

    if firstChar == '{' {
        if err := json.Unmarshal(raw[3], &em.Payload); err != nil {
            return nil, fmt.Errorf("error parsing Payload: %w", err)
        }
    } else {
        return nil, fmt.Errorf("unexpected JSON token for Payload: %c", firstChar)
    }
    }
    return &em, nil
}


func NewWebSocketConn(opts Options) (*WebSocketConn, error) {
    websocket.DefaultDialer.TLSClientConfig = &tls.Config{
        InsecureSkipVerify: opts.InsecureSkipVerify, // ⚠️ true só para testes!
    }
   log.Println("👂 Iniciando escuta de mensagens WebSocket..." + opts.SocketURL )
    // Add query parameters
    u, err := url.Parse(opts.SocketURL)
    if err != nil {
        panic(err)
    }
    query := u.Query()
    query.Set("channel", opts.ChannelRef)
    u.RawQuery = query.Encode() 
    //

    conn, resp,  err := websocket.DefaultDialer.Dial(u.String(), nil)
    
    if err != nil {
        log.Fatalf("Erro  conectar: %v (HTTP status: %v)", err, resp.Status)
        return nil, err
    }
    return &WebSocketConn{conn: conn}, nil
}

func (w *WebSocketConn) SendPing() error {
   // log.Println("👂 SendPing" )
    return w.conn.WriteMessage(websocket.PingMessage, nil)
}


func (w *WebSocketConn) Listen(handler func(*EventMessage)) {
    for {
        _, msgBytes, err := w.conn.ReadMessage()
        if err != nil {
            return
        }
        evtMsg, err := ParseEventMessage(msgBytes)
        if err != nil {
            log.Println("Error parsing event message:", err)
            continue
        }
        handler(evtMsg) // pasa puntero directamente
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
