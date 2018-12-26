package messages

import (
	"encoding/base64"
	"encoding/json"
	"pandora/vars"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
)

var coder = base64.NewEncoding(base64Table)

type Message struct {
	Id   string `json:"id"`
	Type int    `json:"type"`
	Ori  string `json:"ori"`
	Data string `json:"data"`
	Tar  string `json:"tar"`
}

type Method struct {
	MethodName   string `json:"methodName"`
	MethodParams string `json:"methodParams"`
}

func (m *Message) ToString() string {
	m.Data = string(base64Encode([]byte(m.Data)))
	b, _ := json.Marshal(m)
	return "[" + string(b) + "]"
}

func (m *Message) GetDecodedData() string {

	res, err := base64Decode([]byte(m.Data))
	if err != nil {
		return "消息解码失败:" + err.Error() + m.Data
	} else {
		return string(res)
	}

}

func NewRegisterModuleMsg(id string, ori string) Message {
	var r Message
	r.Id = id
	r.Type = vars.MsgType_registerModule
	r.Ori = ori
	return r
}

func NewRequestMsg(id string, ori string, targetModule string, methodName string, methodParams string) Message {
	var r Message
	r.Id = id
	r.Type = vars.MsgType_request
	r.Tar = targetModule
	r.Ori = ori
	var method Method
	method.MethodName = methodName
	method.MethodParams = methodParams
	b, _ := json.Marshal(method)
	r.Data = string(b)
	return r
}

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}
