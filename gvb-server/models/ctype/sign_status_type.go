package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ     SignStatus = 1 // QQ
	SignGitee  SignStatus = 2 // Gitee
	SignWeChat SignStatus = 3 // 微信
	SignZhiHu  SignStatus = 4 // 知乎
)

func (s SignStatus) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignWeChat:
		str = "微信"
	case SignZhiHu:
		str = "知乎"
	default:
		str = "其他"
	}
	return str
}
