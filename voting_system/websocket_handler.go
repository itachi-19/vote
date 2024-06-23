package voting_system

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
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
    prompt := "[-]:$ "
    for {
      write(conn, prompt + "Enter a command")
      _, message, err := conn.ReadMessage()
      msg := strings.TrimSpace(string(message))
      if err != nil {
        break
      }

      if msg == "login" {
        write(conn, prompt + "Enter username: ")
        username := read(conn)
        write(conn, prompt + "Enter password: ")
        password := read(conn)
        user = Login(username, password)

        if user == nil {
          write(conn, prompt + "Auth Failed! Try again")
        } else {
          prompt = "[" + user.Username + "]:$ "
          write(conn, prompt + fmt.Sprintf("Auth Successful!. Welcome %v !", user.Name))
          dao.UserConnections[username] = conn
        }
      } else if msg == "logout" {
        // Close connection
        write(conn, prompt + "Logging out!...")
        conn.Close()

        return
      } else if msg == "register" {
        write(conn, prompt + "Enter username: ")
        username := read(conn)
        write(conn, prompt + "Enter password: ")
        password := read(conn)
        write(conn, prompt + "Enter full name: ")
        name := read(conn)
        user = dao.RegisterNewUser(username, password, name)
        prompt = "[" + user.Username + "]:$ "
        write(conn, prompt + "User: " + user.Name + " registered!")
        dao.UserConnections[username] = conn
      } else if msg == "vote" {

      } else if msg == "create_vote_session" {
        write(conn, prompt + "Enter name of Voting Session:")
        name := read(conn)
        v_session := dao.NewVotingSession(name)
        dao.VotingSesssions = append(dao.VotingSesssions, v_session)

        write(conn, prompt + "Voting session created for: " + name)
      }
    }
  }()

}

