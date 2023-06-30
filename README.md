# brainwave
闪念笔记

## gRPC 使用说明
### 编译
#### 编译全部
```bash
./gen.sh
```
#### 编译指定文件
```bash
./gen.sh socket.proto
```
> [!TIP]
mac sed 低版本无 -r生成客户端代码失败
使用gsed代替 sed
https://stackoverflow.com/questions/73574692/sed-illegal-option-r-usage-error-in-visual-studio-code/73601736#73601736


## 在项目中使用
1. 初始化客户端
```go
    cfgs := []client.ClientConfig{
        {
            ServiceName: service.Service_socket_Socket,
            ServiceAddr: "127.0.0.1:13041",
        },
    }
    client.InitClients(cfgs...)
```
2. 访问服务
```go
    req := &socket.SendReq{}
	resp, err := socket.SendRpc(ctx, req)
```