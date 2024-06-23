package dao

import (
	"crypto/md5"
	"encoding/hex"
	"time"
  "fmt"

	"github.com/gorilla/websocket"
)

type DAO struct {
//TODO REDIS
  /*
    voting_session : subscribed_connections map[VotingSession]WebsocketConnections
    WebsocketConnections: map[connections]bool
*/ 
}

//(dao *DAO) func GetConnections(votingSessions VotingSeession, )

var VotingSesssions []VotingSession = 
  []VotingSession{NewVotingSession("best_smartphone"), NewVotingSession("best_tv") }

var Ram User = NewUser("ram.0114", "Abc", "Ram Verma")
var Shyam User = NewUser("shyam_24", "vdf", "Shyam Kumar")

var Users map[string]User = map[string]User{
  "ram.0114": Ram,
  "shyam_24": Shyam,
}

var AuthZTokens map[string][]string = map[string][]string{ // allow access to specific voting sessions
  "token_xyz": {"best_smartphone", "best_tv"},
  "token_abc": {"best_tv"},
}

var UserConnections map[string]*websocket.Conn = make(map[string]*websocket.Conn)


type VotingSession struct {
  Name string
  Votes map[string]int
  Subscribers map[User]bool
}

type User struct {
  Username string
  Name string
  PasswordHash string
  Token string
}

func Broadcast(votingSession VotingSession) {
  for user := range votingSession.Subscribers {
    notifyUser(user, votingSession.Name, votingSession.Votes)
  }
}

func notifyUser(user User, name string, votes map[string]int) {
  conn, err := UserConnections[user.Username]
  if err {
    return
  }
  conn.WriteMessage(websocket.TextMessage,[]byte("Results for " + name + ": "))
  for k, v := range votes {
    conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf("%v -> %v", k, v)))
  }
}

func NewVotingSession(name string) VotingSession {
  var subs map[User]bool = make(map[User]bool)
  subs[Ram] = true
  subs[Shyam] = true
  var votes map[string]int = make(map[string]int)
  votes["default"] = 100

  delay := 20 * time.Second
  vs := VotingSession{Name: name, Subscribers:subs, Votes: votes} 

  time.AfterFunc(delay, func() {
    //Broadcast(vs)
  })

  return vs
}

func NewUser(username string, passwd string, name string) User {
  hash := md5.Sum([]byte(passwd))
  token := username + "_233fdfFEFMXCX_token"
  return User{Name: name, Username: username, PasswordHash: hex.EncodeToString(hash[:]), Token: token}
}

func GetUser(username string) User {
  return Users[username]
}

func GetAuthZToken(username string) string {
  return GetUser(username).Token
}

func RegisterNewUser(username string, passwd string, name string) *User {
  user := NewUser(username, passwd, name)
  Users[username] = user
  return &user
}
