package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pandora/modules/logger"
	"pandora/modules/utils/apitool"
	"services"
)

type CommonApi struct {
}

func NewCommonApi() *CommonApi {
	c := &CommonApi{}
	return c
}
func (a *CommonApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	apiHandler(r, w)
}

func apiHandler(r *http.Request, w http.ResponseWriter) {
	var cmd = apitool.Get("c", r)
	switch cmd {
	case "ReportEvent":
		reportEvent(r, w)
		break
	case "Query":
		queryEvent(r, w)
		break
	case "QueryByEventType":
		QueryByEventType(r, w)
		break
	case "ReportUser":
		reportUser(r, w)
		break
	case "ReportProject":
		reportProject(r, w)
		break
	case "DeleteProject":
		deleteProject(r,w)
	}
}

func reportEvent(r *http.Request, w http.ResponseWriter) {
	evtId := apitool.Get("evtId", r)
	evtAttrJSON := apitool.Get("evtAttr", r)
	ok := services.SaveEvent(evtId, evtAttrJSON)
	if ok {
		apitool.ApiResponse(r, w, 0, "", "")
	} else {
		apitool.ApiResponse(r, w, -1, "fail", "")
	}
}

func queryEvent(r *http.Request, w http.ResponseWriter) {
	uid := apitool.Get("uid", r)
	result, ok := services.QueryEvent(uid)
	if ok == nil {
		apitool.ApiResponse(r, w, 0, "", string(result))
	} else {
		apitool.ApiResponse(r, w, -1, "fail", "")
	}
}

func QueryByEventType(r *http.Request, w http.ResponseWriter) {
	uid := apitool.Get("uid", r)
	etype := apitool.Get("eType", r)
	result, err := services.QueryEventByType(uid, etype)
	if err == nil {
		apitool.ApiResponse(r, w, 0, "", string(result))
	} else {
		apitool.ApiResponse(r, w, -1, "fail", "")
	}

}

func reportUser(r *http.Request, w http.ResponseWriter) {
	userJson := apitool.Get("userAttr", r)
	var usAttr services.UserAttr
	usErr := json.Unmarshal([]byte(userJson), &usAttr)
	if usErr != nil {
		logger.E(usErr)
	}
	fmt.Println(userJson)
	if usAttr.Phone == "" { //用户手机号为空，直接返回错误
		apitool.ApiResponse(r, w, -1, "手机号为空", "")
	} else if ok := services.CheckUser(usAttr.Phone); ok { //老登录，返回用户信息
		info := services.QueryUserInfo(usAttr.Phone)
		apitool.ApiResponse(r, w, 0, "老用户", info)
	} else { //新用户
		re := services.CreateAndInsertUser(usAttr)
		apitool.ApiResponse(r, w, 0, "新用户", re)
	}
}

func reportProject(r *http.Request, w http.ResponseWriter) {
	AttrJson:= apitool.Get("ProjectAttr", r)
	result ,ok:= services.AddProject(AttrJson)
	if ok == "" {
		apitool.ApiResponse(r, w, 0, "项目创建成功", result)
	}else {
		apitool.ApiResponse(r, w, -1, ok, "")
	}


}
func deleteProject(r *http.Request, w http.ResponseWriter)  {
	AttrJson:= apitool.Get("ProjectAttr", r)
	ok := services.DeleteProject(AttrJson)
	if ok == "" {
		apitool.ApiResponse(r,w,0,"删除成功","")
	}else {
		apitool.ApiResponse(r,w,-1,ok,"")

	}

}