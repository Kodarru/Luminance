package Luminance

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/coder/websocket"
)

// *** Configurations ***

var pkg = Package{
    Name:    "Luminance",
    Author: "Kodarru",
    Version: "1.0.0",
}

var bot Bot
var hasReceivedMessage = false

func New(username, key, websocket_uri string) {
	bot.Username = username
	bot.Key = key
	bot.WebsocketURI = websocket_uri
	bot.CTX = context.Background()
}

func GetPackage() Package {
	return pkg
}

// *** Main ***

func (bot *Bot) Initialize() {
    // Initialize context
    bot.CTX = context.Background()

    // Connect to WebSocket
    conn, _, err := websocket.Dial(bot.CTX, bot.WebsocketURI, nil)
    if err != nil {
        fmt.Println("Error connecting to the websocket: ", err)
        return
    }

    bot.Conn = conn
}

// *** Websocket ***

func (bot *Bot) SendMessage(content string) {
	message := SendMessage{Content: content, Username: bot.Username, Key: bot.Key}
	messageJSON, _ := json.Marshal(message)
    bot.Conn.Write(bot.CTX, websocket.MessageText, messageJSON)
}

func (bot *Bot) OnMessage(callback func(OnMessage)) {
    if bot.Conn == nil || bot.CTX == nil {
        fmt.Println("WebSocket connection or context is not initialized")
        return
    }

    for {
        var msg OnMessage
		
        _, message, err := bot.Conn.Read(bot.CTX)
        if err != nil {
            fmt.Println("Error reading message: ", err)
            return
        }
        err = json.Unmarshal(message, &msg)
        if err != nil {
            fmt.Println("Error unmarshalling message: ", err)
            return
        }
        callback(msg)
    }
}

func (bot *Bot) Close() {
    bot.Conn.Close(websocket.StatusNormalClosure, "closing connection")
}