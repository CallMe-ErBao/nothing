package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"pandora/modules/db"
	"pandora/modules/logger"
	"reflect"
)

type UserAttr struct {
	Id        string `json:"id"`
	Token     string `json:"token"`
	Status    string `json:"status"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Mail      string `json:"mail"`
	Note      string `json:"note"`
	Icon      string `json:"icon"`
	Pass      string `json:"pass"`
	Phone     string `json:"phone"`
	CreatTime string `json:"creat_time"`
	Extra1    string `json:"extra1"`
	Extra2    string `json:"extra2"`
}

//将新用户信息插入表中
func CreateAndInsertUser(userAttr UserAttr) string {
	value := reflect.ValueOf(userAttr)
	s := ParseInsertStr("users",value)
	lastId := db.Insert(s)
	fmt.Println(lastId)
	info := QueryUserInfo(userAttr.Phone)
	return info

}

//创建用户信息插入语句
//func CreateUserSQL(userAttr UserAttr) string {
//
//	s := `INSERT INTO users VALUES (`
//	value := reflect.ValueOf(userAttr)
//	ty := reflect.TypeOf(userAttr)
//	for i := 0; i < ty.NumField(); i++ {
//		s = s + `"` + value.Field(i).String() + `"` + ","
//	}
//	s = s[:len(s)-1]
//	s = s + `)`
//
//	return s
//}

//检查请求用户是否已经存在
func CheckUser(phone string) bool {

	result, err := QueryUser(phone)
	if err != nil {
		logger.E(err)
		panic(err)
	}
	if result.Phone == "" {
		return false
	}
	return true

}

//通过电话号码查询用户信息并返回json字符串
func QueryUserInfo(phone string) string {
	userAttr, _ := QueryUser(phone)
	bytes, er := json.Marshal(userAttr)
	if er != nil {
		logger.E(er)
	}
	return string(bytes)
}

//将查询结果逐行取出
func QueryUser(phone string) (attr UserAttr, er error) {
	s := `SELECT * FROM users WHERE phone=` + phone
	query := db.Query(s)
	if query == nil {
		return UserAttr{}, errors.New("没有查询到用户信息")
	}
	var finRe UserAttr
	for query.Next() {
		result := ScanUser(query)
		finRe = *result
	}
	return finRe, nil
}

//取一行数据
func ScanUser(query *sql.Rows) *UserAttr {
	var result UserAttr
	err := query.Scan(&result.Id, &result.Token, &result.Status, &result.Name,
		&result.Role, &result.Mail, &result.Note, &result.Icon,
		&result.Pass, &result.Phone, &result.CreatTime, &result.Extra1, &result.Extra2)
	if err != nil {
		//log.Println(err)
		logger.W(errors.New("没有查询到用户信息"))
		return nil
	}
	return &result
}
