package main

import (
    "fmt"
    "strings"
    "time"

    irc "github.com/fluffle/goirc/client"
)

func main() {
    // Or, create a config and fiddle with it first:
    cfg := irc.NewConfig("skybot")
    cfg.Server = "129.12.4.54"
    cfg.NewNick = func(n string) string { return n + "^" }
    c := irc.Client(cfg)
    quit := make(chan bool)
    ticker := time.NewTicker(time.Millisecond * 5000)

    // Add handlers to do things here!
    // e.g. join a channel on connect.
    c.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
      conn.Join("#skybot")

      // Setup a ticker.
      go func() {
          for t := range ticker.C {
              fmt.Println("Tick at", t)
          }
      }()
    })

    // Add handlers to do things here!
    // e.g. join a channel on connect.
    c.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
      command := line.Text()
      if isme := strings.Index(command, "skybot: "); isme != 0 {
        return
      }

      command = command[8:]
      switch command {
        case "quit":
          if line.Nick == "sky" {
            quit <- true
          }
        case "info":
          conn.Privmsg(line.Target(), "SkyBot version 0.1")
        default:
          fmt.Printf(command)
      }
    })

    // And a signal on disconnect
    c.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
      quit <- true
    })

    // Tell client to connect.
    if err := c.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    // Wait for disconnect
    <-quit
    ticker.Stop()
}
