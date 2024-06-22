package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
  "context" 
)

var upgrader = websocket.Upgrader{}
var connections = map[*websocket.Conn]bool{}
var messages = make(chan []byte, 1000)

func handleWs(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  connections[conn] = true
  go func() {
    defer func() {
      delete(connections, conn)
      conn.Close()
    }()
    for {
      messageType, message, err := conn.ReadMessage()
      fmt.Print(messageType, string(message))
      conn.WriteMessage(websocket.TextMessage, []byte("Your vote recorded"))
      if err != nil {

        break
      }
    }
  }()

}

func broadcastUpdateToClients(ctx context.Context, updateMessage []byte) {
  // TODO
  updateMessage = nil
  ctx = nil
  // ======================================
}


func main() {
  http.HandleFunc("/", handleWs)
  fmt.Println("Starting WebSocket server on port 8081")
  http.ListenAndServe(":8081", nil)
}
