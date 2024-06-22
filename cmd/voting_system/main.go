package main

import (
  "net/http"
  "fmt"

  vs "vote/voting_system"
)

func main() {
  http.HandleFunc("/", vs.HandleWs)
  fmt.Println("Starting WebSocket server on port 8081")
  vs.Login("random_user", "r@nd0m")
  http.ListenAndServe(":8081", nil)
}
