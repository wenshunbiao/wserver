# wserver

[![Build Status](https://www.travis-ci.org/alfred-zhong/wserver.svg?branch=master)](https://www.travis-ci.org/alfred-zhong/wserver) [![GoDoc](https://godoc.org/github.com/alfred-zhong/wserver?status.svg)](https://godoc.org/github.com/alfred-zhong/wserver) [![Go Report Card](https://goreportcard.com/badge/github.com/alfred-zhong/wserver)](https://goreportcard.com/report/github.com/alfred-zhong/wserver)

轻量级的 Websocket 消息推送服务

## 授权

### HTTP 推送授权机制

在 `header` 传递 `Authorization` 请求头，如：

```shell
curl -X POST -H "Content-Type: application/json" -H "Authorization: zouaR7n1ZoX2YPZFjyYRX6Lu7vCbS82D" -d '{"userId":"all","Event":"topic1","Message":"Hello in 2021-09-24 16:07:55.966"}' -i http://127.0.0.1:12345/push
```

### Websocket 连接注册消息的token授权机制

采用标准的 [JWT -- JSON WEB TOKEN](https://jwt.io/) 来构建 token，其载荷如：

```json
{
  "sub": "user3688",
  "iss": "wserver",
  "iat": 1632473835,
  "exp": 1632560253,
  "nbf": 1632473835,
  "jti": "37c107e4609ddbcc9c096ea5ee76c667",
  "aud": "dev"
}
```

其中 `sub` 是必须的，它是你的用户ID。其他字段都是可选的，但 `wserver` 服务接受并验证 `iat`、`nbf`、`exp` 字段，并建议你传递 `exp` 字段以控制你的 `JWT` 有效期。

## Basic Usage

Try to start a wserver, you just need to write like this.

### Start wserver

```
go run main.go
```

Now wserver listens on port: 12345.  
启动后会自动在当前目录生成配置文件 `configs/app.toml`

### Browser connecting

Now browser can connect to `ws://ip:12345/ws`. After connection established, browser should send a message to register. Register message looks like this.

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI1NTEzODksInN1YiI6InVzZXIzNjg4In0.EVK26QeHd0d31ZAK0J9xY9wmBAhBRvm5U1sS80D3vIM",
    "event": "topic1"
}
```

The `token` is used for identification and `event` means what kind of messages the client interested (Like topic in MQ). 

### Push messages

Now you can send a request to `http:/ip:12345/push` to push a message. Message should look like this.

```json
{
    "userId": "user3688",
    "event": "topic1",
    "message": "Hello World"
}
```

The `userId` is equal to token is not specified (customize by using `server.AuthToken`). The `event` is equal to that above and `message` is the real content will be sent to each websocket connection.

If you want to run a demo. Then follow the steps below:

* run `go run main.go` to start a wserver.
* Open the webpage **_examples/index.html** in the browser.
* run `go run _examples/push.go` to send some messages.

If all success, you will see the content like this:

![demo-success](./demo-success.png)

## PS

Package is still not stable and I will improve it then. 

PRs will be welcomed. 🍺
