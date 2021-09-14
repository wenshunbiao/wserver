package main

import (
	"fmt"
	"log"
	"net/http"
	"wserver/conf"
	wserver "wserver/server"
)

func main() {
	serverConf := conf.AppConf.Server
	server := wserver.NewServer(serverConf.Addr)

	// Define websocket connect url, default "/ws"
	server.WSPath = "/ws"
	// Define push message url, default "/push"
	server.PushPath = "/push"

	// Set AuthToken func to authorize websocket connection, token is sent by
	// client for registe.
	server.AuthToken = func(token string) (userID string, ok bool) {
		// TODO: check if token is valid and calculate userID
		//if token == "aaa" {
		//	return "jack", true
		//}
		//
		//return "", false
		fmt.Println("user: " + token + " connect")
		return token, true
	}

	// Set PushAuth func to check push request. If the request is valid, returns
	// true. Otherwise return false and request will be ignored.
	server.PushAuth = func(r *http.Request) bool {
		// TODO: check if request is valid
		if serverConf.PushAuthToken != r.Header.Get("Authorization") {
			return false
		}

		return true
	}

	// Run server
	log.Printf("Server: %s  WSPath: %s  PushPath: %s", server.Addr, server.WSPath, server.PushPath)
	log.Println("启动服务...")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
