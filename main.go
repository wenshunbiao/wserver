package main

import (
	"log"
	"net/http"
	"wserver/conf"
	wserver "wserver/server"
	"wserver/utils"
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
	server.AuthToken = func(token string, conn *wserver.Conn) (userID string, ok bool) {
		// 不配置授权key 直接返回token作为用户id
		if serverConf.WsAuthKey == "" {
			return token, true
		}

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			log.Println(err)
			conn.Write([]byte(err.Error()))
			return "", false
		}

		if claims.Subject == "" {
			log.Println("sub is required")
			conn.Write([]byte("sub is required"))
			return "", false
		}

		log.Println("user auth success: " + claims.Subject)
		return claims.Subject, true
	}

	// Set PushAuth func to check push request. If the request is valid, returns
	// true. Otherwise return false and request will be ignored.
	server.PushAuth = func(r *http.Request) bool {
		// 不配置授权key 直接通过
		if serverConf.PushAuthKey == "" {
			return true
		}

		if serverConf.PushAuthKey != r.Header.Get("Authorization") {
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
