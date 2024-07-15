package wxapi

// 微信接口回调推送XML数据包
// ToUserName	开发者 微信号
// FromUserName	发送方账号（一个OpenID，此时发送方是系统账号）
// CreateTime	消息创建时间 （整型），时间戳
// MsgType	消息类型，event
// Event	事件类型 qualification_verify_success
// ExpiredTime	有效期 (整型)，指的是时间戳，将于该时间戳认证过期
type WxEvent struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

type AccessToken struct {
	Value     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

// Value: 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
// ExpireSeconds: 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
// Url: 二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
type Ticket struct {
	Value         string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	SceneStr      string `json:"scene_str"`
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
