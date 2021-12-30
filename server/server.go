package server

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v7"
)

var rdb *redis.Client

type counter struct {
	v int
}

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func echo(w http.ResponseWriter, r *http.Request, globalQuit chan struct{}, hub *hub, f func(data string, c *websocket.Conn, mt int)) {
// 	c, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}
// 	defer c.Close()
// 	for {
// 		mt, message, err := c.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		if mt == 1 {
// 			f(string(message), c, mt)
// 			// log.Printf("recv: %s", message)
// 			// if string(message) == "\"123\"" {
// 			// 	c.WriteMessage(mt, []byte("got it"))
// 			// 	log.Println("send:", []byte("got it"))
// 			// }
// 		}
// 	}

// }

// func ServerRun(addr string, f func(data string, c *websocket.Conn, mt int)) {
// 	log.Println("Starting game server: ", addr)

// 	counterCh := make(chan counter)
// 	globalQuit := make(chan struct{})
// 	hub := NewHub(counterCh, globalQuit)

// 	defer close(globalQuit)

// 	go hub.start()
// 	go updateCounterEvery(time.Second, counterCh)

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		echo(w, r, globalQuit, hub, f)
// 	})

// 	err := http.ListenAndServe(addr, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func updateCounterEvery(d time.Duration, counterCh chan counter) {
// 	c := counter{}
// 	ticker := time.NewTicker(d)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			c.v++
// 			counterCh <- c
// 		}
// 	}
// }

func ServerRun(addr string) {
	log.Println("Starting game server: ", addr)
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
