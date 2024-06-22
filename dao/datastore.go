package dao

import "github.com/gorilla/websocket"

type DAO struct {
//TODO REDIS
  /*
    voting_session : subscribed_connections map[VotingSession]WebsocketConnections
    WebsocketConnections: map[connections]bool
*/ 
}

//(dao *DAO) func GetConnections(votingSessions VotingSeession, )

type VotingSession struct {
  id string 
  name string
  subscribers map[User]bool
}

type User struct {
  name string
  id string // EPOCH_name
  conn *websocket.Conn // TODO VALIDATE STATE before interacting with voting system
}
