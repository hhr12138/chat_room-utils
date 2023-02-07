package encrypt

import (
	"fmt"
	"reflect"
	"testing"
)

type NullString struct {
	String string
	Valid  bool
}

type Test struct {
	Name NullString
}

type Test1 struct {
	Name string
}

type User struct {
	Id     int64
	Name   string
	Weight float64
	Height float64
}

func TestSqlObjToObj(t *testing.T) {
	test := &Test{
		Name: NullString{
			String: "zhangsan",
			Valid:  true,
		},
	}
	test1 := new(Test1)
	reflect.ValueOf(test1).Elem().FieldByName("Name").SetString("haha")
	SqlObjToObj(test, test1)
	fmt.Println(test1)
}

func TestEncryptPassword(t *testing.T) {
	password := EncryptPassword("hhr12138", "qwer")
	fmt.Println(password)
}
