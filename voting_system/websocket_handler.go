package voting_system

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
  "context" 
  "vote/dao"
  "strings"
)

var upgrader = websocket.Upgrader{}
var connections = map[*websocket.Conn]bool{}
var messages = make(chan []byte, 1000)

func read(conn *websocket.Conn) string {
  _, msg, _ := conn.ReadMessage()
  return strings.TrimSpace(string(msg))
}

func write(conn *websocket.Conn, msg string) {
  conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
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

    var user *dao.User
    for {
      _, message, err := conn.ReadMessage()
      msg := strings.TrimSpace(string(message))
      if err != nil {
        break
      }
      //write(conn, msg)
      if msg == "login" {
        write(conn, "Ack")
        write(conn, "Enter username: ")
        username := read(conn)
        write(conn, "Enter password: ")
        password := read(conn)
        user = Login(username, password)


        if user == nil {
          write(conn, "Auth Failed! Try again")
        } else {
          write(conn, fmt.Sprintf("Auth Successful!. Welcome %v !", user.Name))
        }


      } else if msg == "logout" {
        // Close connection
        break
      } else if msg == "register" {
        //TODO
      } else if msg == "vote" {

      } else if msg == "create_vote_session" {

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


