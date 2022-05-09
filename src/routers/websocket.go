package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"golang.org/x/exp/slices"
	"log"
)

var WebsocketConnections []*websocket.Conn

type WebsocketMessage[D interface{}] struct {
	Type string `json:"type"`
	Data D      `json:"data"`
}

func handleWebsocket(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/connect", websocket.New(func(c *websocket.Conn) {
		WebsocketConnections = append(WebsocketConnections, c)

		// c.Locals are added to the *websocket.Conn
		//log.Println(c.Locals("allowed"))  // true
		//log.Println(c.Query("v"))         // 1.0
		//log.Println(c.Cookies("session")) // ""

		c.SetCloseHandler(func(code int, text string) error {
			log.Println("close:", code, text)
			index := slices.Index(WebsocketConnections, c)
			WebsocketConnections = slices.Delete(WebsocketConnections, index, index+1)
			return nil
		})

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			err error
		)

		for {
			if _, _, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
		}
	}))
}

func SendWebsocketBroadcast[D interface{}](message WebsocketMessage[D]) {
	for _, conn := range WebsocketConnections {
		_ = conn.WriteJSON(message)
	}
}
