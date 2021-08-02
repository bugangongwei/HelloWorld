package verification

import (
	"fmt"
	"time"
)

/* 关于 time 包的用法 */

func ParseSec() {
	t := time.Unix(1626333388, 0).Format("200601")
	fmt.Println(t)
}
