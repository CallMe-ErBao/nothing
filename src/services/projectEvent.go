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

var ProjectSql = `
create table IF not exists project(
     id int(64) not null  auto_increment,
	 name varchar(128),
	 ownerID varchar int(11) not null，
     status varchar(64) default 1,
     deleteFlag varchar(2),
     balance varchar(64)，
	 extra0 varchar(128),
	 extra1 varchar(128),
	 extra2 varchar(128)
    )
`

type ProAttr struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	OwnerID    string `json:"ownerID"`
	States     string `json:"states"`
	DeleteFlag string `json:"deleteFlag"`
	Balance    string `json:"balance"`
	Extra0     string `json:"extra0"`
	Extra1     string `json:"extra1"`
	Extra2     string `json:"extra2"`
}

//添加项目
func AddProject(JsonAttr string) (res, ok string) {

	var attr ProAttr
	err := json.Unmarshal([]byte(JsonAttr), &attr)
	if err != nil {
		logger.E(err)
	}
	proAt, er := QueryProject(attr.OwnerID, attr.Name)
	if er != nil {
		logger.E(er)
		return "", "查询失败"
	}
	if proAt != nil {
		return "", "项目名称已经存在"
	}

	value := reflect.ValueOf(attr)

	str := ParseInsertStr("project", value)
	lastId := db.Insert(str)
	fmt.Println(lastId)
	proAttr, err := QueryProject(attr.OwnerID, attr.Name)
	fmt.Println()
	if err != nil {
		logger.E(err)
		panic(err)
	}
	bytesAttr, e := json.Marshal(proAttr)
	if e != nil {
		logger.E(e)
	}

	return string(bytesAttr), ""

}

//func CreateProSQL(PA ProAttr) string {
//
//	s := `INSERT INTO project VALUES (`
//	value := reflect.ValueOf(PA)
//	ty := reflect.TypeOf(PA)
//	for i := 0; i < ty.NumField(); i++ {
//		s = s + `"` + value.Field(i).String() + `"` + ","
//	}
//	s = s[:len(s)-1]
//	s = s + `)`
//
//	return s
//}

//查询项目
func QueryProject(ownerId, name string) (attr []ProAttr, er error) {
	s := `SELECT * FROM project WHERE name="` + name + `" AND ownerId=` + ownerId + ` AND deleteFlag=` + "1"
	query := db.Query(s)
	if query == nil {
		return []ProAttr{}, errors.New("查询出错")
	}
	var finRe []ProAttr
	for query.Next() {
		result := ScanProject(query)
		finRe = append(finRe, *result)
	}
	return finRe, nil
}

//取一行数据
func ScanProject(query *sql.Rows) *ProAttr {
	var result ProAttr
	err := query.Scan(&result.Id, &result.Name, &result.OwnerID,
		&result.States, &result.DeleteFlag, &result.Balance, &result.Extra0, &result.Extra1, &result.Extra2)
	if err != nil {
		//log.Println(err)
		logger.W(errors.New("没有查询到项目信息"))
		return nil
	}
	return &result
}

//删除项目
func DeleteProject(JsonAttr string) string {
	var attr ProAttr
	err := json.Unmarshal([]byte(JsonAttr), &attr)
	if err != nil {
		logger.E(err)
	}
	proAt, er := QueryProject(attr.OwnerID, attr.Name)
	if er != nil {
		logger.E(er)
		return "查询项目失败"
	}
	if proAt == nil {
		return "所删除项目不存在"
	}

	s := `update project set deleteFlag=0 where id="` + attr.Id + `"`
	num := db.Update(s)
	if num == 0 {
		return "删除失败"
	} else {
		return ""
	}

}
