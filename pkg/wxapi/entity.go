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
