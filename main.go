package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"thdr_vault_go/database"
	"thdr_vault_go/http_server"
)

func main() {
	//ln, tcpErr := net.Listen("tcp4", ":8081")
	//if tcpErr != nil {
	//	log.Println("Something went wrong!")
	//}

	port := ":8083"
	dbConn, err := database.InitDatabase("database/store/thdr-vault-database.sql")
	if err != nil {
		log.Printf("Failed to initialize sqlite: %v", err.Error())
	}
	defer dbConn.Close()

	customServer := http_server.NewCustomHttpServer(dbConn)
	mux := customServer.GetMuxHandler()

	server := http.Server{
		Addr: port,
		Handler: mux,
	}
	err = server.ListenAndServe()
	if err
		log.Printf("Something went wrong starting server: %v", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	// Kill all thread when signal sent to process
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	return

	//for {
	//	conn, err := ln.Accept()
	//	if err != nil {
	//		log.Println("Error accepting message")
	//	}

	//	go func(c net.Conn) {
	//		tmp := make([]byte, 256)
	//		conn.Read(tmp)
	//		log.Printf("message: %s", tmp)
	//		io.Copy(c, c)
	//		c.Close()
	//	}(conn)
	//}
}

//func tcpServer() {
//
//}
