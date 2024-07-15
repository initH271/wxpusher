# 微信公众号开发实例

已完成接口功能:
- 验证域名URL回调
- 获取AccessToken
- 生成临时二维码
- 接收公众号订阅/取消订阅/扫码事件推送回调
- 发送模板消息
- 微信扫码登录demo

## 快速开始

wxconfig.toml配置文件示例, 如果不配置环境变量,那么应该在项目或可执行文件目录下创建配置文件文件

```toml
[WxConfig]
AppID = "wx7cdsadasdas79bc"
AppSecret = "f8816edasdasdas4ff8f47f656d3"
Token = "2zeCr3Tco8J8dasdasdasBHMFiGz"
EncodingAESKey = " "

AccessToken = "82_CAB55wxIq472ZjMLKTdasdasdasVZ2EdKQUuT8IkW0iJ486cdsadasdasUt_JDzA30X-6kLM5ygEFYZgAHATMN"

[Server]
Port = 8080
```

> NOTE: 出现access token过期的情况, 可以先请求/wx/getAccessToken接口刷新一下

使用redis缓存用户扫码登录状态, **运行Redis在端口16379**:
```sh
docker compose up
```

**Linux服务器中**: 
1. 编译项目
```sh
go mod tidy && go build .
```
2. 登录微信公众号后台, 根据`wxconfig.toml`文件添加对应环境变量

3. 运行`./wxpusher`, 并解析域名到8080端口

4. 在微信公众号后台配置服务器URL与Token


**Windows 本地机器调试**:
1. 登录微信公众号后台, 修改`wxconfig.toml`文件中对应环境变量

2. 运行项目
```sh
go run main.go
```
3. 使用内网穿透将本机8080端口穿透到服务器, 并为服务器端口配置域名解析

4. . 在微信公众号后台配置服务器URL与Token

**微信扫码登录demo**: 浏览器访问 /public接口