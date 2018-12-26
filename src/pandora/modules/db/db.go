package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"pandora/modules/logger"
	"strconv"
	"strings"
	"sync"
)

var currDb *sql.DB = nil
var lock sync.Mutex
var MysqlURL string

func getDb() *sql.DB {
	if currDb != nil {
		return currDb
	}
	db, err := sql.Open("mysql", MysqlURL)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(20)
	if checkErr(err) {
		logger.E(err)
		return nil
	}
	currDb = db
	return db
}

//插入demo
func Insert(sql string, args ...interface{}) int64 {
	lock.Lock()
	db := getDb()
	if db == nil {
		lock.Unlock()
		return -100
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		logger.E(err)
		lock.Unlock()
		return -100
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		logger.E(err)
		stmt.Close()
		lock.Unlock()
		return -100
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.E(err)
		stmt.Close()
		lock.Unlock()
		return -100
	}
	stmt.Close()
	lock.Unlock()
	return id
}

func Test() {
	logger.D("db", "testing")
}
func GetRealSql(sql string, args ...interface{}) string {
	str := sql
	for _, v := range args {
		vs := fmt.Sprint(v)
		vs = strings.Trim(vs, "\n")
		vs = strings.TrimSpace(vs)
		_, e := strconv.ParseInt(vs, 10, 64)
		if e != nil {
			vs = "'" + vs + "'"
		}
		str = strings.Replace(str, "?", vs, 1)
	}
	return str
}

//查询demo
func Query(sql string, args ...interface{}) *sql.Rows {
	db := getDb()
	if db == nil {
		logger.E("db is nil")
		return nil
	}
	//fmt.Println(db)
	//logger.D("db", "querying:", GetRealSql(sql, args...))
	stmt, err := db.Prepare(sql)
	if err != nil {
		logger.E(err)
		return nil
	}
	//fmt.Println(stmt)
	row, err := stmt.Query(args...)
	if err != nil {
		logger.E(err)
		stmt.Close()
		return nil
	}
	defer stmt.Close()
	return row
}

//获取唯一结果，如count 等
func QueryUnigueResult(sql string, args ...interface{}) string {
	r := Query(sql, args...)
	if r == nil {
		return ""
	}
	var s string
	for r.Next() {
		r.Scan(&s)
	}
	return s
}
func QuerySqlToJSON(jsonKeys []string, sql string, args ...interface{}) string {
	logger.D("prepare query", sql)
	r := Query(sql, args...)
	//logger.D("query end")
	if r == nil {
		return "[]"
	}
	return convertRowsToJSON(r, jsonKeys)
}
func ConvertQueryToJSON(rows *sql.Rows, jsonKeys ...string) string {
	return convertRowsToJSON(rows, jsonKeys)
}
func convertRowsToJSON(rows *sql.Rows, jsonKeys []string) string {
	cols, _ := rows.Columns()
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	values := make([][]byte, len(cols))
	for i := range values {
		dest[i] = &values[i]
	}
	//var data [][]map[string]interface{}
	var results []map[string]interface{}
	for rows.Next() {
		rows.Scan(dest...)
		//logger.D("db",dest)
		m := make(map[string]interface{})
		for i, k := range jsonKeys {
			//logger.D(values[i])
			m[k] = string(values[i])

		}
		results = append(results, m)
	}
	if len(results) == 0 {
		return "[]"
	}
	s, _ := json.Marshal(results)
	rows.Close()
	return string(s)
}

//更新数据
func Update(sql string, args ...interface{}) int64 {

	lock.Lock()
	db := getDb()
	if db == nil {
		lock.Unlock()
		return -100
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		logger.E(err)
		lock.Unlock()
		return -100
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		logger.E(err)
		lock.Unlock()
		stmt.Close()
		return -100
	}
	num, err := res.RowsAffected()
	if err != nil {
		logger.E(err)
		stmt.Close()
		lock.Unlock()
		return -100
	}
	stmt.Close()
	lock.Unlock()
	return num
}

//删除数据
func Remove(sql string, args ...interface{}) int64 {
	lock.Lock()
	db := getDb()
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(args...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	stmt.Close()
	lock.Unlock()
	return num
}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		logger.E(err)
		return true
	}
	return false
}
