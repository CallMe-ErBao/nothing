package services

import (
	"fmt"
	"testing"
)

func Test_user(T *testing.T) {
	s := `{
			"id":"147258",
	 		"token":"4hk34vb5jk2bk3",
			"status":"01",
			"name":"wallace",
			"role":"master",
			"mail":"hfeihjf@qq.com",
	        "note":"01",
			"icon":"2015",
			"pass":"mac",
			"phone":"12345657894",
			"create_time":"20180109",
			"extra1":"123",
			"extra2":"456"
	}`
	fmt.Println(s)
}
