package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type subscribingServer struct {
	Host         string
	Port         int
	RedisHost    string
	RedisPort    int
	RedisChannel string
}

func NewSubscribingServer(host string, port int, rHost string, rPort int, channel string) WebsocketServer {
	return &subscribingServer{
		Host:         host,
		Port:         port,
		RedisHost:    rHost,
		RedisPort:    rPort,
		RedisChannel: channel,
	}
}

func (ss subscribingServer) Start() error {
	log.Println("Starting subscribing websocket server:", ss.Addr())
	server := http.Server{
		Addr:    ss.Addr(),
		Handler: ss,
	}

	return server.ListenAndServe()
}

func (ss subscribingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ERROR:", "[upgrade]", err)
		return
	}
	defer c.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     ss.RedisAddr(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Println("connected to redis at", ss.RedisAddr())

	ctx := context.Background()
	pubsub := rdb.Subscribe(ctx, ss.RedisChannel)

	log.Println("subscribed to redis channel", ss.RedisChannel)

	// Close the subscription when we are done.
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("ERROR:", "[pubsub]", err)
			break
		}

		c.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	}
}

func (ss subscribingServer) Addr() string {
	return fmt.Sprintf("%s:%d", ss.Host, ss.Port)
}

func (ss subscribingServer) RedisAddr() string {
	return fmt.Sprintf("%s:%d", ss.RedisHost, ss.RedisPort)
}
