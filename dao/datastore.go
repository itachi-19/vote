package dao

import (
	"encoding/hex"
  "crypto/md5"

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
  []VotingSession{NewVotingSession("best_stock"), NewVotingSession("worst_stock") }

var Users map[string]User = map[string]User{
  "ram.0114": NewUser("ram.0114", "Abc", "Ram Verma"),
  "shyam_24": NewUser("shyam_24", "vdf", "Shyam Kumar"),
}

var AuthZTokens map[string][]string = map[string][]string{ // allow access to specific voting sessions
  "token_xyz": {"best_stock", "worst_stock"},
  "token_abc": {"worst_stock"},
}


type VotingSession struct {
  name string
  subscribers map[User]bool
}

type ConnectionSession struct {
  conn *websocket.Conn // TODO VALIDATE STATE before interacting with voting system
  token string // 
}


type User struct {
  Username string
  Name string
  PasswordHash string
  Token string
}

func NewVotingSession(name string) VotingSession {
  var subs map[User]bool
  return VotingSession{name: name, subscribers:subs} 
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


/*
establish conn
register conn in set
login
user pass -> get token or invalid 
*/
