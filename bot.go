package main

import (
    //"crypto/tls"
    "fmt"

    irc "github.com/fluffle/goirc/client"
)

func main() {
    // Or, create a config and fiddle with it first:
    cfg := irc.NewConfig("skybot")
    //cfg.SSL = true
    //cfg.SSLConfig = &tls.Config{ServerName: "irc.cs.kent.ac.uk"}
    cfg.Server = "129.12.4.54"
    cfg.NewNick = func(n string) string { return n + "^" }
    c := irc.Client(cfg)

    // Add handlers to do things here!
    // e.g. join a channel on connect.
    c.HandleFunc(irc.CONNECTED,
        func(conn *irc.Conn, line *irc.Line) { conn.Join("#skybot") })

    // And a signal on disconnect
    quit := make(chan bool)
    c.HandleFunc(irc.DISCONNECTED,
        func(conn *irc.Conn, line *irc.Line) { quit <- true })

    // Tell client to connect.
    if err := c.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    // Wait for disconnect
    <-quit
}
