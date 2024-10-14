package Luminance

import (
	"context"
	"github.com/coder/websocket"
)

// ** Package Structure **

type Package struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
}

// ** Bot Init Structure **

type Bot struct {
	Conn         *websocket.Conn
	Username     string `json:"username"`
	Key          string `json:"key"`
	WebsocketURI string `json:"websocket_uri"`
	CTX          context.Context
}

// ** Websocket Structure **

type OnConnection struct {
	OP       int      `json:"op"`
	Messages []string `json:"messages"`
}

type OnMessage struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

type SendMessage struct {
	Content  string `json:"content"`
	Username string `json:"username"`
	Key      string `json:"key"`
}
