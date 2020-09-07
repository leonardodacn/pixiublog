package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

/**
* 提供函数给template使用
 */

// NumEq eq int and interface{}, interface is type of float64
// see: http://www.seaxiang.com/blog/Zz6zEj
func NumEq(a int, b interface{}) bool {
	if b == nil {
		return false
	}
	c := 0
	switch b.(type) {
	case string:
		if d, err := strconv.Atoi(b.(string)); err == nil {
			c = d
		} else {
			return false
		}
	case float32:
		c = int(b.(float32))
	case float64:
		c = int(b.(float64))
	case int:
		c = b.(int)
	case int32:
		c = int(b.(int32))
	case int64:
		c = int(b.(int64))
	}
	if a == c {
		return true
	}
	return false
}

func ModEq(a, b, c int) bool {
	if a < b {
		return false
	}
	if a%b == c {
		return true
	}
	return false
}

// ContainStr 判断a是否包含b，默认的分隔符使用 , 例子： a = "12,34,56"; b = 4; 结果为false
func ContainNum(a string, b int, separator ...string) bool {
	if len(a) == 0 {
		return false
	}
	sep := ","
	if len(separator) > 0 {
		sep = separator[0]
	}
	a = sep + a + sep
	c := sep + strconv.Itoa(b) + sep
	if strings.Contains(a, c) {
		return true
	}
	return false
}

func RandNum(max, min int) int {
	if max <= min {
		fmt.Printf("param error, max: %d, min: %d", max, min)
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	for {
		v := rand.Intn(max)
		if v >= min {
			return v
		}
	}
}

func Add(a, b int) int {
	return a + b
}

func DateStr(time time.Time) string {
	d := beego.Date(time, "Y-m-d")
	return d
}
