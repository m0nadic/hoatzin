package server

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"hoatzin/internal/app/utils/functions"
	"log"
	"net/http"
	"text/template"
	"time"
)

type tickingTemplateServer struct {
	Host            string
	Port            int
	MessageTemplate string
	Interval        int
	Count           int
}

func NewTickingTemplateServer(host string, port int, template string, interval int, count int) WebsocketServer {
	return &tickingTemplateServer{
		Host:            host,
		Port:            port,
		MessageTemplate: template,
		Interval:        interval,
		Count:           count,
	}
}

func (tts *tickingTemplateServer) Addr() string {
	return fmt.Sprintf("%s:%d", tts.Host, tts.Port)
}

func (tts *tickingTemplateServer) Start() error {
	log.Println("Starting websocket server:", tts.Addr())
	server := http.Server{
		Addr:    tts.Addr(),
		Handler: tts,
	}

	return server.ListenAndServe()
}

func (tts *tickingTemplateServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ERROR:", "[upgrade]", err)
		return
	}
	defer c.Close()

	ticker := time.NewTicker(time.Duration(tts.Interval) * time.Millisecond)

	log.Println("sending total", tts.Count, "messages, 1 in every", tts.Interval, "milliseconds")
	for i := 0; i < tts.Count; i++ {
		select {
		case _ = <-ticker.C:
			writeTemplate(c, tts.MessageTemplate)
		}
	}

}

func writeTemplate(c *websocket.Conn, mTemplate string) {

	tt := template.New("message")
	tt.Funcs(functions.GlobalFunctions())
	_, err := tt.Parse(mTemplate)
	if err != nil {
		log.Print("ERROR:", "[template]", err)
		return
	}
	var data bytes.Buffer
	err = tt.Execute(&data, nil)
	if err != nil {
		log.Print("ERROR:", "[template]", err)
		return
	}
	err = c.WriteMessage(websocket.TextMessage, data.Bytes())
	if err != nil {
		log.Println("ERROR:", "[write]:", err)
	}
}
