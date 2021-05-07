package sqls

import (
	"testing"
)

func TestHumiRes(t *testing.T) {
	go func() {

		for i := 0; i < 200; i++ {
			if !HumiRes("50") {
				t.Error(i, "错误")
			}
		}
	}()

	go func() {
		for i := 0; i < 200; i++ {
			if !TempRes("50") {
				t.Error(i, "错误")
			}
		}
	}()

	go func() {
		for i := 0; i < 200; i++ {
			GetWeekTempHumi()
		}
	}()
}
