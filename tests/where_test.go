package tests

import (
	"github.com/jiajun-c/where"
	"testing"
)

func TestWhere(t *testing.T) {
	ans, _ := where.IsArea("洪山区")
	if ans[0].City != "武汉市" || ans[0].Province != "湖北省" {
		t.Error("failed")
	}
}
