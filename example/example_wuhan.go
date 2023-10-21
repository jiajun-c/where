package example

import (
	"fmt"
	"github.com/jiajun-c/where"
)

func main() {
	ans, _ := where.Where("武汉")
	fmt.Printf("len: %d 省:%s 市：%s 级别：%s", len(ans), ans[0].Province, ans[0].City, ans[0].Level)
}
