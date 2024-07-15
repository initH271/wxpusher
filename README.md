# 微信公众号开发实例

已使用接口实例:
- 验证域名URL回调
- 获取AccessToken
- 生成临时二维码
- 接收公众号订阅/取消订阅/扫码事件推送回调
- 发送模板消息

## 快速开始

.env文件示例, 如果不配置环境变量,那么应该在项目或可执行文件目录下创建.env文件

```conf
APP_ID=wx7cd7fdsadasd179bc
APP_SECRET=f8816ee02dasdasdasdasf47f656d3
TOKEN=2zeCr3TcoxaxaxsadasfadfasdasGz
ENCODING_AES_KEY=Dy9potcczZPOKJsssfDHk95IF

ACCESS_TOKEN_TMP=

PORT=8080
```

Linux服务器中: 
1. 编译项目
```sh
go mod tidy && go build .
```
2. 登录微信公众号后台, 根据`.env`文件添加对应环境变量

3. 运行`./wxpusher`, 并解析域名到8080端口

4. 在微信公众号后台配置服务器URL与Token


Windows 本地机器调试
1. 登录微信公众号后台, 修改`.env`文件中对应环境变量

2. 运行项目
```sh
go run main.go
```
3. 使用内网穿透将本机8080端口穿透到服务器, 并为服务器端口配置域名解析

4. . 在微信公众号后台配置服务器URL与Token
