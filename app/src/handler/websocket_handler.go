package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/shinya-ac/GoChat/domain"
)

type WebsocketHandler struct {
	hub *domain.Hub
}

func NewWebsocketHandler(hub *domain.Hub) *WebsocketHandler {
	return &WebsocketHandler{
		hub: hub,
	}
}

func (h *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := domain.NewClient(ws)
	go client.ReadLoop(h.hub.BroadcastCh, h.hub.UnRegisterCh)
	go client.WriteLoop()
	h.hub.RegisterCh <- client

	// 初回メッセージ
	//以下のコードはこの記事から→https://qiita.com/hiro_nico/items/db6cb98916fdf3e8c4cc
	// err = ws.WriteMessage(websocket.TextMessage, []byte(`Server (gorilla): Hello, Client!`))
	// if err != nil {
	// 	log.Println("WriteMessage:", err)
	// 	return
	// }

	// for {
	// 	//mt=message tupe
	// 	mt, message, err := ws.ReadMessage()
	// 	if err != nil {
	// 		log.Println("ReadMessage:", err)
	// 		break
	// 	}
	// 	err = ws.WriteMessage(mt, []byte(fmt.Sprintf("Server (gorilla): '%s' received.", message)))
	// 	if err != nil {
	// 		log.Println("WirteMessage:", err)
	// 		break
	// 	}
	// }

}
