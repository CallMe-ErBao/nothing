package apitool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ResultObject struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

///tools
func JsonP(req *http.Request, w http.ResponseWriter, result string) {
	//logger.D("api", "jsonp", result)
	w.Header().Add("Content-type", "text/json; charset=utf-8")
	w.Write([]byte(Get("callback", req) + "(" + result + ")"))
}
func Response(msg string, w http.ResponseWriter) {
	w.Header().Add("Content-type", "text/json; charset=utf-8")
	w.Header().Add("Cache-control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(msg))
}
func GenJsonSucc(result interface{}) string {
	return GenJsonStatus("success", result)
}
func GenJsonStatus(status string, result interface{}) string {
	r := &ResultObject{}
	r.Result = result
	r.Status = status
	b, _ := json.Marshal(r)
	return string(b)
}
func ApiResponse(r *http.Request, w http.ResponseWriter, errorCode int64, errMsg, content string) string {
	if strings.HasPrefix(content, "{") || strings.HasPrefix(content, "[") {
		//json形式不处理
	} else {
		content = `"` + content + `"`
	}

	s := `{"code":` + fmt.Sprint(errorCode) + `,"msg":"` + errMsg + `","data":` + content + `}`
	if Get("reqType", r) == "jsonp" {
		JsonP(r, w, s)
	} else {
		Response(s, w)
	}
	return s
}
func GetWithCheck(key string, r *http.Request) string {
	re := strings.Trim(r.Form.Get(key), " ")
	return checkSafeSql(re)
}
func Get(key string, r *http.Request) string {
	return strings.Trim(r.Form.Get(key), " ")
}

//简单防注入
func checkSafeSql(s string) string {
	if strings.Contains(s, "--") {
		s = ""
	}
	return s
}
