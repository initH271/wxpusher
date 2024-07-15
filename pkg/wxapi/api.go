package wxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// NOTICE: 持久化失效才重新请求微信服务器获取
// accessToken有效2小时, 考虑存储在数据库
// acToken = jsonMap["access_token"]
// expiresIn := jsonMap["expires_in"] // 7200 seconds
func GetAccessToken(appID, appSecret string) (acToken AccessToken, err error) {
	api := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		appID, appSecret)
	resp, err := http.Get(api)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 解析JSON响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading the response body:", err)
		return
	}
	err = json.Unmarshal(body, &acToken)
	if err != nil {
		log.Println("Error decoding the JSON:", err)
		return
	}
	return
}

// 获取临时二维码
func GetTicket(accessToken string) (ticket Ticket, err error) {
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		log.Println("Error generating UUID v1:", err)
		return
	}
	api := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + accessToken

	ticket.SceneStr = uuidV1.String()
	tReq := &ticketRequest{
		ActionName: "QR_STR_SCENE", // 生成临时二维码
		ActionInfo: actionInfo{
			Scene: scene{
				SceneStr: ticket.SceneStr,
			},
		},
		ExpireSeconds: 120,
	}
	data, err := json.Marshal(tReq)
	if err != nil {
		log.Println("Error encoding the JSON:", err)
		return
	}

	resp, err := http.Post(api, "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Println("Error http.Post:", err)
		return
	}
	defer resp.Body.Close()

	// 解析JSON响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading the response body:", err)
		return
	}
	err = json.Unmarshal(body, &ticket)
	if err != nil {
		log.Println("Error decoding the JSON:", err)
		return
	}
	if ticket.ErrCode != 0 {
		log.Println("Error get ticket:", ticket.ErrMsg)
		return
	}
	return

}

func SendLoginTemplateMsg(templateId, openId, accessToken string) {
	api := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	jsonStr := fmt.Sprintf(`{
    "touser": "%s",
    "template_id": "%s",
    "url": "https://inith271.top",
    "topcolor": "#FF0000",
    "data": {}
}`, openId, templateId)

	_, err := http.Post(api, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Println("模板消息发送失败.")
	}

}
