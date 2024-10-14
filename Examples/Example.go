package main

import (
	"runtime"
	"strings"
	"github.com/Kodarru/Luminance"
)

func main() {
    // Initialize the bot with your credentials and WebSocket URI
    bot := &Luminance.Bot{
        Username:     "your_username_here",
        Key:          "your_key_here",
        WebsocketURI: "wss://guhws.nin0.dev",
    }

    // Initialize the bot connection
    bot.Initialize()

    var pkg Luminance.Package = Luminance.GetPackage()
	bot.SendMessage("Bot is ready! Running on " + runtime.Version() + " with " + pkg.Name + " v" + pkg.Version)

    // Define a callback function to handle incoming messages
    messageHandler := func(msg Luminance.OnMessage) {
        if strings.HasPrefix(msg.Content, "!") {
			command := strings.Split(msg.Content, "!")[1]
			switch command {
				case "ping":
					bot.SendMessage("Pong!")
				case "hello":
					bot.SendMessage("Hello, " + msg.Username + "!")
				case "close":
					bot.SendMessage("Goodbye!")
					bot.Close()
				default:
					bot.SendMessage("Command not found")
			} 
		}
    }

    // Start listening for messages
    bot.OnMessage(messageHandler)
}