package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {

	//sec, err := setting.Cfg.GetSection("database")
	//if err != nil {
	//	fmt.Printf("Fail to get section 'database': %v", err)
	//	return
	//}
	//
	//dbName := sec.Key("NAME").String()
	//user := sec.Key("USER").String()
	//password := sec.Key("PASSWORD").String()
	//host := sec.Key("HOST").String()

	dbName := ""
	user := ""
	password := ""
	host := ""

	dsn := user + ":" + password + "@tcp" + "(" + host + ")" + "/" + dbName + "?charset=utf8"
	fmt.Println(dsn) // root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8
	err := converter.NewTable2Struct().
		SavePath("/Users/luxuefeng/go/gin-simple/models/version.go").
		Dsn(dsn).
		TagKey("gorm").
		EnableJsonTag(true).
		Table("blog_auth").
		Run()
	if err != nil {
		fmt.Println(err)
	}

}

func GenerateStruct(path string, dsn string, table string) {
	//err := converter.NewTable2Struct().
	//	SavePath("./model.go").
	//	Dsn("用户名:密码@tcp(IP:端口号)/数据库名?charset=utf8").
	//	TagKey("gorm").
	//	EnableJsonTag(true).
	//	Table("表名").
	//	Run()
	//fmt.Println(err)
	err := converter.NewTable2Struct().
		SavePath(path).
		Dsn(dsn).
		TagKey("gorm").
		EnableJsonTag(true).
		Table(table).
		Run()
	fmt.Println(err)
}
