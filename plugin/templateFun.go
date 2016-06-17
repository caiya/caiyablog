//模板函数
package plugin

import (
	"time"

	"github.com/astaxie/beego"
)

func formatTimeStamp(in interface{}, layout string) (out string) {
	timeStr := in.(int)
	month, err := time.Parse(layout, string(timeStr))
	if err != nil {
		return time.Now().Format(layout)
	}
	return month.Format(layout)
}

func init() {
	beego.AddFuncMap("formatTimeStamp", formatTimeStamp)
}
