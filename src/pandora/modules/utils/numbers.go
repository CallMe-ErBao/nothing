package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GetInt64(str string) int64 {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	n, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	return n
}
func GetInt(str string) int {

	return int(GetInt64(str))
}
func GetFloat64(str string) float64 {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	n, e := strconv.ParseFloat(str, 64)
	if e != nil {
		fmt.Println(e.Error())
	}
	return n
}
