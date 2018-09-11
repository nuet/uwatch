package excel

import (
	"testing"
	"fmt"
)

func TestImport(t *testing.T)  {
	//列号映射成key
	cellmaps := []string{"empno", "username", "realname", "dep_1", "dep_2", "dep_3", "dep_4", "dep_5", "position", "phone", "email", "comp"}

	startColumn := 1 //起始列，第1列
	endColumn := 12 //结束列
	startRow := 19 //起始行,第2行
	filename := "./user.xlsx"
	res,err := Import(filename, cellmaps, startColumn, endColumn, startRow)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(len(res))
	fmt.Println(res[0])
}