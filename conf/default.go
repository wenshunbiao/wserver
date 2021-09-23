package conf

const DefaultConfig = `[Server]
# http推送授权key
PushAuthKey="{PushAuthKey}"

# ws连接授权key  AES加密算法,建议长度32
WsAuthKey="{WsAuthKey}"

# 服务监听地址
Addr="0.0.0.0:12345"
`
