package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

type restStreamingServer struct {
	Host string
	Port int
}

func NewRESTStreamingServer(host string, port int) WebsocketServer {
	return &restStreamingServer{
		Host: host,
		Port: port,
	}
}

type WSHandler struct {
	Messages chan string
}

func (ws *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ERROR:", "[upgrade]", err)
		return
	}
	defer c.Close()

	for {
		m := <-ws.Messages
		_ = c.WriteMessage(websocket.TextMessage, []byte(m))
	}
}

type APIHandler struct {
	Messages chan string
}

func (api *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("ERROR:", "[API]", err)
		return
	}
	api.Messages <- string(message)
	_, _ = w.Write([]byte("OK"))
}

func (rss *restStreamingServer) Start() error {
	log.Println("Starting REST streaming websocket server:", rss.Addr())
	mux := http.NewServeMux()
	ch := make(chan string, 10)
	mux.Handle("/", &WSHandler{Messages: ch})
	mux.Handle("/api/v1/messages", &APIHandler{Messages: ch})
	server := http.Server{
		Addr:    rss.Addr(),
		Handler: mux,
	}

	return server.ListenAndServe()
}

func (rss *restStreamingServer) Addr() string {
	return fmt.Sprintf("%s:%d", rss.Host, rss.Port)
}
