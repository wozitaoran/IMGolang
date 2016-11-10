package proto

import (
	"encoding/json"
)

//接受的消息结构
type RecvMessage struct {
	Target_type string `json:"target_type"`
	Msg_type    int64  `json:"msg_type"`
	Msg         string `json:"msg"`
	From        int64  `json:"from"`
	Send_time   string `json:"send_time"`
}

//发送的消息结构
type SendMessage struct {
	Target_type string          `json:"target_type"`
	Target      int64           `json:"target"`
	Msg_type    int64           `json:"msg_type"`
	Msg         json.RawMessage `json:"msg"`
	From        int64           `json:"from"`
}
