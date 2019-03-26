package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/majidbigdeli/websocket/websocket"
)

func main() {
	app := iris.New()
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(func(c websocket.Connection) {
		//go func() {
		//	<-time.After(20 * time.Second)
		//	c.Disconnect()
		//}()

		c.On("chat", func(message string) {
			c.To(websocket.Broadcast).Emit("chat", c.ID()+": "+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("Connection with ID: %s has been disconnected!\n", c.ID())
		})
	})

	app.Get("/socket", ws.Handler())

	app.Run(iris.Addr(":8080"))
}
