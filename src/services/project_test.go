package services

import (
	"fmt"
	"testing"
)

func Test_Project(t *testing.T) {
	s:=`
{
"Id":"123445",
"name":"wallace",
"ownerId":"0192384",
"states":"01",
"deleteFlag":"1",
"balance":"123",
"extra0":"123",
"extra1":"123",
"extra2":"123"
}
`
	fmt.Println(s)
}
