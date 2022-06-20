package server

import (
	"fmt"
	"log"
	"net/http"
)

type echoServer struct {
	Host string
	Port int
}

func NewEchoServer(host string, port int) WebsocketServer {
	return &echoServer{
		Host: host,
		Port: port,
	}
}

func (es *echoServer) Addr() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}

func (es *echoServer) Start() error {
	log.Println("Starting websocket server:", es.Addr())
	server := http.Server{
		Addr:    es.Addr(),
		Handler: es,
	}

	return server.ListenAndServe()
}

func (es *echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ERROR:", "[upgrade]", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("ERROR:", "[read]", err)
			break
		}
		log.Println("recv:", string(message))

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("ERROR:", "[write]", err)
			break
		}
	}

}
