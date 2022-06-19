package server

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type streamingServer struct {
	Host    string
	Port    int
	Command []string
}

func NewStreamingServer(host string, port int, cmd []string) WebsocketServer {
	return &streamingServer{
		Host:    host,
		Port:    port,
		Command: cmd,
	}
}

func (ss *streamingServer) Addr() string {
	return fmt.Sprintf("%s:%d", ss.Host, ss.Port)
}

func (ss *streamingServer) Start() error {
	log.Println("Starting websocket server:", ss.Addr())
	server := http.Server{
		Addr:    ss.Addr(),
		Handler: ss,
	}

	return server.ListenAndServe()
}

func (ss *streamingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ERROR:", "[upgrade]", err)
		return
	}
	defer c.Close()

	cmdLine := strings.Join(ss.Command, " ")
	cmd := exec.Command("sh", "-c", cmdLine)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return
	}

	err = cmd.Start()
	log.Println("executing command:", cmd.String())
	if err != nil {
		log.Println(err)
		return
	}

	// print the output of the subprocess
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		data := scanner.Text()
		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("ERROR:", "[write]:", err)
		}
	}
	_ = cmd.Wait()
	log.Println("command completed")
}
