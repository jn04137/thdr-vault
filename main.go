package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	ln, tcpErr := net.Listen("tcp4", ":8081")
	if tcpErr != nil {
		log.Println("Something went wrong!")
	}

	go httpServer()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting message")
		}

		go func(c net.Conn) {
			tmp := make([]byte, 256)
			conn.Read(tmp)
			log.Printf("message: %s", tmp)
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}

func httpServer() {
	port := ":8080"
	http.HandleFunc("/", httpHandler)

	log.Printf("Starting http server at port: %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi there! I love %v", r.URL.Path[1:])
}

func tcpServer() {

}
