package logger

import (
	"fmt"
	"log"
	"os"
	"pandora/modules/utils/datetimeutil"
	"runtime"
	"strings"
)

const (
	VERBOS int64 = 1
	DEBUG  int64 = 2
	INFO   int64 = 3
	WARN   int64 = 4
	ERROR  int64 = 5
)

var level = DEBUG

func D(args ...interface{}) {
	if level <= DEBUG {
		out("DEBUG", args)
	}
}
func W(args ...interface{}) {
	if level <= WARN {
		out("WARN", args)
	}
}
func E(args ...interface{}) {
	if level <= ERROR {
		out("ERROR", args)
	}
	//panic(args)
}
func SetLevel(l int64) {
	level = l
}

var mLog *log.Logger
var f = ""
var sep = "/"
var std = log.New(os.Stderr, "", 0)

//记录函数调用信息，输出日志
func out(level string, args []interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	if strings.Contains(file, sep) {
		sr := strings.Split(file, sep)
		file = sr[len(sr)-1]
	}
	fpathstr := "at:" + file + " " + fmt.Sprint(line)
	str := "[" + datetimeutil.FormatDateTimeMillsNow() + "] " + level + "/:"
	for _, s := range args {
		if s == nil {
			s = "nil"
		}
		str += fmt.Sprint(s) + " "
	}
	str += " " + fpathstr

	f1 := "logs" + sep + datetimeutil.FormatDateNow() + ".log"
	if f1 != f {
		f = f1
	}
	if level == "ERROR" {
		std.Println(str)
	} else {
		fmt.Println(str)
	}

	saveToFile(f, str)
}

//将日志保存到文件
func saveToFile(filePath, str string) int64 {
	//fmt.Println("saving", filePath, str)
	userFile := filePath
	fout, err := os.OpenFile(userFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	defer fout.Close()
	if err != nil {
		//fmt.Println(err)
	}
	n, er := fout.WriteString(str + "\n")
	if er != nil {
		n = 0
		//fmt.Println(er.Error())
	}
	return int64(n)
}
