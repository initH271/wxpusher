package wxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"wxpusher/config"

	"github.com/google/uuid"
)

type AccessToken struct {
	Value     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

// NOTICE: 持久化失效才重新请求微信服务器获取
// accessToken有效2小时, 考虑存储在数据库
// acToken = jsonMap["access_token"]
// expiresIn := jsonMap["expires_in"] // 7200 seconds
func GetAccessToken() (acToken AccessToken, err error) {
	api := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		config.WxConfig.AppID, config.WxConfig.AppSecret)
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

// Value: 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
// ExpireSeconds: 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
// Url: 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
type Ticket struct {
	Value         string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
}

// action_name	二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
// expire_seconds	该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为60秒。
// action_info	二维码详细信息
//
//	|-- scene
//		|-- scene_str	场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
type ticketRequest struct {
	ActionName    string     `json:"action_name"`
	ActionInfo    actionInfo `json:"action_info"`
	ExpireSeconds int        `json:"expire_seconds"`
}

type actionInfo struct {
	Scene scene `json:"scene"`
}

type scene struct {
	SceneStr string `json:"scene_str"`
}

// 获取临时二维码
func GetTicket() (ticket Ticket, err error) {
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		log.Println("Error generating UUID v1:", err)
		return
	}
	api := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + config.WxConfig.AccessToken
	tReq := &ticketRequest{
		ActionName: "QR_STR_SCENE", // 生成临时二维码
		ActionInfo: actionInfo{
			Scene: scene{
				SceneStr: uuidV1.String(),
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

func SendLoginTemplateMsg(templateId, openId, acToken string) {
	api := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", acToken)
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
