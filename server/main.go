package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris"
	"github.com/majidbigdeli/websocket/websocket"
)

var cho = func(r *http.Request) bool {
	// allow all connections by default
	dd := r.Header.Get("Upgrade")
	fmt.Println(dd)
	return true
}

func main() {
	app := iris.New()
	ws := websocket.New(websocket.Config{CheckOrigin: cho})
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
