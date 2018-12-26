package services

import (
	"fmt"
	"testing"
)

func Test_event(T *testing.T) {
	//	s := `{
	//		"uid":"147258",
	//		"eType":"01",
	//		"location_x":"42.123",
	//		"location_y":"121.45",
	//		"uip":"192.168.1.1",
	//       "dType":"01",
	//		"DMan":"2015",
	//		"DTypeVersion":"mac",
	//		"osVersion":"os",
	//		"AppVersion":"1.0.0",
	//		"AppChannel":"pc",
	//		"Extra1":"123",
	//		"Extra2":"456"
	//}`
	con := [][2]string{
		{"uid", "12"},
		{"eType", "34"}}
	cl := []string{"*"}
	s := SQLselect(cl, "event", con)
	fmt.Println(s)
}
