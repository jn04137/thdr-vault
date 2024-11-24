package main

import (
	//"io"
	"log"
	"net/http"
	//"net"

	"thdr_vault_go/database"
	"thdr_vault_go/http_server"
)

func main() {
	//ln, tcpErr := net.Listen("tcp4", ":8081")
	//if tcpErr != nil {
	//	log.Println("Something went wrong!")
	//}

	dbConn, err := database.InitDatabase("database/store/thdr-vault-database.sql")
	if err != nil {
		log.Printf("Failed to initialize sqlite: %v", err.Error())
	}
	defer dbConn.Close()

	httpServer := http_server.NewCustomHttpServer(dbConn)
	mux := httpServer.HttpServer()
	http.ListenAndServe(":8080", mux)

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

func tcpServer() {

}
