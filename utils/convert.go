package utils

import "fmt"

//查看数据类型
func DataType(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
