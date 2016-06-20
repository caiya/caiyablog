//模板函数
package plugin

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
)

func formatTimeStamp(in int, layout string) (out string) {
	month, err := time.Parse(layout, string(in))
	if err != nil {
		return time.Now().Format(layout)
	}
	return month.Format(layout)
}

func formatToURL(in int) (out string) {
	mtime := time.Unix(int64(in), 0)
	s := mtime.Format("2006-01-02")
	sarr := strings.Split(s, "-")
	return sarr[0] + "/" + sarr[1] + "/" + sarr[2]
}

func init() {
	beego.AddFuncMap("formatTimeStamp", formatTimeStamp)
	beego.AddFuncMap("formatToURL", formatToURL)
}
