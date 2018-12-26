package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"pandora/modules/db"
	"pandora/modules/logger"
	"reflect"
)

type attr struct {
	Id           string `json:"id"`
	Type         string `json:"eType"`
	Uid          string `json:"uid"`
	Location_x   string `json:"location_x"`
	Location_y   string `json:"location_y"`
	Uip          string `json:"uip"`
	DType        string `json:"dType"` //device type
	DMan         string `json:"dMan"`
	DTypeVersion string `json:"dTypeVersion"`
	OsVersion    string `json:"osVersion"`
	AppVersion   string `json:"appVersion"`
	AppChannel   string `json:"appChannel"`
	CreatTime    string `json:"creatTime"`
	Extra1       string `json:"extra1"`
	Extra2       string `json:"extra2"`
}

type QueryCondition [][2]string //查询条件 名称：值
type QueryColumn []string       //需要查询的列名

//存储事件
func SaveEvent(evtId, evtAttr string) bool {
	var enAttr attr
	enAttr.Id = evtId
	enErr := json.Unmarshal([]byte(evtAttr), &enAttr)
	if enErr != nil {
		logger.E(enErr)
		return false
	}
	value := reflect.ValueOf(enAttr)
	event := ParseInsertStr("event",value)
	_ = db.Insert(event)
	return true
}

//通过用户ID查询
func QueryEvent(id string) (str []byte, err error) {
	s := `SELECT * FROM event WHERE uid=` + id
	finRe := Query(s)
	return json.Marshal(finRe)
}

//返回插入语句
func ParseInsertStr(tabName string,at reflect.Value) string {

	s := `INSERT INTO `+tabName+` VALUES (`
	for i := 0; i < at.NumField(); i++ {
		s = s + `"` + at.Field(i).String() + `"` + ","
	}
	s = s[:len(s)-1]
	s = s + `)`
	fmt.Println(s)
	return s
}

func QueryEventByType(uid, eType string) (str []byte, err error) {
	//注意空格
	con := QueryCondition{{"uid", uid},
		{"eType", eType}}
	cl := QueryColumn{"*"}
	s := SQLselect(cl, "event", con)
	fmt.Println(s)
	finRe := Query(s)
	return json.Marshal(finRe)
}

//根据查询列，数据库名，查询条件返回查询语句
func SQLselect(args []string, dbName string, condition [][2]string) string {
	s := `SELECT`
	for i := 0; i < len(args); i++ {
		s = s + ` ` + args[i]
	}
	s = s + ` FROM ` + dbName + ` ` + `WHERE `
	for i := 0; i < len(condition); i++ {
		s = s + condition[i][0] + `=` + condition[i][1]
		if i < len(condition)-1 {
			s += ` AND `
		}
	}
	return s

}

//将查询结果逐行取出
func Query(s string) *[]attr {

	query := db.Query(s)
	var finRe []attr
	for query.Next() {
		result := Scan(query)
		//fmt.Println(*result)
		finRe = append(finRe, *result)
	}
	return &finRe
}

//取一行数据
func Scan(query *sql.Rows) *attr {
	var result attr
	err := query.Scan(&result.Id, &result.Type, &result.Uid,
		&result.Location_x, &result.Location_y, &result.Uip, &result.DType,
		&result.DMan, &result.DTypeVersion, &result.OsVersion, &result.AppVersion,
		&result.AppChannel, &result.CreatTime, &result.Extra1, &result.Extra2)
	if err != nil {
		log.Println(err)
		logger.E(errors.New("读取查询数据错误"))
		return new(attr)
	}
	return &result
}
