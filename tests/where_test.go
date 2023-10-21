package tests

import (
	"fmt"
	"github.com/jiajun-c/where"
	"testing"
)

func TestWhere(t *testing.T) {
	ans, _ := where.Where("武汉")
	fmt.Println(ans)
}
