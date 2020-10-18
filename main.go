package main

import (
	"flag"
	broker "github.com/IT108/achieve-broker-go"
	"github.com/IT108/achieve-gateway-go/processor"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var clientsHub *Hub
var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	 http.ServeFile(w, r, "static/home.html")
}

func shutdownService()  {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		os.Exit(0)
	}()
}

func main() {
	processor.GenerateGateId()
	shutdownService()
	broker.ConfigureFromEnv()
	flag.Parse()
	clientsHub = newHub()
	go clientsHub.run()
	startConsumer()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(clientsHub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
