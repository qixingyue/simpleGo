package libs

import (
	"fmt"
)

func insertSql(tableName string, fieldValues map[string]interface{}) {
	kvStr := ""
	for k, v := range fieldValues {
		if sv, ok := v.(string); ok {
			kvStr = fmt.Sprintf("%s,`%s`=\"%s\"", kvStr, k, sv)
		}
		if iv, ok := v.(int); ok {
			kvStr = fmt.Sprintf("%s,`%s`=%d", kvStr, k, iv)
		}
	}
	kvStr = kvStr[1:]
	sql := fmt.Sprintf("INSERT INTO `%s` SET %s ", tableName, kvStr)
	fmt.Println(sql)
}

//func main() {
//
//	//	    -> id int primary key not null auto_increment,
//	//		    -> path varchar(200),
//	//			    -> realname varchar(200),
//	//				    -> description text,
//	//					    -> mime varchar(10)
//	//
//	m := make(map[string]interface{})
//	m["path"] = ""
//	m["realname"] = "hello "
//	m["description"] = "new"
//	m["mime"] = "mp3"
//	insertDb("info", m)
//
//}
