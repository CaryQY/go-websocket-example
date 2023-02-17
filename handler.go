package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"sync"
)

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"mode": mode, "domain": domain})
}

var wsUpgrader = websocket.Upgrader{}

var globalConns = make([]*websocket.Conn, 0)
var lk = sync.RWMutex{}

func wsHandler(c *gin.Context) {
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": "failed to set websocket upgrade"})
		return
	}

	lk.Lock()
	defer lk.Unlock()
	globalConns = append(globalConns, conn)
}

type Message struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func apiHandler(c *gin.Context) {
	id := c.Query("id")
	value := c.Query("value")

	msg := Message{
		Id:    id,
		Value: value,
	}

	data, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	lk.RLock()
	defer lk.RUnlock()
	end := len(globalConns)
	for i := 0; i < end; i++ {
		conn := globalConns[i]
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err)
			conn = nil
			globalConns[i], globalConns[end-1] = globalConns[end-1], globalConns[i]
			end = end - 1
			i--
		}
	}
	if end < len(globalConns) {
		globalConns = globalConns[:end]
	}
	fmt.Printf("total conns: %d\n", len(globalConns))
}

func staticFileHandler(c *gin.Context) {
	file, err := f.ReadFile("rest_api_example.png")
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.Data(200, "image/png", file)
}
